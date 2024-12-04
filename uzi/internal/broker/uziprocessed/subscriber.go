package uziprocessed

import (
	"context"
	"errors"
	"fmt"

	"yirv2/pkg/brokerlib"
	"yirv2/uzi/internal/domain"
	pb "yirv2/uzi/internal/generated/broker/consume/uziprocessed"
	"yirv2/uzi/internal/services/node"

	"github.com/google/uuid"
)

const (
	groupID = "uziprocessed"
	topic   = "uziprocessed"
)

type subscriber struct {
	nodeSrv node.Service
}

func New(
	nodeSrv node.Service,
) brokerlib.SubscriberStrategy {
	return &subscriber{
		nodeSrv: nodeSrv,
	}
}

func (h *subscriber) GetConfig() brokerlib.SubscriberConfig {
	return brokerlib.SubscriberConfig{
		GroupID: groupID,
		Topics:  []string{topic},
	}
}

func (h *subscriber) ProcessMessage(ctx context.Context, msg any) error {
	payload, ok := msg.(*pb.UziProcessed)
	if !ok {
		return errors.New("wrong msg type. uziprocessed required")
	}

	nodes := make([]domain.Node, 0, len(payload.Nodes))
	segments := make([]domain.Segment, 0, len(payload.Segments))

	for _, v := range payload.Nodes {
		nodes = append(nodes, domain.Node{
			Id:       uuid.MustParse(v.Id),
			Tirads23: v.Tirads_23,
			Tirads4:  v.Tirads_4,
			Tirads5:  v.Tirads_5,
		})
	}

	for _, v := range payload.Segments {
		segments = append(segments, domain.Segment{
			Id:       uuid.MustParse(v.Id),
			ImageID:  uuid.MustParse(v.ImageId),
			NodeID:   uuid.MustParse(v.NodeId),
			Contor:   v.Contor,
			Tirads23: v.Tirads_23,
			Tirads4:  v.Tirads_4,
			Tirads5:  v.Tirads_5,
		})
	}

	if err := h.nodeSrv.InsertAiNodeWithSegments(ctx, nodes, segments); err != nil {
		return fmt.Errorf("isert ai nodes && segments: %w", err)
	}
	return nil
}
