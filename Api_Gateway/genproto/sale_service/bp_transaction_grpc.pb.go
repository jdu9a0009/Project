// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: bp_transaction.proto

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

// BpTransactionServiceClient is the client API for BpTransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BpTransactionServiceClient interface {
	Create(ctx context.Context, in *CreateBpTransactionRequest, opts ...grpc.CallOption) (*IdResponse, error)
	Get(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*GetBpTransactionResponse, error)
	GetAll(ctx context.Context, in *GetAllBpTransactionRequest, opts ...grpc.CallOption) (*GetAllBpTransactionResponse, error)
	Update(ctx context.Context, in *UpdateBpTransactionRequest, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Response, error)
}

type bpTransactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBpTransactionServiceClient(cc grpc.ClientConnInterface) BpTransactionServiceClient {
	return &bpTransactionServiceClient{cc}
}

func (c *bpTransactionServiceClient) Create(ctx context.Context, in *CreateBpTransactionRequest, opts ...grpc.CallOption) (*IdResponse, error) {
	out := new(IdResponse)
	err := c.cc.Invoke(ctx, "/sale_service.BpTransactionService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bpTransactionServiceClient) Get(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*GetBpTransactionResponse, error) {
	out := new(GetBpTransactionResponse)
	err := c.cc.Invoke(ctx, "/sale_service.BpTransactionService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bpTransactionServiceClient) GetAll(ctx context.Context, in *GetAllBpTransactionRequest, opts ...grpc.CallOption) (*GetAllBpTransactionResponse, error) {
	out := new(GetAllBpTransactionResponse)
	err := c.cc.Invoke(ctx, "/sale_service.BpTransactionService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bpTransactionServiceClient) Update(ctx context.Context, in *UpdateBpTransactionRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/sale_service.BpTransactionService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bpTransactionServiceClient) Delete(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/sale_service.BpTransactionService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BpTransactionServiceServer is the server API for BpTransactionService service.
// All implementations must embed UnimplementedBpTransactionServiceServer
// for forward compatibility
type BpTransactionServiceServer interface {
	Create(context.Context, *CreateBpTransactionRequest) (*IdResponse, error)
	Get(context.Context, *IdRequest) (*GetBpTransactionResponse, error)
	GetAll(context.Context, *GetAllBpTransactionRequest) (*GetAllBpTransactionResponse, error)
	Update(context.Context, *UpdateBpTransactionRequest) (*Response, error)
	Delete(context.Context, *IdRequest) (*Response, error)
	mustEmbedUnimplementedBpTransactionServiceServer()
}

// UnimplementedBpTransactionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBpTransactionServiceServer struct {
}

func (UnimplementedBpTransactionServiceServer) Create(context.Context, *CreateBpTransactionRequest) (*IdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBpTransactionServiceServer) Get(context.Context, *IdRequest) (*GetBpTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBpTransactionServiceServer) GetAll(context.Context, *GetAllBpTransactionRequest) (*GetAllBpTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedBpTransactionServiceServer) Update(context.Context, *UpdateBpTransactionRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBpTransactionServiceServer) Delete(context.Context, *IdRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBpTransactionServiceServer) mustEmbedUnimplementedBpTransactionServiceServer() {}

// UnsafeBpTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BpTransactionServiceServer will
// result in compilation errors.
type UnsafeBpTransactionServiceServer interface {
	mustEmbedUnimplementedBpTransactionServiceServer()
}

func RegisterBpTransactionServiceServer(s grpc.ServiceRegistrar, srv BpTransactionServiceServer) {
	s.RegisterService(&BpTransactionService_ServiceDesc, srv)
}

func _BpTransactionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBpTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BpTransactionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.BpTransactionService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BpTransactionServiceServer).Create(ctx, req.(*CreateBpTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BpTransactionService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BpTransactionServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.BpTransactionService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BpTransactionServiceServer).Get(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BpTransactionService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllBpTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BpTransactionServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.BpTransactionService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BpTransactionServiceServer).GetAll(ctx, req.(*GetAllBpTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BpTransactionService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBpTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BpTransactionServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.BpTransactionService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BpTransactionServiceServer).Update(ctx, req.(*UpdateBpTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BpTransactionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BpTransactionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.BpTransactionService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BpTransactionServiceServer).Delete(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BpTransactionService_ServiceDesc is the grpc.ServiceDesc for BpTransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BpTransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sale_service.BpTransactionService",
	HandlerType: (*BpTransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BpTransactionService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _BpTransactionService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _BpTransactionService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BpTransactionService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BpTransactionService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bp_transaction.proto",
}