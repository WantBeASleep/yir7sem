package broker

import (
	"fmt"

	"github.com/WantBeASleep/goooool/brokerlib"

	pb "gateway/internal/generated/broker/produce/uziupload"

	"google.golang.org/protobuf/proto"
)

const (
	uziuploadTopic = "uziupload"
)

type BrokerAdapter interface {
	SendUziUpload(msg *pb.UziUpload) error
}

// TODO: переписать библу/хотя бы в интерфейс обернуть продьюсера
func New(
	producer brokerlib.Producer,
) BrokerAdapter {
	return &adapter{
		producer: producer,
	}
}

type adapter struct {
	producer brokerlib.Producer
}

func (a *adapter) SendUziUpload(msg *pb.UziUpload) error {
	// TODO: когда будем делать партицированние пробрасывать сюда ключи
	payload, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal uziupload event: %w", err)
	}
	return a.producer.Send(uziuploadTopic, "52", payload)
}
