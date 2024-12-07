package broker

import (
	"fmt"

	"pkg/brokerlib"

	uziuploadpb "gateway/internal/generated/broker/produce/uziupload"

	"google.golang.org/protobuf/proto"
)

const (
	uziuploadTopic = "uziupload"
)

type Adapter interface {
	SendUziUpload(msg *uziuploadpb.UziUpload) error
}

// TODO: переписать библу/хотя бы в интерфейс обернуть продьюсера
func New(
	producer brokerlib.Producer,
) Adapter {
	return &adapter{
		producer: producer,
	}
}

const (
	uzisplittedTopic = "uzisplitted"
)

type adapter struct {
	producer brokerlib.Producer
}

func (a *adapter) SendUziUpload(msg *uziuploadpb.UziUpload) error {
	// TODO: когда будем делать партицированние пробрасывать сюда ключи
	payload, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal uziupload event: %w", err)
	}
	return a.producer.Send(uziuploadTopic, "52", payload)
}
