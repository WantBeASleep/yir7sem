package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
)

type Producer struct {
	Addrs []string
	Topic string
}

func New(addr []string, topic string) *Producer {
	return &Producer{
		Addrs: addr,
		Topic: topic,
	}
}

func (k *Producer) Send(key string, payload []byte) error {
	prod, err := sarama.NewSyncProducer(k.Addrs, nil)
	if err != nil {
		return fmt.Errorf("failed to open kafka producer: %w", err)
	}

	prod.SendMessage(&sarama.ProducerMessage{
		Topic: k.Topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(payload),
	})
	if err := prod.Close(); err != nil {
		return fmt.Errorf("close kafka producer: %w", err)
	}
	return nil
}
