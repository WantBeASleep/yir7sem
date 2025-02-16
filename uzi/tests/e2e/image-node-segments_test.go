//go:build e2e

package e2e_test

import (
	"math/rand"
	"os"
	"path/filepath"
	"time"

	pbBrokerProcessed "uzi/internal/generated/broker/consume/uziprocessed"
	pbBrokerUpload "uzi/internal/generated/broker/consume/uziupload"
	pb "uzi/internal/generated/grpc/service"

	"github.com/IBM/sarama"
	"github.com/WantBeASleep/med_ml_lib/slicer"
	"github.com/google/uuid"
	minio "github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func (suite *TestSuite) TestUziProcessed_Success_FakeMLService() {
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
	msgUpload, err := proto.Marshal(&pbBrokerUpload.UziUpload{
		UziId: uziResp.Id,
	})
	require.NoError(suite.T(), err)

	_, _, err = suite.dbusClient.SendMessage(
		&sarama.ProducerMessage{
			Topic: "uziupload",
			Value: sarama.ByteEncoder(msgUpload),
		},
	)
	require.NoError(suite.T(), err)

	// TODO: расписать состояние узи, принято оно в обработку или нет
	// здесь сделаем костыль с time.sleep

	time.Sleep(time.Second * 5)

	// TODO: подумать над стабильностью теста в зависимости от версии нейронки
	// на данный момент нейронка делает 20 частей

	uziImagesResp, err := suite.grpcClient.GetUziImages(suite.ctx, &pb.GetUziImagesIn{
		UziId: uziResp.Id,
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), 20, len(uziImagesResp.Images))

	// создаем событие сохранения узлов и сегментов
	msgProcessed := pbBrokerProcessed.UziProcessed{
		Nodes:    []*pbBrokerProcessed.UziProcessed_Node{},
		Segments: []*pbBrokerProcessed.UziProcessed_Segment{},
	}

	type imageNodesSegments struct {
		Nodes    map[*pb.Node]struct{}
		Segments []*pb.Segment
	}

	expectedImageNodesSegments := map[string]imageNodesSegments{}

	// создаем 10 узлов
	for range 10 {
		node := pbBrokerProcessed.UziProcessed_Node{
			Id:        uuid.New().String(),
			UziId:     uziResp.Id,
			Tirads_23: 1,
			Tirads_4:  1,
			Tirads_5:  1,
		}

		msgProcessed.Nodes = append(msgProcessed.Nodes, &node)

		// TODO: добавить AI в segment
		// создаем 5 сегментов
		for range 5 {
			segmentImg := rand.Intn(len(uziImagesResp.Images))
			segment := pbBrokerProcessed.UziProcessed_Segment{
				Id:        uuid.New().String(),
				NodeId:    node.Id,
				ImageId:   uziImagesResp.Images[segmentImg].Id,
				Contor:    "contor",
				Tirads_23: 1,
				Tirads_4:  1,
				Tirads_5:  1,
			}

			msgProcessed.Segments = append(msgProcessed.Segments, &segment)

			// добавляем узел и сегмент в ожидаемый результат
			v := expectedImageNodesSegments[uziImagesResp.Images[segmentImg].Id]
			if v.Nodes == nil {
				v.Nodes = map[*pb.Node]struct{}{}
			}
			v.Nodes[uziProcessedNodeToPb(&node)] = struct{}{}
			v.Segments = append(v.Segments, uziProcessedSegmentToPb(&segment))
			expectedImageNodesSegments[uziImagesResp.Images[segmentImg].Id] = v
		}
	}

	// тригерим событие сохранения узлов и сегментов
	msgProcessedBytes, err := proto.Marshal(&msgProcessed)
	require.NoError(suite.T(), err)

	_, _, err = suite.dbusClient.SendMessage(
		&sarama.ProducerMessage{
			Topic: "uziprocessed",
			Value: sarama.ByteEncoder(msgProcessedBytes),
		},
	)
	require.NoError(suite.T(), err)

	// ждем сохранения узлов и сегментов
	time.Sleep(time.Second * 5)

	// проверяем, что узлы и сегменты сохранились правильно
	for _, image := range uziImagesResp.Images {
		imageNodesSegments, err := suite.grpcClient.GetImageSegmentsWithNodes(
			suite.ctx,
			&pb.GetImageSegmentsWithNodesIn{
				Id: image.Id,
			},
		)
		require.NoError(suite.T(), err)

		expectedImageNodes := slicer.MapToSlice(expectedImageNodesSegments[image.Id].Nodes)
		require.ElementsMatch(suite.T(), expectedImageNodes, imageNodesSegments.Nodes)
		require.ElementsMatch(suite.T(), expectedImageNodesSegments[image.Id].Segments, imageNodesSegments.Segments)
	}
}

func uziProcessedNodeToPb(node *pbBrokerProcessed.UziProcessed_Node) *pb.Node {
	return &pb.Node{
		Id: node.Id,
		// при uziprocessed узел всегда AI
		Ai:        true,
		Tirads_23: node.Tirads_23,
		Tirads_4:  node.Tirads_4,
		Tirads_5:  node.Tirads_5,
	}
}

func uziProcessedSegmentToPb(segment *pbBrokerProcessed.UziProcessed_Segment) *pb.Segment {
	return &pb.Segment{
		Id:        segment.Id,
		NodeId:    segment.NodeId,
		ImageId:   segment.ImageId,
		Contor:    segment.Contor,
		Tirads_23: segment.Tirads_23,
		Tirads_4:  segment.Tirads_4,
		Tirads_5:  segment.Tirads_5,
	}
}
