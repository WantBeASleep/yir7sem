package uzi

import (
	"context"
	pb "yir/uzi/api/broker"
	"yir/uzi/internal/api/mvpmappers"
	"yir/uzi/internal/api/usecases"

	"google.golang.org/protobuf/proto"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Broker struct {
	uziUseCase usecases.Uzi

	logger *zap.Logger
}

func NewBroker(
	uziUseCase usecases.Uzi,
	logger *zap.Logger,
) *Broker {
	return &Broker{
		uziUseCase: uziUseCase,
		logger:     logger,
	}
}

func (b *Broker) ProcessingEvents(ctx context.Context, topic string, msg []byte) error {
	b.logger.Info("[EVENTS] New event", zap.String("Topic", topic))

	// dependency injection
	switch topic {
	case "uziUpload":
		var eventData pb.UziUpload
		if err := proto.Unmarshal(msg, &eventData); err != nil {
			b.logger.Error("[EVENTS] parse data", zap.Error(err))
			return nil
		}

		uziID, err := uuid.Parse(eventData.Id)
		if err != nil {
			b.logger.Error("[EVENTS] parse uzi uuid", zap.Error(err))
			return nil
		}

		pagesID, err := b.uziUseCase.SplitLoadSaveUzi(ctx, uziID)
		if err != nil {
			b.logger.Error("[EVENTS] split load save", zap.Error(err))
			return nil
		}

		b.logger.Info("[EVENTS] Successfull process uziUpload event", zap.Any("pages id", pagesID))

	case "uziProcessed":
		var eventData pb.UziProcessed
		if err := proto.Unmarshal(msg, &eventData); err != nil {
			b.logger.Error("[EVENTS] parse data", zap.Error(err))
			return nil
		}

		formations := mvpmappers.KafkaFormationsToDTOFormations(eventData.Formations)
		segments := mvpmappers.KafkaSegmentsToDTOSegments(eventData.Segments)

		if err := b.uziUseCase.InsertFormationsAndSegemetsSeparately(ctx, formations, segments); err != nil {
			b.logger.Error("[EVENTS] insert formations and segments", zap.Error(err))
			return nil
		}

		b.logger.Info("[EVENTS] Successfull process uziProcessed event")

	default:
		b.logger.Error("[EVENTS] unsupported event")
	}

	return nil
}
