//go:build e2e

package e2e_test

import (
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

func (suite *TestSuite) TestNode_Success() {
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

	// создаем узел с сегментами
	createNodeResp, err := suite.grpcClient.CreateNode(suite.ctx, &pb.CreateNodeIn{
		UziId: uziResp.Id,
		Segments: []*pb.CreateNodeIn_NestedSegment{
			{
				ImageId:   uziImagesResp.Images[0].Id,
				Contor:    "contor1",
				Tirads_23: 0.1,
				Tirads_4:  0.2,
				Tirads_5:  0.3,
			},
			{
				ImageId:   uziImagesResp.Images[1].Id,
				Contor:    "contor2",
				Tirads_23: 0.4,
				Tirads_4:  0.5,
				Tirads_5:  0.6,
			},
		},
		Tirads_23: 0.1,
		Tirads_4:  0.2,
		Tirads_5:  0.3,
	})
	require.NoError(suite.T(), err)

	// проверяем что узел создался
	nodesResp, err := suite.grpcClient.GetAllNodes(suite.ctx, &pb.GetAllNodesIn{
		UziId: uziResp.Id,
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), 1, len(nodesResp.Nodes))
	require.Equal(suite.T(), createNodeResp.Id, nodesResp.Nodes[0].Id)
	require.Equal(suite.T(), false, nodesResp.Nodes[0].Ai)
	require.Equal(suite.T(), 0.1, nodesResp.Nodes[0].Tirads_23)
	require.Equal(suite.T(), 0.2, nodesResp.Nodes[0].Tirads_4)
	require.Equal(suite.T(), 0.3, nodesResp.Nodes[0].Tirads_5)

	// проверяем что сегменты создались
	segmentsResp, err := suite.grpcClient.GetNodeSegments(suite.ctx, &pb.GetNodeSegmentsIn{
		NodeId: createNodeResp.Id,
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), 2, len(segmentsResp.Segments))
	require.Equal(suite.T(), uziImagesResp.Images[0].Id, segmentsResp.Segments[0].ImageId)
	require.Equal(suite.T(), "contor1", segmentsResp.Segments[0].Contor)
	require.Equal(suite.T(), 0.1, segmentsResp.Segments[0].Tirads_23)
	require.Equal(suite.T(), 0.2, segmentsResp.Segments[0].Tirads_4)
	require.Equal(suite.T(), 0.3, segmentsResp.Segments[0].Tirads_5)
	require.Equal(suite.T(), uziImagesResp.Images[1].Id, segmentsResp.Segments[1].ImageId)
	require.Equal(suite.T(), "contor2", segmentsResp.Segments[1].Contor)
	require.Equal(suite.T(), 0.4, segmentsResp.Segments[1].Tirads_23)
	require.Equal(suite.T(), 0.5, segmentsResp.Segments[1].Tirads_4)
	require.Equal(suite.T(), 0.6, segmentsResp.Segments[1].Tirads_5)

	// обновляем узел
	updateNodeResp, err := suite.grpcClient.UpdateNode(suite.ctx, &pb.UpdateNodeIn{
		Id:        createNodeResp.Id,
		Tirads_23: gtc.ValueToPointer(0.4),
		Tirads_4:  gtc.ValueToPointer(0.5),
		Tirads_5:  gtc.ValueToPointer(0.6),
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), createNodeResp.Id, updateNodeResp.Node.Id)
	require.Equal(suite.T(), false, updateNodeResp.Node.Ai)
	require.Equal(suite.T(), 0.4, updateNodeResp.Node.Tirads_23)
	require.Equal(suite.T(), 0.5, updateNodeResp.Node.Tirads_4)
	require.Equal(suite.T(), 0.6, updateNodeResp.Node.Tirads_5)

	// проверяем что узел обновился
	nodesResp, err = suite.grpcClient.GetAllNodes(suite.ctx, &pb.GetAllNodesIn{
		UziId: uziResp.Id,
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), 1, len(nodesResp.Nodes))
	require.Equal(suite.T(), createNodeResp.Id, nodesResp.Nodes[0].Id)
	require.Equal(suite.T(), false, nodesResp.Nodes[0].Ai)
	require.Equal(suite.T(), 0.4, nodesResp.Nodes[0].Tirads_23)
	require.Equal(suite.T(), 0.5, nodesResp.Nodes[0].Tirads_4)
	require.Equal(suite.T(), 0.6, nodesResp.Nodes[0].Tirads_5)

	// удаляем узел
	_, err = suite.grpcClient.DeleteNode(suite.ctx, &pb.DeleteNodeIn{
		Id: createNodeResp.Id,
	})
	require.NoError(suite.T(), err)

	// проверяем что узел удалился
	nodesResp, err = suite.grpcClient.GetAllNodes(suite.ctx, &pb.GetAllNodesIn{
		UziId: uziResp.Id,
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), 0, len(nodesResp.Nodes))

	// проверяем что сегменты удалились
	segmentsResp, err = suite.grpcClient.GetNodeSegments(suite.ctx, &pb.GetNodeSegmentsIn{
		NodeId: createNodeResp.Id,
	})
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), 0, len(segmentsResp.Segments))
}
