// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package patient

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MedPatientClient is the client API for MedPatient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MedPatientClient interface {
	// Добавить пациента
	//
	// Принимает данные пациента и его карты.
	AddPatient(ctx context.Context, in *CreatePatientRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Получить список пациентов
	//
	// Принимает пагинацию (Пока без нее). Возвращает список пациентов
	GetPatientList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PatientListResponse, error)
	// Получить пациента и его карту
	//
	// Получает id пациента. Возвращает пациента и его карту
	GetPatientInfoByID(ctx context.Context, in *PatientInfoRequest, opts ...grpc.CallOption) (*PatientInfoResponse, error)
	// UNIMPLEMENTED!!!
	//
	// Получает id пациента. Возвращает пациента, карту, снимки
	// rpc PatientShots (google.protobuf.Empty) returns (google.protobuf.Empty) {
	//   option (google.api.http) = {
	//     get: "/med/patient/shots"
	//   };
	// }
	// Обновить данные пациента
	//
	// Получает пациента, карту. Возвращает id
	UpdatePatient(ctx context.Context, in *PatientUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type medPatientClient struct {
	cc grpc.ClientConnInterface
}

func NewMedPatientClient(cc grpc.ClientConnInterface) MedPatientClient {
	return &medPatientClient{cc}
}

func (c *medPatientClient) AddPatient(ctx context.Context, in *CreatePatientRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/yir.med.MedPatient/AddPatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medPatientClient) GetPatientList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PatientListResponse, error) {
	out := new(PatientListResponse)
	err := c.cc.Invoke(ctx, "/yir.med.MedPatient/GetPatientList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medPatientClient) GetPatientInfoByID(ctx context.Context, in *PatientInfoRequest, opts ...grpc.CallOption) (*PatientInfoResponse, error) {
	out := new(PatientInfoResponse)
	err := c.cc.Invoke(ctx, "/yir.med.MedPatient/GetPatientInfoByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medPatientClient) UpdatePatient(ctx context.Context, in *PatientUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/yir.med.MedPatient/UpdatePatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MedPatientServer is the server API for MedPatient service.
// All implementations must embed UnimplementedMedPatientServer
// for forward compatibility
type MedPatientServer interface {
	// Добавить пациента
	//
	// Принимает данные пациента и его карты.
	AddPatient(context.Context, *CreatePatientRequest) (*emptypb.Empty, error)
	// Получить список пациентов
	//
	// Принимает пагинацию (Пока без нее). Возвращает список пациентов
	GetPatientList(context.Context, *emptypb.Empty) (*PatientListResponse, error)
	// Получить пациента и его карту
	//
	// Получает id пациента. Возвращает пациента и его карту
	GetPatientInfoByID(context.Context, *PatientInfoRequest) (*PatientInfoResponse, error)
	// UNIMPLEMENTED!!!
	//
	// Получает id пациента. Возвращает пациента, карту, снимки
	// rpc PatientShots (google.protobuf.Empty) returns (google.protobuf.Empty) {
	//   option (google.api.http) = {
	//     get: "/med/patient/shots"
	//   };
	// }
	// Обновить данные пациента
	//
	// Получает пациента, карту. Возвращает id
	UpdatePatient(context.Context, *PatientUpdateRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMedPatientServer()
}

// UnimplementedMedPatientServer must be embedded to have forward compatible implementations.
type UnimplementedMedPatientServer struct {
}

func (UnimplementedMedPatientServer) AddPatient(context.Context, *CreatePatientRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPatient not implemented")
}
func (UnimplementedMedPatientServer) GetPatientList(context.Context, *emptypb.Empty) (*PatientListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPatientList not implemented")
}
func (UnimplementedMedPatientServer) GetPatientInfoByID(context.Context, *PatientInfoRequest) (*PatientInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPatientInfoByID not implemented")
}
func (UnimplementedMedPatientServer) UpdatePatient(context.Context, *PatientUpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePatient not implemented")
}
func (UnimplementedMedPatientServer) mustEmbedUnimplementedMedPatientServer() {}

// UnsafeMedPatientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MedPatientServer will
// result in compilation errors.
type UnsafeMedPatientServer interface {
	mustEmbedUnimplementedMedPatientServer()
}

func RegisterMedPatientServer(s grpc.ServiceRegistrar, srv MedPatientServer) {
	s.RegisterService(&MedPatient_ServiceDesc, srv)
}

func _MedPatient_AddPatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedPatientServer).AddPatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yir.med.MedPatient/AddPatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedPatientServer).AddPatient(ctx, req.(*CreatePatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedPatient_GetPatientList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedPatientServer).GetPatientList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yir.med.MedPatient/GetPatientList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedPatientServer).GetPatientList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedPatient_GetPatientInfoByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatientInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedPatientServer).GetPatientInfoByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yir.med.MedPatient/GetPatientInfoByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedPatientServer).GetPatientInfoByID(ctx, req.(*PatientInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedPatient_UpdatePatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatientUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedPatientServer).UpdatePatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yir.med.MedPatient/UpdatePatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedPatientServer).UpdatePatient(ctx, req.(*PatientUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MedPatient_ServiceDesc is the grpc.ServiceDesc for MedPatient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MedPatient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yir.med.MedPatient",
	HandlerType: (*MedPatientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPatient",
			Handler:    _MedPatient_AddPatient_Handler,
		},
		{
			MethodName: "GetPatientList",
			Handler:    _MedPatient_GetPatientList_Handler,
		},
		{
			MethodName: "GetPatientInfoByID",
			Handler:    _MedPatient_GetPatientInfoByID_Handler,
		},
		{
			MethodName: "UpdatePatient",
			Handler:    _MedPatient_UpdatePatient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "patient.proto",
}
