// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: category.proto

package product_service

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

// CategoriesServiceClient is the client API for CategoriesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoriesServiceClient interface {
	Create(ctx context.Context, in *CreateCategoriesRequest, opts ...grpc.CallOption) (*IdResponse, error)
	Get(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*GetCategoriesResponse, error)
	GetAll(ctx context.Context, in *GetAllCategoriesRequest, opts ...grpc.CallOption) (*GetAllCategoriesResponse, error)
	Update(ctx context.Context, in *UpdateCategoriesRequest, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Response, error)
}

type categoriesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoriesServiceClient(cc grpc.ClientConnInterface) CategoriesServiceClient {
	return &categoriesServiceClient{cc}
}

func (c *categoriesServiceClient) Create(ctx context.Context, in *CreateCategoriesRequest, opts ...grpc.CallOption) (*IdResponse, error) {
	out := new(IdResponse)
	err := c.cc.Invoke(ctx, "/product_service.CategoriesService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoriesServiceClient) Get(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*GetCategoriesResponse, error) {
	out := new(GetCategoriesResponse)
	err := c.cc.Invoke(ctx, "/product_service.CategoriesService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoriesServiceClient) GetAll(ctx context.Context, in *GetAllCategoriesRequest, opts ...grpc.CallOption) (*GetAllCategoriesResponse, error) {
	out := new(GetAllCategoriesResponse)
	err := c.cc.Invoke(ctx, "/product_service.CategoriesService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoriesServiceClient) Update(ctx context.Context, in *UpdateCategoriesRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/product_service.CategoriesService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoriesServiceClient) Delete(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/product_service.CategoriesService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoriesServiceServer is the server API for CategoriesService service.
// All implementations must embed UnimplementedCategoriesServiceServer
// for forward compatibility
type CategoriesServiceServer interface {
	Create(context.Context, *CreateCategoriesRequest) (*IdResponse, error)
	Get(context.Context, *IdRequest) (*GetCategoriesResponse, error)
	GetAll(context.Context, *GetAllCategoriesRequest) (*GetAllCategoriesResponse, error)
	Update(context.Context, *UpdateCategoriesRequest) (*Response, error)
	Delete(context.Context, *IdRequest) (*Response, error)
	mustEmbedUnimplementedCategoriesServiceServer()
}

// UnimplementedCategoriesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCategoriesServiceServer struct {
}

func (UnimplementedCategoriesServiceServer) Create(context.Context, *CreateCategoriesRequest) (*IdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCategoriesServiceServer) Get(context.Context, *IdRequest) (*GetCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCategoriesServiceServer) GetAll(context.Context, *GetAllCategoriesRequest) (*GetAllCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedCategoriesServiceServer) Update(context.Context, *UpdateCategoriesRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCategoriesServiceServer) Delete(context.Context, *IdRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCategoriesServiceServer) mustEmbedUnimplementedCategoriesServiceServer() {}

// UnsafeCategoriesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoriesServiceServer will
// result in compilation errors.
type UnsafeCategoriesServiceServer interface {
	mustEmbedUnimplementedCategoriesServiceServer()
}

func RegisterCategoriesServiceServer(s grpc.ServiceRegistrar, srv CategoriesServiceServer) {
	s.RegisterService(&CategoriesService_ServiceDesc, srv)
}

func _CategoriesService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoriesServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.CategoriesService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoriesServiceServer).Create(ctx, req.(*CreateCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoriesService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoriesServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.CategoriesService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoriesServiceServer).Get(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoriesService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoriesServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.CategoriesService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoriesServiceServer).GetAll(ctx, req.(*GetAllCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoriesService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoriesServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.CategoriesService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoriesServiceServer).Update(ctx, req.(*UpdateCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoriesService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoriesServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.CategoriesService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoriesServiceServer).Delete(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoriesService_ServiceDesc is the grpc.ServiceDesc for CategoriesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoriesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product_service.CategoriesService",
	HandlerType: (*CategoriesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CategoriesService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CategoriesService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _CategoriesService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CategoriesService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CategoriesService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "category.proto",
}
