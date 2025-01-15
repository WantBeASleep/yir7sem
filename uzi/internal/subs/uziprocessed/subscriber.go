package uziprocessed

import (
	"context"
	"errors"
	"fmt"

	"github.com/WantBeASleep/goooool/brokerlib"

	"uzi/internal/domain"
	pb "uzi/internal/generated/broker/consume/uziprocessed"
	"uzi/internal/services/node"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
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

func (h *subscriber) ProcessMessage(ctx context.Context, msg []byte) error {
	var event pb.UziProcessed
	if err := proto.Unmarshal(msg, &event); err != nil {
		return errors.New("wrong msg type. uziprocessed required")
	}

	nodes := make([]domain.Node, 0, len(event.Nodes))
	segments := make([]domain.Segment, 0, len(event.Segments))

	for _, v := range event.Nodes {
		nodes = append(nodes, domain.Node{
			Id:       uuid.MustParse(v.Id),
			UziID:    uuid.MustParse(v.UziId),
			Tirads23: v.Tirads_23,
			Tirads4:  v.Tirads_4,
			Tirads5:  v.Tirads_5,
		})
	}

	for _, v := range event.Segments {
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
