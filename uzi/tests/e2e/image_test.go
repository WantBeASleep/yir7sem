//go:build e2e

package e2e_test

import (
	"os"
	"path/filepath"
	"time"

	pbBrokerUpload "uzi/internal/generated/broker/consume/uziupload"
	pb "uzi/internal/generated/grpc/service"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
	minio "github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func (suite *TestSuite) TestSplitImage_Success() {
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

	page := 1
	for _, image := range uziImagesResp.Images {
		_, err = uuid.Parse(image.Id)
		require.NoError(suite.T(), err)

		require.Equal(suite.T(), page, int(image.Page))
		page++
	}
}
