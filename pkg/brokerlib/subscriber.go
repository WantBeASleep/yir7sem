// TODO: пакет полного говна, нужно будет переделать. (ничего критичного нет вроде бы как)
package brokerlib

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
)

type SubscriberConfig struct {
	GroupID string
	Topics  []string
}

type SubscriberStrategy interface {
	GetConfig() SubscriberConfig
	ProcessMessage(ctx context.Context, message []byte) error
}

type SubscriberHandler interface {
	Start(ctx context.Context) error
	Close() error
}

type handler struct {
	sub      SubscriberStrategy
	consumer sarama.ConsumerGroup
}

// TODO: переделать это говнище, добавить прокидывание ID в кафке по человечески
func (*handler) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*handler) Cleanup(sarama.ConsumerGroupSession) error { return nil }
func (h *handler) ConsumeClaim(s sarama.ConsumerGroupSession, c sarama.ConsumerGroupClaim) error {
	for msg := range c.Messages() {
		if err := h.sub.ProcessMessage(s.Context(), msg.Value); err != nil {
			return fmt.Errorf("processing msg: %w", err)
		}
		s.MarkMessage(msg, "")
		s.Commit()
	}

	return nil
}

func (h *handler) Start(ctx context.Context) error {
	for {
		if err := h.consumer.Consume(ctx, h.sub.GetConfig().Topics, h); err != nil {
			return fmt.Errorf("listen topics: %v, error: %w", h.sub.GetConfig().Topics, err)
		}
	}
}

func (h *handler) Close() error { return nil }

// TODO: тут нужны опции настройки нормальные
func GetSubscriberHandler(sub SubscriberStrategy, hosts []string, cfg *sarama.Config) (SubscriberHandler, error) {
	consumer, err := sarama.NewConsumerGroup(hosts, sub.GetConfig().GroupID, cfg)
	if err != nil {
		return nil, fmt.Errorf("create new group: %w", err)
	}

	return &handler{
		sub:      sub,
		consumer: consumer,
	}, nil
}
