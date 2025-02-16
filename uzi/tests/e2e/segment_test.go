//go:build e2e

package e2e_test

import (
	"math/rand"
	"os"
	"path/filepath"
	"time"

	pbBrokerUpload "uzi/internal/generated/broker/consume/uziupload"
	pb "uzi/internal/generated/grpc/service"

	"github.com/IBM/sarama"
	"github.com/WantBeASleep/med_ml_lib/gtc"
	"github.com/google/uuid"
	minio "github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func (suite *TestSuite) TestSegment_Success() {
	// создаем девайс
	deviceResp, err := suite.grpcClient.CreateDevice(suite.ctx, &pb.CreateDeviceIn{
		Name: "test_device",
	})
	require.NoError(suite.T(), err)

	// создаем uzi
	uziResp, err := suite.grpcClient.CreateUzi(suite.ctx, &pb.CreateUziIn{
		Projection: "axial",
		PatientId:  uuid.New().String(),
		DeviceId:   deviceResp.Id,
	})
	require.NoError(suite.T(), err)

	// Открываем локальный TIFF файл
	tiffFile, err := os.Open("assets/sample.tiff")
	require.NoError(suite.T(), err)
	defer tiffFile.Close()

	// Читаем содержимое файла
	fileInfo, err := tiffFile.Stat()
	require.NoError(suite.T(), err)

	_, err = suite.s3Client.PutObject(
		suite.ctx,
		suite.s3Bucket,
		filepath.Join(uziResp.Id, uziResp.Id),
		tiffFile,
		fileInfo.Size(),
		minio.PutObjectOptions{ContentType: "image/tiff"},
	)
	require.NoError(suite.T(), err)

	// тригерим событие обработки загруженного узи
	msg, err := proto.Marshal(&pbBrokerUpload.UziUpload{
		UziId: uziResp.Id,
	})
	require.NoError(suite.T(), err)

	_, _, err = suite.dbusClient.SendMessage(
		&sarama.ProducerMessage{
			Topic: "uziupload",
			Value: sarama.ByteEncoder(msg),
		},
	)
	require.NoError(suite.T(), err)

	// ждем обработки изображений
	time.Sleep(time.Second * 5)

	// получаем изображения
	uziImagesResp, err := suite.grpcClient.GetUziImages(suite.ctx, &pb.GetUziImagesIn{
		UziId: uziResp.Id,
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), 20, len(uziImagesResp.Images))

	// создаем узел
	createNodeResp, err := suite.grpcClient.CreateNode(suite.ctx, &pb.CreateNodeIn{
		UziId:     uziResp.Id,
		Segments:  []*pb.CreateNodeIn_NestedSegment{},
		Tirads_23: 0.1,
		Tirads_4:  0.2,
		Tirads_5:  0.3,
	})
	require.NoError(suite.T(), err)

	// создаем сегмент
	createSegmentResp, err := suite.grpcClient.CreateSegment(suite.ctx, &pb.CreateSegmentIn{
		NodeId:    createNodeResp.Id,
		ImageId:   uziImagesResp.Images[0].Id,
		Contor:    "contor1",
		Tirads_23: 0.1,
		Tirads_4:  0.2,
		Tirads_5:  0.3,
	})
	require.NoError(suite.T(), err)

	// проверяем что сегмент создался
	segmentsResp, err := suite.grpcClient.GetNodeSegments(suite.ctx, &pb.GetNodeSegmentsIn{
		NodeId: createNodeResp.Id,
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), 1, len(segmentsResp.Segments))
	require.Equal(suite.T(), createSegmentResp.Id, segmentsResp.Segments[0].Id)
	require.Equal(suite.T(), uziImagesResp.Images[0].Id, segmentsResp.Segments[0].ImageId)
	require.Equal(suite.T(), "contor1", segmentsResp.Segments[0].Contor)
	require.Equal(suite.T(), 0.1, segmentsResp.Segments[0].Tirads_23)
	require.Equal(suite.T(), 0.2, segmentsResp.Segments[0].Tirads_4)
	require.Equal(suite.T(), 0.3, segmentsResp.Segments[0].Tirads_5)

	// обновляем сегмент
	updateSegmentResp, err := suite.grpcClient.UpdateSegment(suite.ctx, &pb.UpdateSegmentIn{
		Id:        createSegmentResp.Id,
		Tirads_23: gtc.ValueToPointer(0.4),
		Tirads_4:  gtc.ValueToPointer(0.5),
		Tirads_5:  gtc.ValueToPointer(0.6),
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), createSegmentResp.Id, updateSegmentResp.Segment.Id)
	require.Equal(suite.T(), uziImagesResp.Images[0].Id, updateSegmentResp.Segment.ImageId)
	require.Equal(suite.T(), "contor1", updateSegmentResp.Segment.Contor)
	require.Equal(suite.T(), 0.4, updateSegmentResp.Segment.Tirads_23)
	require.Equal(suite.T(), 0.5, updateSegmentResp.Segment.Tirads_4)
	require.Equal(suite.T(), 0.6, updateSegmentResp.Segment.Tirads_5)

	// проверяем что сегмент обновился
	segmentsResp, err = suite.grpcClient.GetNodeSegments(suite.ctx, &pb.GetNodeSegmentsIn{
		NodeId: createNodeResp.Id,
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), 1, len(segmentsResp.Segments))
	require.Equal(suite.T(), createSegmentResp.Id, segmentsResp.Segments[0].Id)
	require.Equal(suite.T(), uziImagesResp.Images[0].Id, segmentsResp.Segments[0].ImageId)
	require.Equal(suite.T(), "contor1", segmentsResp.Segments[0].Contor)
	require.Equal(suite.T(), 0.4, segmentsResp.Segments[0].Tirads_23)
	require.Equal(suite.T(), 0.5, segmentsResp.Segments[0].Tirads_4)
	require.Equal(suite.T(), 0.6, segmentsResp.Segments[0].Tirads_5)

	// удаляем сегмент
	_, err = suite.grpcClient.DeleteSegment(suite.ctx, &pb.DeleteSegmentIn{
		Id: createSegmentResp.Id,
	})
	require.NoError(suite.T(), err)

	// проверяем что сегмент удалился
	segmentsResp, err = suite.grpcClient.GetNodeSegments(suite.ctx, &pb.GetNodeSegmentsIn{
		NodeId: createNodeResp.Id,
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), 0, len(segmentsResp.Segments))
}

func (suite *TestSuite) TestSegments_MultipleImages() {
	// создаем девайс
	deviceResp, err := suite.grpcClient.CreateDevice(suite.ctx, &pb.CreateDeviceIn{
		Name: "test_device",
	})
	require.NoError(suite.T(), err)

	// создаем uzi
	uziResp, err := suite.grpcClient.CreateUzi(suite.ctx, &pb.CreateUziIn{
		Projection: "axial",
		PatientId:  uuid.New().String(),
		DeviceId:   deviceResp.Id,
	})
	require.NoError(suite.T(), err)

	// Открываем локальный TIFF файл
	tiffFile, err := os.Open("assets/sample.tiff")
	require.NoError(suite.T(), err)
	defer tiffFile.Close()

	// Читаем содержимое файла
	fileInfo, err := tiffFile.Stat()
	require.NoError(suite.T(), err)

	_, err = suite.s3Client.PutObject(
		suite.ctx,
		suite.s3Bucket,
		filepath.Join(uziResp.Id, uziResp.Id),
		tiffFile,
		fileInfo.Size(),
		minio.PutObjectOptions{ContentType: "image/tiff"},
	)
	require.NoError(suite.T(), err)

	// тригерим событие обработки загруженного узи
	msg, err := proto.Marshal(&pbBrokerUpload.UziUpload{
		UziId: uziResp.Id,
	})
	require.NoError(suite.T(), err)

	_, _, err = suite.dbusClient.SendMessage(
		&sarama.ProducerMessage{
			Topic: "uziupload",
			Value: sarama.ByteEncoder(msg),
		},
	)
	require.NoError(suite.T(), err)

	// ждем обработки изображений
	time.Sleep(time.Second * 5)

	// получаем изображения
	uziImagesResp, err := suite.grpcClient.GetUziImages(suite.ctx, &pb.GetUziImagesIn{
		UziId: uziResp.Id,
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), 20, len(uziImagesResp.Images))

	// создаем узел
	createNodeResp, err := suite.grpcClient.CreateNode(suite.ctx, &pb.CreateNodeIn{
		UziId:     uziResp.Id,
		Segments:  []*pb.CreateNodeIn_NestedSegment{},
		Tirads_23: 0.1,
		Tirads_4:  0.2,
		Tirads_5:  0.3,
	})
	require.NoError(suite.T(), err)

	// создаем мапу ожидаемых сегментов по картинкам
	expectedSegments := make(map[string][]*pb.Segment)

	// создаем 10 сегментов на разных картинках
	for i := 0; i < 10; i++ {
		imageIdx := rand.Intn(len(uziImagesResp.Images))
		imageId := uziImagesResp.Images[imageIdx].Id

		createSegmentResp, err := suite.grpcClient.CreateSegment(suite.ctx, &pb.CreateSegmentIn{
			NodeId:    createNodeResp.Id,
			ImageId:   imageId,
			Contor:    "contor",
			Tirads_23: 0.1,
			Tirads_4:  0.2,
			Tirads_5:  0.3,
		})
		require.NoError(suite.T(), err)
		// TODO: здесь воспользоваться маппером из хендлера
		segment := &pb.Segment{
			Id:        createSegmentResp.Id,
			NodeId:    createNodeResp.Id,
			ImageId:   imageId,
			Contor:    "contor",
			Tirads_23: 0.1,
			Tirads_4:  0.2,
			Tirads_5:  0.3,
		}

		expectedSegments[imageId] = append(expectedSegments[imageId], segment)
	}

	// проверяем сегменты на каждой картинке
	for _, image := range uziImagesResp.Images {
		imageSegments, err := suite.grpcClient.GetImageSegmentsWithNodes(suite.ctx, &pb.GetImageSegmentsWithNodesIn{
			Id: image.Id,
		})
		require.NoError(suite.T(), err)
		require.ElementsMatch(suite.T(), expectedSegments[image.Id], imageSegments.Segments)
	}
}
