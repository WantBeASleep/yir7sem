package image

import (
	"context"

	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/services/image"

	// TODO: вынести в отдельный пакет маппреы
	nodemapper "uzi/internal/grpc/node"
	segmentmapper "uzi/internal/grpc/segment"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ImageHandler interface {
	GetUziImages(ctx context.Context, in *pb.GetUziImagesIn) (*pb.GetUziImagesOut, error)
	GetImageSegmentsWithNodes(ctx context.Context, in *pb.GetImageSegmentsWithNodesIn) (*pb.GetImageSegmentsWithNodesOut, error)
}

type handler struct {
	imageSrv image.Service
}

func New(
	imageSrv image.Service,
) ImageHandler {
	return &handler{
		imageSrv: imageSrv,
	}
}

func (h *handler) GetUziImages(ctx context.Context, in *pb.GetUziImagesIn) (*pb.GetUziImagesOut, error) {
	images, err := h.imageSrv.GetUziImages(ctx, uuid.MustParse(in.UziId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	out := pb.GetUziImagesOut{}
	for _, v := range images {
		out.Images = append(out.Images, &pb.Image{
			Id:   v.Id.String(),
			Page: int64(v.Page),
		})
	}

	return &out, nil
}

// TODO: вынести это в сегменты или ноды, однозначно не в image
func (h *handler) GetImageSegmentsWithNodes(ctx context.Context, in *pb.GetImageSegmentsWithNodesIn) (*pb.GetImageSegmentsWithNodesOut, error) {
	nodes, segments, err := h.imageSrv.GetImageSegmentsWithNodes(ctx, uuid.MustParse(in.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	pbnodes := make([]*pb.Node, 0, len(nodes))
	for _, v := range nodes {
		pbnodes = append(pbnodes, nodemapper.DomainNodeToPb(&v))
	}

	pbsegments := make([]*pb.Segment, 0, len(segments))
	for _, v := range segments {
		pbsegments = append(pbsegments, segmentmapper.DomainSegmentToPb(&v))
	}

	return &pb.GetImageSegmentsWithNodesOut{
		Nodes:    pbnodes,
		Segments: pbsegments,
	}, nil
}
