// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: staff_transaction.proto

package sale_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StaffTransactionServiceClient is the client API for StaffTransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StaffTransactionServiceClient interface {
	Create(ctx context.Context, in *CreateStaffTransactionRequest, opts ...grpc.CallOption) (*IdResponse, error)
	Get(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*GetStaffTransactionResponse, error)
	GetAll(ctx context.Context, in *GetAllStaffTransactionRequest, opts ...grpc.CallOption) (*GetAllStaffTransactionResponse, error)
	Update(ctx context.Context, in *UpdateStaffTransactionRequest, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Response, error)
}

type staffTransactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStaffTransactionServiceClient(cc grpc.ClientConnInterface) StaffTransactionServiceClient {
	return &staffTransactionServiceClient{cc}
}

func (c *staffTransactionServiceClient) Create(ctx context.Context, in *CreateStaffTransactionRequest, opts ...grpc.CallOption) (*IdResponse, error) {
	out := new(IdResponse)
	err := c.cc.Invoke(ctx, "/sale_service.StaffTransactionService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffTransactionServiceClient) Get(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*GetStaffTransactionResponse, error) {
	out := new(GetStaffTransactionResponse)
	err := c.cc.Invoke(ctx, "/sale_service.StaffTransactionService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffTransactionServiceClient) GetAll(ctx context.Context, in *GetAllStaffTransactionRequest, opts ...grpc.CallOption) (*GetAllStaffTransactionResponse, error) {
	out := new(GetAllStaffTransactionResponse)
	err := c.cc.Invoke(ctx, "/sale_service.StaffTransactionService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffTransactionServiceClient) Update(ctx context.Context, in *UpdateStaffTransactionRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/sale_service.StaffTransactionService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffTransactionServiceClient) Delete(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/sale_service.StaffTransactionService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaffTransactionServiceServer is the server API for StaffTransactionService service.
// All implementations must embed UnimplementedStaffTransactionServiceServer
// for forward compatibility
type StaffTransactionServiceServer interface {
	Create(context.Context, *CreateStaffTransactionRequest) (*IdResponse, error)
	Get(context.Context, *IdRequest) (*GetStaffTransactionResponse, error)
	GetAll(context.Context, *GetAllStaffTransactionRequest) (*GetAllStaffTransactionResponse, error)
	Update(context.Context, *UpdateStaffTransactionRequest) (*Response, error)
	Delete(context.Context, *IdRequest) (*Response, error)
	mustEmbedUnimplementedStaffTransactionServiceServer()
}

// UnimplementedStaffTransactionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStaffTransactionServiceServer struct {
}

func (UnimplementedStaffTransactionServiceServer) Create(context.Context, *CreateStaffTransactionRequest) (*IdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedStaffTransactionServiceServer) Get(context.Context, *IdRequest) (*GetStaffTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStaffTransactionServiceServer) GetAll(context.Context, *GetAllStaffTransactionRequest) (*GetAllStaffTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedStaffTransactionServiceServer) Update(context.Context, *UpdateStaffTransactionRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStaffTransactionServiceServer) Delete(context.Context, *IdRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStaffTransactionServiceServer) mustEmbedUnimplementedStaffTransactionServiceServer() {
}

// UnsafeStaffTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StaffTransactionServiceServer will
// result in compilation errors.
type UnsafeStaffTransactionServiceServer interface {
	mustEmbedUnimplementedStaffTransactionServiceServer()
}

func RegisterStaffTransactionServiceServer(s grpc.ServiceRegistrar, srv StaffTransactionServiceServer) {
	s.RegisterService(&StaffTransactionService_ServiceDesc, srv)
}

func _StaffTransactionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStaffTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffTransactionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.StaffTransactionService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffTransactionServiceServer).Create(ctx, req.(*CreateStaffTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffTransactionService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffTransactionServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.StaffTransactionService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffTransactionServiceServer).Get(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffTransactionService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllStaffTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffTransactionServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.StaffTransactionService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffTransactionServiceServer).GetAll(ctx, req.(*GetAllStaffTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffTransactionService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStaffTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffTransactionServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.StaffTransactionService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffTransactionServiceServer).Update(ctx, req.(*UpdateStaffTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffTransactionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffTransactionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.StaffTransactionService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffTransactionServiceServer).Delete(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StaffTransactionService_ServiceDesc is the grpc.ServiceDesc for StaffTransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StaffTransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sale_service.StaffTransactionService",
	HandlerType: (*StaffTransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _StaffTransactionService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _StaffTransactionService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _StaffTransactionService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StaffTransactionService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _StaffTransactionService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staff_transaction.proto",
}
