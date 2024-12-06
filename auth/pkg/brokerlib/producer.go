package brokerlib

import (
	"fmt"

	"github.com/IBM/sarama"
)

type Producer interface {
	Send(topic, key string, payload []byte) error
}

type producer struct {
	prod sarama.SyncProducer
}

func NewProducer(addr []string) (Producer, error) {
	prod, err := sarama.NewSyncProducer(addr, nil)
	if err != nil {
		return nil, fmt.Errorf("create new producer: %w", err)
	}

	return &producer{
		prod: prod,
	}, nil
}

func (p *producer) Send(topic, key string, payload []byte) error {
	_, _, err := p.prod.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(payload),
	})
	if err != nil {
		return fmt.Errorf("send msg: %w", err)
	}

	return nil
}