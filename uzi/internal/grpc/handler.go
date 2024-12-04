package grpc

import (
	"yirv2/uzi/internal/generated/grpc/service"
	"yirv2/uzi/internal/grpc/device"
	"yirv2/uzi/internal/grpc/image"
	"yirv2/uzi/internal/grpc/node"
	"yirv2/uzi/internal/grpc/segment"
	"yirv2/uzi/internal/grpc/uzi"
)

type Handler struct {
	device.DeviceHandler
	uzi.UziHandler
	image.ImageHandler
	node.NodeHandler
	segment.SegmentHandler

	service.UnsafeUziSrvServer
}

func New(
	deviceHandler device.DeviceHandler,
	uziHandler uzi.UziHandler,
	imageHandler image.ImageHandler,
	nodeHandler node.NodeHandler,
	segmentHandler segment.SegmentHandler,
) *Handler {
	return &Handler{
		DeviceHandler:  deviceHandler,
		UziHandler:     uziHandler,
		ImageHandler:   imageHandler,
		NodeHandler:    nodeHandler,
		SegmentHandler: segmentHandler,
	}
}
