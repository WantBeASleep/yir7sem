package uziupload

import (
	"context"
	"errors"
	"fmt"

	"yirv2/pkg/brokerlib"
	pb "yirv2/uzi/internal/generated/broker/consume/uziupload"
	"yirv2/uzi/internal/services/image"

	"github.com/google/uuid"
)

const (
	groupID = "uziupload"
	topic   = "uziupload"
)

type subscriber struct {
	imageSrv image.Service
}

func New(
	imageSrv image.Service,
) brokerlib.SubscriberStrategy {
	return &subscriber{
		imageSrv: imageSrv,
	}
}

func (h *subscriber) GetConfig() brokerlib.SubscriberConfig {
	return brokerlib.SubscriberConfig{
		GroupID: groupID,
		Topics:  []string{topic},
	}
}

func (h *subscriber) ProcessMessage(ctx context.Context, msg any) error {
	payload, ok := msg.(*pb.UziUpload)
	if !ok {
		return errors.New("wrong msg type. uziupload required")
	}

	if err := h.imageSrv.SplitUzi(ctx, uuid.MustParse(payload.UziId)); err != nil {
		return fmt.Errorf("process uziupload: %w", err)
	}
	return nil
}
