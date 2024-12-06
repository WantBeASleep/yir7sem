// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: uzi.proto

package uzi

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UziSrv_GetDeviceList_FullMethodName             = "/UziSrv/getDeviceList"
	UziSrv_CreateUzi_FullMethodName                 = "/UziSrv/createUzi"
	UziSrv_UpdateUzi_FullMethodName                 = "/UziSrv/updateUzi"
	UziSrv_UpdateEchographic_FullMethodName         = "/UziSrv/updateEchographic"
	UziSrv_GetUzi_FullMethodName                    = "/UziSrv/getUzi"
	UziSrv_GetUziImages_FullMethodName              = "/UziSrv/getUziImages"
	UziSrv_GetImageSegmentsWithNodes_FullMethodName = "/UziSrv/getImageSegmentsWithNodes"
	UziSrv_CreateSegment_FullMethodName             = "/UziSrv/createSegment"
	UziSrv_DeleteSegment_FullMethodName             = "/UziSrv/deleteSegment"
	UziSrv_UpdateSegment_FullMethodName             = "/UziSrv/updateSegment"
	UziSrv_CreateNode_FullMethodName                = "/UziSrv/createNode"
	UziSrv_DeleteNode_FullMethodName                = "/UziSrv/deleteNode"
	UziSrv_UpdateNode_FullMethodName                = "/UziSrv/updateNode"
)

// UziSrvClient is the client API for UziSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UziSrvClient interface {
	GetDeviceList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDeviceListOut, error)
	CreateUzi(ctx context.Context, in *CreateUziIn, opts ...grpc.CallOption) (*CreateUziOut, error)
	UpdateUzi(ctx context.Context, in *UpdateUziIn, opts ...grpc.CallOption) (*UpdateUziOut, error)
	UpdateEchographic(ctx context.Context, in *UpdateEchographicIn, opts ...grpc.CallOption) (*UpdateEchographicOut, error)
	GetUzi(ctx context.Context, in *GetUziIn, opts ...grpc.CallOption) (*GetUziOut, error)
	GetUziImages(ctx context.Context, in *GetUziImagesIn, opts ...grpc.CallOption) (*GetUziImagesOut, error)
	GetImageSegmentsWithNodes(ctx context.Context, in *GetImageSegmentsWithNodesIn, opts ...grpc.CallOption) (*GetImageSegmentsWithNodesOut, error)
	CreateSegment(ctx context.Context, in *CreateSegmentIn, opts ...grpc.CallOption) (*CreateSegmentOut, error)
	DeleteSegment(ctx context.Context, in *DeleteSegmentIn, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateSegment(ctx context.Context, in *UpdateSegmentIn, opts ...grpc.CallOption) (*UpdateSegmentOut, error)
	CreateNode(ctx context.Context, in *CreateNodeIn, opts ...grpc.CallOption) (*CreateNodeOut, error)
	DeleteNode(ctx context.Context, in *DeleteNodeIn, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateNode(ctx context.Context, in *UpdateNodeIn, opts ...grpc.CallOption) (*UpdateNodeOut, error)
}

type uziSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewUziSrvClient(cc grpc.ClientConnInterface) UziSrvClient {
	return &uziSrvClient{cc}
}

func (c *uziSrvClient) GetDeviceList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDeviceListOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDeviceListOut)
	err := c.cc.Invoke(ctx, UziSrv_GetDeviceList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) CreateUzi(ctx context.Context, in *CreateUziIn, opts ...grpc.CallOption) (*CreateUziOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUziOut)
	err := c.cc.Invoke(ctx, UziSrv_CreateUzi_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) UpdateUzi(ctx context.Context, in *UpdateUziIn, opts ...grpc.CallOption) (*UpdateUziOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUziOut)
	err := c.cc.Invoke(ctx, UziSrv_UpdateUzi_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) UpdateEchographic(ctx context.Context, in *UpdateEchographicIn, opts ...grpc.CallOption) (*UpdateEchographicOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateEchographicOut)
	err := c.cc.Invoke(ctx, UziSrv_UpdateEchographic_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) GetUzi(ctx context.Context, in *GetUziIn, opts ...grpc.CallOption) (*GetUziOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUziOut)
	err := c.cc.Invoke(ctx, UziSrv_GetUzi_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) GetUziImages(ctx context.Context, in *GetUziImagesIn, opts ...grpc.CallOption) (*GetUziImagesOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUziImagesOut)
	err := c.cc.Invoke(ctx, UziSrv_GetUziImages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) GetImageSegmentsWithNodes(ctx context.Context, in *GetImageSegmentsWithNodesIn, opts ...grpc.CallOption) (*GetImageSegmentsWithNodesOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetImageSegmentsWithNodesOut)
	err := c.cc.Invoke(ctx, UziSrv_GetImageSegmentsWithNodes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) CreateSegment(ctx context.Context, in *CreateSegmentIn, opts ...grpc.CallOption) (*CreateSegmentOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSegmentOut)
	err := c.cc.Invoke(ctx, UziSrv_CreateSegment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) DeleteSegment(ctx context.Context, in *DeleteSegmentIn, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UziSrv_DeleteSegment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) UpdateSegment(ctx context.Context, in *UpdateSegmentIn, opts ...grpc.CallOption) (*UpdateSegmentOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateSegmentOut)
	err := c.cc.Invoke(ctx, UziSrv_UpdateSegment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) CreateNode(ctx context.Context, in *CreateNodeIn, opts ...grpc.CallOption) (*CreateNodeOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateNodeOut)
	err := c.cc.Invoke(ctx, UziSrv_CreateNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) DeleteNode(ctx context.Context, in *DeleteNodeIn, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UziSrv_DeleteNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) UpdateNode(ctx context.Context, in *UpdateNodeIn, opts ...grpc.CallOption) (*UpdateNodeOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateNodeOut)
	err := c.cc.Invoke(ctx, UziSrv_UpdateNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UziSrvServer is the server API for UziSrv service.
// All implementations must embed UnimplementedUziSrvServer
// for forward compatibility.
type UziSrvServer interface {
	GetDeviceList(context.Context, *emptypb.Empty) (*GetDeviceListOut, error)
	CreateUzi(context.Context, *CreateUziIn) (*CreateUziOut, error)
	UpdateUzi(context.Context, *UpdateUziIn) (*UpdateUziOut, error)
	UpdateEchographic(context.Context, *UpdateEchographicIn) (*UpdateEchographicOut, error)
	GetUzi(context.Context, *GetUziIn) (*GetUziOut, error)
	GetUziImages(context.Context, *GetUziImagesIn) (*GetUziImagesOut, error)
	GetImageSegmentsWithNodes(context.Context, *GetImageSegmentsWithNodesIn) (*GetImageSegmentsWithNodesOut, error)
	CreateSegment(context.Context, *CreateSegmentIn) (*CreateSegmentOut, error)
	DeleteSegment(context.Context, *DeleteSegmentIn) (*emptypb.Empty, error)
	UpdateSegment(context.Context, *UpdateSegmentIn) (*UpdateSegmentOut, error)
	CreateNode(context.Context, *CreateNodeIn) (*CreateNodeOut, error)
	DeleteNode(context.Context, *DeleteNodeIn) (*emptypb.Empty, error)
	UpdateNode(context.Context, *UpdateNodeIn) (*UpdateNodeOut, error)
	mustEmbedUnimplementedUziSrvServer()
}

// UnimplementedUziSrvServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUziSrvServer struct{}

func (UnimplementedUziSrvServer) GetDeviceList(context.Context, *emptypb.Empty) (*GetDeviceListOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceList not implemented")
}
func (UnimplementedUziSrvServer) CreateUzi(context.Context, *CreateUziIn) (*CreateUziOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUzi not implemented")
}
func (UnimplementedUziSrvServer) UpdateUzi(context.Context, *UpdateUziIn) (*UpdateUziOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUzi not implemented")
}
func (UnimplementedUziSrvServer) UpdateEchographic(context.Context, *UpdateEchographicIn) (*UpdateEchographicOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEchographic not implemented")
}
func (UnimplementedUziSrvServer) GetUzi(context.Context, *GetUziIn) (*GetUziOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUzi not implemented")
}
func (UnimplementedUziSrvServer) GetUziImages(context.Context, *GetUziImagesIn) (*GetUziImagesOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUziImages not implemented")
}
func (UnimplementedUziSrvServer) GetImageSegmentsWithNodes(context.Context, *GetImageSegmentsWithNodesIn) (*GetImageSegmentsWithNodesOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImageSegmentsWithNodes not implemented")
}
func (UnimplementedUziSrvServer) CreateSegment(context.Context, *CreateSegmentIn) (*CreateSegmentOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSegment not implemented")
}
func (UnimplementedUziSrvServer) DeleteSegment(context.Context, *DeleteSegmentIn) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSegment not implemented")
}
func (UnimplementedUziSrvServer) UpdateSegment(context.Context, *UpdateSegmentIn) (*UpdateSegmentOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSegment not implemented")
}
func (UnimplementedUziSrvServer) CreateNode(context.Context, *CreateNodeIn) (*CreateNodeOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNode not implemented")
}
func (UnimplementedUziSrvServer) DeleteNode(context.Context, *DeleteNodeIn) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNode not implemented")
}
func (UnimplementedUziSrvServer) UpdateNode(context.Context, *UpdateNodeIn) (*UpdateNodeOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNode not implemented")
}
func (UnimplementedUziSrvServer) mustEmbedUnimplementedUziSrvServer() {}
func (UnimplementedUziSrvServer) testEmbeddedByValue()                {}

// UnsafeUziSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UziSrvServer will
// result in compilation errors.
type UnsafeUziSrvServer interface {
	mustEmbedUnimplementedUziSrvServer()
}

func RegisterUziSrvServer(s grpc.ServiceRegistrar, srv UziSrvServer) {
	// If the following call pancis, it indicates UnimplementedUziSrvServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UziSrv_ServiceDesc, srv)
}

func _UziSrv_GetDeviceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).GetDeviceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_GetDeviceList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).GetDeviceList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_CreateUzi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUziIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).CreateUzi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_CreateUzi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).CreateUzi(ctx, req.(*CreateUziIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_UpdateUzi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUziIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).UpdateUzi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_UpdateUzi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).UpdateUzi(ctx, req.(*UpdateUziIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_UpdateEchographic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEchographicIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).UpdateEchographic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_UpdateEchographic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).UpdateEchographic(ctx, req.(*UpdateEchographicIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_GetUzi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUziIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).GetUzi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_GetUzi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).GetUzi(ctx, req.(*GetUziIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_GetUziImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUziImagesIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).GetUziImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_GetUziImages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).GetUziImages(ctx, req.(*GetUziImagesIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_GetImageSegmentsWithNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageSegmentsWithNodesIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).GetImageSegmentsWithNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_GetImageSegmentsWithNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).GetImageSegmentsWithNodes(ctx, req.(*GetImageSegmentsWithNodesIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_CreateSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSegmentIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).CreateSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_CreateSegment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).CreateSegment(ctx, req.(*CreateSegmentIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_DeleteSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSegmentIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).DeleteSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_DeleteSegment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).DeleteSegment(ctx, req.(*DeleteSegmentIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_UpdateSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSegmentIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).UpdateSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_UpdateSegment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).UpdateSegment(ctx, req.(*UpdateSegmentIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_CreateNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).CreateNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_CreateNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).CreateNode(ctx, req.(*CreateNodeIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_DeleteNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodeIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).DeleteNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_DeleteNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).DeleteNode(ctx, req.(*DeleteNodeIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_UpdateNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodeIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).UpdateNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_UpdateNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).UpdateNode(ctx, req.(*UpdateNodeIn))
	}
	return interceptor(ctx, in, info, handler)
}

// UziSrv_ServiceDesc is the grpc.ServiceDesc for UziSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UziSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UziSrv",
	HandlerType: (*UziSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getDeviceList",
			Handler:    _UziSrv_GetDeviceList_Handler,
		},
		{
			MethodName: "createUzi",
			Handler:    _UziSrv_CreateUzi_Handler,
		},
		{
			MethodName: "updateUzi",
			Handler:    _UziSrv_UpdateUzi_Handler,
		},
		{
			MethodName: "updateEchographic",
			Handler:    _UziSrv_UpdateEchographic_Handler,
		},
		{
			MethodName: "getUzi",
			Handler:    _UziSrv_GetUzi_Handler,
		},
		{
			MethodName: "getUziImages",
			Handler:    _UziSrv_GetUziImages_Handler,
		},
		{
			MethodName: "getImageSegmentsWithNodes",
			Handler:    _UziSrv_GetImageSegmentsWithNodes_Handler,
		},
		{
			MethodName: "createSegment",
			Handler:    _UziSrv_CreateSegment_Handler,
		},
		{
			MethodName: "deleteSegment",
			Handler:    _UziSrv_DeleteSegment_Handler,
		},
		{
			MethodName: "updateSegment",
			Handler:    _UziSrv_UpdateSegment_Handler,
		},
		{
			MethodName: "createNode",
			Handler:    _UziSrv_CreateNode_Handler,
		},
		{
			MethodName: "deleteNode",
			Handler:    _UziSrv_DeleteNode_Handler,
		},
		{
			MethodName: "updateNode",
			Handler:    _UziSrv_UpdateNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "uzi.proto",
}
