package uzi

import (
	"context"
	pb "yir/uzi/api/broker"
	"yir/uzi/internal/api/mvpmappers"
	"yir/uzi/internal/api/usecases"

	"yir/pkg/kafka"

	"google.golang.org/protobuf/proto"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Broker struct {
	uziUseCase usecases.Uzi

	logger *zap.Logger
	prod   *kafka.Producer
}

func NewBroker(
	uziUseCase usecases.Uzi,
	logger *zap.Logger,
	prod *kafka.Producer,
) *Broker {
	return &Broker{
		uziUseCase: uziUseCase,
		logger:     logger,
		prod:       prod,
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

		event := &pb.UziSplitted{
			UziId:   uziID.String(),
			PagesId: pagesID.Strings(),
		}
		payload, err := proto.Marshal(event)
		if err != nil {
			b.logger.Error("[EVENTS] marshal uziSplitted event", zap.Error(err))
			return nil
		}
		if err := b.prod.Send("1", payload); err != nil {
			b.logger.Error("[EVENTS] send uziSplitted event", zap.Error(err))
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

		b.logger.Info("[EVENT] INPUT EVENT DATA", zap.Any("Formations", formations), zap.Any("Segments", segments))

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
