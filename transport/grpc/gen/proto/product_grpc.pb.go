// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// ProductApiClient is the client API for ProductApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductApiClient interface {
	GetProducts(ctx context.Context, in *PaginationRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	GetProductById(ctx context.Context, in *ProductIdRequest, opts ...grpc.CallOption) (*Product, error)
}

type productApiClient struct {
	cc grpc.ClientConnInterface
}

func NewProductApiClient(cc grpc.ClientConnInterface) ProductApiClient {
	return &productApiClient{cc}
}

func (c *productApiClient) GetProducts(ctx context.Context, in *PaginationRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/main.ProductApi/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productApiClient) GetProductById(ctx context.Context, in *ProductIdRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/main.ProductApi/GetProductById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductApiServer is the server API for ProductApi service.
// All implementations must embed UnimplementedProductApiServer
// for forward compatibility
type ProductApiServer interface {
	GetProducts(context.Context, *PaginationRequest) (*ProductResponse, error)
	GetProductById(context.Context, *ProductIdRequest) (*Product, error)
	mustEmbedUnimplementedProductApiServer()
}

// UnimplementedProductApiServer must be embedded to have forward compatible implementations.
type UnimplementedProductApiServer struct {
}

func (UnimplementedProductApiServer) GetProducts(context.Context, *PaginationRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (UnimplementedProductApiServer) GetProductById(context.Context, *ProductIdRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductById not implemented")
}
func (UnimplementedProductApiServer) mustEmbedUnimplementedProductApiServer() {}

// UnsafeProductApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductApiServer will
// result in compilation errors.
type UnsafeProductApiServer interface {
	mustEmbedUnimplementedProductApiServer()
}

func RegisterProductApiServer(s grpc.ServiceRegistrar, srv ProductApiServer) {
	s.RegisterService(&ProductApi_ServiceDesc, srv)
}

func _ProductApi_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaginationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductApiServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ProductApi/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductApiServer).GetProducts(ctx, req.(*PaginationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductApi_GetProductById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductApiServer).GetProductById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ProductApi/GetProductById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductApiServer).GetProductById(ctx, req.(*ProductIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductApi_ServiceDesc is the grpc.ServiceDesc for ProductApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.ProductApi",
	HandlerType: (*ProductApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProducts",
			Handler:    _ProductApi_GetProducts_Handler,
		},
		{
			MethodName: "GetProductById",
			Handler:    _ProductApi_GetProductById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
