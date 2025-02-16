// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: service.proto

package service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MedSrv_RegisterDoctor_FullMethodName    = "/MedSrv/registerDoctor"
	MedSrv_GetDoctor_FullMethodName         = "/MedSrv/getDoctor"
	MedSrv_UpdateDoctor_FullMethodName      = "/MedSrv/updateDoctor"
	MedSrv_CreatePatient_FullMethodName     = "/MedSrv/createPatient"
	MedSrv_GetPatient_FullMethodName        = "/MedSrv/getPatient"
	MedSrv_GetDoctorPatients_FullMethodName = "/MedSrv/getDoctorPatients"
	MedSrv_UpdatePatient_FullMethodName     = "/MedSrv/updatePatient"
	MedSrv_CreateCard_FullMethodName        = "/MedSrv/createCard"
	MedSrv_GetCard_FullMethodName           = "/MedSrv/getCard"
	MedSrv_UpdateCard_FullMethodName        = "/MedSrv/updateCard"
)

// MedSrvClient is the client API for MedSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MedSrvClient interface {
	RegisterDoctor(ctx context.Context, in *RegisterDoctorIn, opts ...grpc.CallOption) (*empty.Empty, error)
	GetDoctor(ctx context.Context, in *GetDoctorIn, opts ...grpc.CallOption) (*GetDoctorOut, error)
	UpdateDoctor(ctx context.Context, in *UpdateDoctorIn, opts ...grpc.CallOption) (*UpdateDoctorOut, error)
	CreatePatient(ctx context.Context, in *CreatePatientIn, opts ...grpc.CallOption) (*CreatePatientOut, error)
	GetPatient(ctx context.Context, in *GetPatientIn, opts ...grpc.CallOption) (*GetPatientOut, error)
	GetDoctorPatients(ctx context.Context, in *GetDoctorPatientsIn, opts ...grpc.CallOption) (*GetDoctorPatientsOut, error)
	UpdatePatient(ctx context.Context, in *UpdatePatientIn, opts ...grpc.CallOption) (*UpdatePatientOut, error)
	CreateCard(ctx context.Context, in *CreateCardIn, opts ...grpc.CallOption) (*empty.Empty, error)
	GetCard(ctx context.Context, in *GetCardIn, opts ...grpc.CallOption) (*GetCardOut, error)
	UpdateCard(ctx context.Context, in *UpdateCardIn, opts ...grpc.CallOption) (*UpdateCardOut, error)
}

type medSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewMedSrvClient(cc grpc.ClientConnInterface) MedSrvClient {
	return &medSrvClient{cc}
}

func (c *medSrvClient) RegisterDoctor(ctx context.Context, in *RegisterDoctorIn, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, MedSrv_RegisterDoctor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medSrvClient) GetDoctor(ctx context.Context, in *GetDoctorIn, opts ...grpc.CallOption) (*GetDoctorOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDoctorOut)
	err := c.cc.Invoke(ctx, MedSrv_GetDoctor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medSrvClient) UpdateDoctor(ctx context.Context, in *UpdateDoctorIn, opts ...grpc.CallOption) (*UpdateDoctorOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateDoctorOut)
	err := c.cc.Invoke(ctx, MedSrv_UpdateDoctor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medSrvClient) CreatePatient(ctx context.Context, in *CreatePatientIn, opts ...grpc.CallOption) (*CreatePatientOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePatientOut)
	err := c.cc.Invoke(ctx, MedSrv_CreatePatient_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medSrvClient) GetPatient(ctx context.Context, in *GetPatientIn, opts ...grpc.CallOption) (*GetPatientOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPatientOut)
	err := c.cc.Invoke(ctx, MedSrv_GetPatient_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medSrvClient) GetDoctorPatients(ctx context.Context, in *GetDoctorPatientsIn, opts ...grpc.CallOption) (*GetDoctorPatientsOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDoctorPatientsOut)
	err := c.cc.Invoke(ctx, MedSrv_GetDoctorPatients_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medSrvClient) UpdatePatient(ctx context.Context, in *UpdatePatientIn, opts ...grpc.CallOption) (*UpdatePatientOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePatientOut)
	err := c.cc.Invoke(ctx, MedSrv_UpdatePatient_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medSrvClient) CreateCard(ctx context.Context, in *CreateCardIn, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, MedSrv_CreateCard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medSrvClient) GetCard(ctx context.Context, in *GetCardIn, opts ...grpc.CallOption) (*GetCardOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCardOut)
	err := c.cc.Invoke(ctx, MedSrv_GetCard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medSrvClient) UpdateCard(ctx context.Context, in *UpdateCardIn, opts ...grpc.CallOption) (*UpdateCardOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCardOut)
	err := c.cc.Invoke(ctx, MedSrv_UpdateCard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MedSrvServer is the server API for MedSrv service.
// All implementations must embed UnimplementedMedSrvServer
// for forward compatibility.
type MedSrvServer interface {
	RegisterDoctor(context.Context, *RegisterDoctorIn) (*empty.Empty, error)
	GetDoctor(context.Context, *GetDoctorIn) (*GetDoctorOut, error)
	UpdateDoctor(context.Context, *UpdateDoctorIn) (*UpdateDoctorOut, error)
	CreatePatient(context.Context, *CreatePatientIn) (*CreatePatientOut, error)
	GetPatient(context.Context, *GetPatientIn) (*GetPatientOut, error)
	GetDoctorPatients(context.Context, *GetDoctorPatientsIn) (*GetDoctorPatientsOut, error)
	UpdatePatient(context.Context, *UpdatePatientIn) (*UpdatePatientOut, error)
	CreateCard(context.Context, *CreateCardIn) (*empty.Empty, error)
	GetCard(context.Context, *GetCardIn) (*GetCardOut, error)
	UpdateCard(context.Context, *UpdateCardIn) (*UpdateCardOut, error)
	mustEmbedUnimplementedMedSrvServer()
}

// UnimplementedMedSrvServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMedSrvServer struct{}

func (UnimplementedMedSrvServer) RegisterDoctor(context.Context, *RegisterDoctorIn) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDoctor not implemented")
}
func (UnimplementedMedSrvServer) GetDoctor(context.Context, *GetDoctorIn) (*GetDoctorOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDoctor not implemented")
}
func (UnimplementedMedSrvServer) UpdateDoctor(context.Context, *UpdateDoctorIn) (*UpdateDoctorOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDoctor not implemented")
}
func (UnimplementedMedSrvServer) CreatePatient(context.Context, *CreatePatientIn) (*CreatePatientOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePatient not implemented")
}
func (UnimplementedMedSrvServer) GetPatient(context.Context, *GetPatientIn) (*GetPatientOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPatient not implemented")
}
func (UnimplementedMedSrvServer) GetDoctorPatients(context.Context, *GetDoctorPatientsIn) (*GetDoctorPatientsOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDoctorPatients not implemented")
}
func (UnimplementedMedSrvServer) UpdatePatient(context.Context, *UpdatePatientIn) (*UpdatePatientOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePatient not implemented")
}
func (UnimplementedMedSrvServer) CreateCard(context.Context, *CreateCardIn) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCard not implemented")
}
func (UnimplementedMedSrvServer) GetCard(context.Context, *GetCardIn) (*GetCardOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCard not implemented")
}
func (UnimplementedMedSrvServer) UpdateCard(context.Context, *UpdateCardIn) (*UpdateCardOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCard not implemented")
}
func (UnimplementedMedSrvServer) mustEmbedUnimplementedMedSrvServer() {}
func (UnimplementedMedSrvServer) testEmbeddedByValue()                {}

// UnsafeMedSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MedSrvServer will
// result in compilation errors.
type UnsafeMedSrvServer interface {
	mustEmbedUnimplementedMedSrvServer()
}

func RegisterMedSrvServer(s grpc.ServiceRegistrar, srv MedSrvServer) {
	// If the following call pancis, it indicates UnimplementedMedSrvServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MedSrv_ServiceDesc, srv)
}

func _MedSrv_RegisterDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDoctorIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedSrvServer).RegisterDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedSrv_RegisterDoctor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedSrvServer).RegisterDoctor(ctx, req.(*RegisterDoctorIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedSrv_GetDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDoctorIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedSrvServer).GetDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedSrv_GetDoctor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedSrvServer).GetDoctor(ctx, req.(*GetDoctorIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedSrv_UpdateDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDoctorIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedSrvServer).UpdateDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedSrv_UpdateDoctor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedSrvServer).UpdateDoctor(ctx, req.(*UpdateDoctorIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedSrv_CreatePatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePatientIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedSrvServer).CreatePatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedSrv_CreatePatient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedSrvServer).CreatePatient(ctx, req.(*CreatePatientIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedSrv_GetPatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPatientIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedSrvServer).GetPatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedSrv_GetPatient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedSrvServer).GetPatient(ctx, req.(*GetPatientIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedSrv_GetDoctorPatients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDoctorPatientsIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedSrvServer).GetDoctorPatients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedSrv_GetDoctorPatients_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedSrvServer).GetDoctorPatients(ctx, req.(*GetDoctorPatientsIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedSrv_UpdatePatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePatientIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedSrvServer).UpdatePatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedSrv_UpdatePatient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedSrvServer).UpdatePatient(ctx, req.(*UpdatePatientIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedSrv_CreateCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCardIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedSrvServer).CreateCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedSrv_CreateCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedSrvServer).CreateCard(ctx, req.(*CreateCardIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedSrv_GetCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCardIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedSrvServer).GetCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedSrv_GetCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedSrvServer).GetCard(ctx, req.(*GetCardIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedSrv_UpdateCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCardIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedSrvServer).UpdateCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedSrv_UpdateCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedSrvServer).UpdateCard(ctx, req.(*UpdateCardIn))
	}
	return interceptor(ctx, in, info, handler)
}

// MedSrv_ServiceDesc is the grpc.ServiceDesc for MedSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MedSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MedSrv",
	HandlerType: (*MedSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "registerDoctor",
			Handler:    _MedSrv_RegisterDoctor_Handler,
		},
		{
			MethodName: "getDoctor",
			Handler:    _MedSrv_GetDoctor_Handler,
		},
		{
			MethodName: "updateDoctor",
			Handler:    _MedSrv_UpdateDoctor_Handler,
		},
		{
			MethodName: "createPatient",
			Handler:    _MedSrv_CreatePatient_Handler,
		},
		{
			MethodName: "getPatient",
			Handler:    _MedSrv_GetPatient_Handler,
		},
		{
			MethodName: "getDoctorPatients",
			Handler:    _MedSrv_GetDoctorPatients_Handler,
		},
		{
			MethodName: "updatePatient",
			Handler:    _MedSrv_UpdatePatient_Handler,
		},
		{
			MethodName: "createCard",
			Handler:    _MedSrv_CreateCard_Handler,
		},
		{
			MethodName: "getCard",
			Handler:    _MedSrv_GetCard_Handler,
		},
		{
			MethodName: "updateCard",
			Handler:    _MedSrv_UpdateCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
