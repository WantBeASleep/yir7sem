package node

import (
	"context"

	"uzi/internal/domain"
	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/services/node"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NodeHandler interface {
	CreateNode(ctx context.Context, in *pb.CreateNodeIn) (*pb.CreateNodeOut, error)
	GetAllNodes(ctx context.Context, in *pb.GetAllNodesIn) (*pb.GetAllNodesOut, error)
	DeleteNode(ctx context.Context, in *pb.DeleteNodeIn) (*empty.Empty, error)
	UpdateNode(ctx context.Context, in *pb.UpdateNodeIn) (*pb.UpdateNodeOut, error)
}

type handler struct {
	nodeSrv node.Service
}

func New(
	nodeSrv node.Service,
) NodeHandler {
	return &handler{
		nodeSrv: nodeSrv,
	}
}

func (h *handler) CreateNode(ctx context.Context, in *pb.CreateNodeIn) (*pb.CreateNodeOut, error) {
	segments := make([]domain.Segment, 0, len(in.Segments))
	for _, v := range in.Segments {
		segments = append(segments, domain.Segment{
			ImageID:  uuid.MustParse(v.ImageId),
			Contor:   v.Contor,
			Tirads23: v.Tirads_23,
			Tirads4:  v.Tirads_4,
			Tirads5:  v.Tirads_5,
		})
	}

	nodeID, err := h.nodeSrv.CreateNode(
		ctx,
		domain.Node{
			UziID:    uuid.MustParse(in.UziId),
			Tirads23: in.Tirads_23,
			Tirads4:  in.Tirads_4,
			Tirads5:  in.Tirads_5,
		},
		segments,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.CreateNodeOut{
		Id: nodeID.String(),
	}, nil
}

func (h *handler) GetAllNodes(ctx context.Context, in *pb.GetAllNodesIn) (*pb.GetAllNodesOut, error) {
	nodes, err := h.nodeSrv.GetAllNodes(ctx, uuid.MustParse(in.UziId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	outNodes := make([]*pb.Node, 0, len(nodes))
	for _, v := range nodes {
		outNodes = append(outNodes, DomainNodeToPb(&v))
	}

	return &pb.GetAllNodesOut{
		Nodes: outNodes,
	}, nil
}

func (h *handler) DeleteNode(ctx context.Context, in *pb.DeleteNodeIn) (*empty.Empty, error) {
	if err := h.nodeSrv.DeleteNode(ctx, uuid.MustParse(in.Id)); err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}
	return &empty.Empty{}, nil
}

func (h *handler) UpdateNode(ctx context.Context, in *pb.UpdateNodeIn) (*pb.UpdateNodeOut, error) {
	node, err := h.nodeSrv.UpdateNode(
		ctx,
		uuid.MustParse(in.Id),
		node.UpdateNode{
			Tirads23: in.Tirads_23,
			Tirads4:  in.Tirads_4,
			Tirads5:  in.Tirads_5,
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.UpdateNodeOut{
		Node: DomainNodeToPb(&node),
	}, nil
}
