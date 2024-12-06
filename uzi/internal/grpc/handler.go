package grpc

import (
	"uzi/internal/generated/grpc/service"
	"uzi/internal/grpc/device"
	"uzi/internal/grpc/image"
	"uzi/internal/grpc/node"
	"uzi/internal/grpc/segment"
	"uzi/internal/grpc/uzi"
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
