package repository

import (
	"yir/pkg/kafka"
)

type Producer struct {
	producer *kafka.Producer
}

func New(addr []string, topic string) *Producer {
	return &Producer{
		producer: kafka.New(addr, topic),
	}
}

func (p *Producer) Send(key string, uziID string) error {
	if err := p.producer.Send(key, []byte(uziID)); err != nil {
		return err
	}
	return nil
}
