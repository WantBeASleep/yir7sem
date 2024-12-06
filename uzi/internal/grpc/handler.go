package grpc

import (
	"yir/uzi/internal/generated/grpc/service"
	"yir/uzi/internal/grpc/device"
	"yir/uzi/internal/grpc/image"
	"yir/uzi/internal/grpc/node"
	"yir/uzi/internal/grpc/segment"
	"yir/uzi/internal/grpc/uzi"
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
