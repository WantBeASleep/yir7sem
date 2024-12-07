package broker

import (
	"fmt"

	"github.com/WantBeASleep/goooool/brokerlib"

	uzicompletepb "uzi/internal/generated/broker/produce/uzicomplete"
	uzisplittedpb "uzi/internal/generated/broker/produce/uzisplitted"

	"google.golang.org/protobuf/proto"
)

const (
	uzisplittedTopic = "uzisplitted"
	uzicompleteTopic = "uzicomplete"
)

type BrokerAdapter interface {
	SendUziSplitted(msg *uzisplittedpb.UziSplitted) error
	SendUziComplete(msg *uzicompletepb.UziComplete) error
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

func (a *adapter) SendUziSplitted(msg *uzisplittedpb.UziSplitted) error {
	// TODO: когда будем делать партицированние пробрасывать сюда ключи
	payload, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal uzisplitted event: %w", err)
	}
	return a.producer.Send(uzisplittedTopic, "52", payload)
}

func (a *adapter) SendUziComplete(msg *uzicompletepb.UziComplete) error {
	// TODO: когда будем делать партицированние пробрасывать сюда ключи
	payload, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal uzicomplete event: %w", err)
	}
	return a.producer.Send(uzicompleteTopic, "52", payload)
}
