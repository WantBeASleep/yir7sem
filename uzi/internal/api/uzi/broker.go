package uzi

import (
	pb "yir/uzi/events"
	"yir/uzi/internal/api/usecases"

	"google.golang.org/protobuf/proto"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// Как же плохо, просто уауауауййй бляяя
type Broker struct {
	logger *zap.Logger

	uziUseCase usecases.Uzi
}

func NewBroker(
	logger *zap.Logger,
	uziUseCase usecases.Uzi,
) *Broker {
	return &Broker{
		logger:     logger,
		uziUseCase: uziUseCase,
	}
}

func (b *Broker) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (b *Broker) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (b *Broker) ConsumeClaim(s sarama.ConsumerGroupSession, g sarama.ConsumerGroupClaim) error {
	b.logger.Info("[EVENTS] Start consume events", zap.String("Topic", g.Topic()), zap.Int64("Lag", g.HighWaterMarkOffset()))
	for msg := range g.Messages() {
		b.logger.Info("[EVENTS] New event", zap.String("Topic", g.Topic()))

		var eventData pb.UziUploadEvent
		err := proto.Unmarshal(msg.Value, &eventData)
		if err != nil {
			// Tech Debt
			b.logger.Error("[EVENTS] Error while processing event; Parse data", zap.Error(err))
			s.MarkMessage(msg, "")
			s.Commit()
			continue
		}

		uziID, err := uuid.Parse(eventData.UziId)
		if err != nil {
			b.logger.Error("[EVENTS] Error while processing event; Parse uuid", zap.Error(err))
			s.MarkMessage(msg, "")
			s.Commit()
			continue
		}

		pagesID, err := b.uziUseCase.SplitLoadSaveUzi(s.Context(), uziID)
		if err != nil {
			b.logger.Error("[EVENTS] Error while processing event; Event run", zap.Error(err))
			s.MarkMessage(msg, "")
			s.Commit()
			continue
		}

		b.logger.Info("[EVENTS] Successfull process event", zap.Any("pages id", pagesID))
	}

	return nil
}
