package uzi

import (
	"context"
	pb "yir/uzi/events"
	"yir/uzi/internal/api/usecases"
	"yir/uzi/internal/api/mvpmappers"
	"yir/uzi/internal/entity/dto"
	

	"google.golang.org/protobuf/proto"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

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

func (b *Broker) ProcessingEvents(ctx context.Context, topic string, msg []byte) error {
	b.logger.Info("[EVENTS] New event", zap.String("Topic", topic))

	// dependency injection
	switch topic {
	case "uziUpload":
		var eventData pb.UziUpload
		err := proto.Unmarshal(msg, &eventData)
		if err != nil {
			b.logger.Error("[EVENTS] parse data", zap.Error(err))
			return nil
		}

		uziID, err := uuid.Parse(eventData.UziId)
		if err != nil {
			b.logger.Error("[EVENTS] parse uzi uuid", zap.Error(err))
			return nil
		}

		pagesID, err := b.uziUseCase.SplitLoadSaveUzi(ctx, uziID)
		if err != nil {
			b.logger.Error("[EVENTS] split load save", zap.Error(err))
		}

		b.logger.Info("[EVENTS] Successfull process event", zap.Any("pages id", pagesID))

	case "uziProcessed":
		var eventData pb.UziProcessed
		err := proto.Unmarshal(msg, &eventData)
		if err != nil {
			b.logger.Error("[EVENTS] parse data", zap.Error(err))
			return nil
		}

		req := make([]dto.FormationWithSegments, 0, len(eventData.Formations))
		for _, pbFormation := range eventData.Formations {
			
		}

	default:
		b.logger.Error("[EVENTS] unsupported event")
	}

	return nil
}
