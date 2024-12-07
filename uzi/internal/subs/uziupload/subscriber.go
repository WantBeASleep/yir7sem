package uziupload

import (
	"context"
	"errors"
	"fmt"

	"github.com/WantBeASleep/goooool/brokerlib"

	pb "uzi/internal/generated/broker/consume/uziupload"
	"uzi/internal/services/image"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
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

func (h *subscriber) ProcessMessage(ctx context.Context, msg []byte) error {
	var event pb.UziUpload
	if err := proto.Unmarshal(msg, &event); err != nil {
		return errors.New("wrong msg type. uziupload required")
	}

	if err := h.imageSrv.SplitUzi(ctx, uuid.MustParse(event.UziId)); err != nil {
		return fmt.Errorf("process uziupload: %w", err)
	}
	return nil
}
