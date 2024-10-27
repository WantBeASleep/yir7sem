package kafka

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
)

type EventHandler func(ctx context.Context, topic string, msg []byte) error

type saramaEventWrapper EventHandler
func (saramaEventWrapper) Setup(sarama.ConsumerGroupSession) error {return nil}
func (saramaEventWrapper) Cleanup(sarama.ConsumerGroupSession) error {return nil}
func (w saramaEventWrapper) ConsumeClaim(s sarama.ConsumerGroupSession, c sarama.ConsumerGroupClaim) error {
	for msg := range c.Messages() {
		if err := w(s.Context(), c.Topic(), msg.Value); err != nil {
			return fmt.Errorf("processing msg: %w", err)
		}
		s.MarkMessage(msg, "")
		s.Commit() // надо исходить из идемпонетности операций, пусть так будет
	}

	return nil
}

type GroupConsumer struct {
	topics []string

	consumer sarama.ConsumerGroup
	handler EventHandler
}

func NewGroupConsumer(
	groupID string,
	topics []string,
	hosts []string,
	cfg *sarama.Config,
	handler EventHandler,
) (*GroupConsumer, error) {
	consumer, err := sarama.NewConsumerGroup(hosts, groupID, cfg)
	if err != nil {
		return nil, fmt.Errorf("create new consumer in group[id %q][hosts %v]: %w", groupID, hosts, err)
	}

	return &GroupConsumer{
		topics: topics,

		consumer: consumer,
		handler: handler,
	}, nil
}

// sync block until listen
func (g *GroupConsumer) Start(ctx context.Context) error {
	handler := saramaEventWrapper(g.handler)
	for {
		if err := g.consumer.Consume(ctx, g.topics, handler); err != nil {
			return fmt.Errorf("broker listen msg [topics: %v]: %w", g.topics, err)
		}
	}
}

func (g *GroupConsumer) Close() error {
	if err := g.consumer.Close(); err != nil {
		return fmt.Errorf("close group consumer: %w", err)
	}
	return nil
}
