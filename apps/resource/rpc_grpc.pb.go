// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: apps/resource/pb/rpc.proto

package resource

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

const (
	Rpc_CreateResource_FullMethodName = "/cmdb.resource.Rpc/CreateResource"
	Rpc_DescResource_FullMethodName   = "/cmdb.resource.Rpc/DescResource"
	Rpc_QueryResource_FullMethodName  = "/cmdb.resource.Rpc/QueryResource"
)

// RpcClient is the client API for Rpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcClient interface {
	CreateResource(ctx context.Context, in *Spec, opts ...grpc.CallOption) (*Resource, error)
	DescResource(ctx context.Context, in *DescResourceRequest, opts ...grpc.CallOption) (*Resource, error)
	QueryResource(ctx context.Context, in *QueryResourceRequest, opts ...grpc.CallOption) (*ResourceSet, error)
}

type rpcClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcClient(cc grpc.ClientConnInterface) RpcClient {
	return &rpcClient{cc}
}

func (c *rpcClient) CreateResource(ctx context.Context, in *Spec, opts ...grpc.CallOption) (*Resource, error) {
	out := new(Resource)
	err := c.cc.Invoke(ctx, Rpc_CreateResource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) DescResource(ctx context.Context, in *DescResourceRequest, opts ...grpc.CallOption) (*Resource, error) {
	out := new(Resource)
	err := c.cc.Invoke(ctx, Rpc_DescResource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) QueryResource(ctx context.Context, in *QueryResourceRequest, opts ...grpc.CallOption) (*ResourceSet, error) {
	out := new(ResourceSet)
	err := c.cc.Invoke(ctx, Rpc_QueryResource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcServer is the server API for Rpc service.
// All implementations must embed UnimplementedRpcServer
// for forward compatibility
type RpcServer interface {
	CreateResource(context.Context, *Spec) (*Resource, error)
	DescResource(context.Context, *DescResourceRequest) (*Resource, error)
	QueryResource(context.Context, *QueryResourceRequest) (*ResourceSet, error)
	mustEmbedUnimplementedRpcServer()
}

// UnimplementedRpcServer must be embedded to have forward compatible implementations.
type UnimplementedRpcServer struct {
}

func (UnimplementedRpcServer) CreateResource(context.Context, *Spec) (*Resource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResource not implemented")
}
func (UnimplementedRpcServer) DescResource(context.Context, *DescResourceRequest) (*Resource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescResource not implemented")
}
func (UnimplementedRpcServer) QueryResource(context.Context, *QueryResourceRequest) (*ResourceSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryResource not implemented")
}
func (UnimplementedRpcServer) mustEmbedUnimplementedRpcServer() {}

// UnsafeRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcServer will
// result in compilation errors.
type UnsafeRpcServer interface {
	mustEmbedUnimplementedRpcServer()
}

func RegisterRpcServer(s grpc.ServiceRegistrar, srv RpcServer) {
	s.RegisterService(&Rpc_ServiceDesc, srv)
}

func _Rpc_CreateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Spec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).CreateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_CreateResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).CreateResource(ctx, req.(*Spec))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_DescResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).DescResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_DescResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).DescResource(ctx, req.(*DescResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_QueryResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).QueryResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_QueryResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).QueryResource(ctx, req.(*QueryResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rpc_ServiceDesc is the grpc.ServiceDesc for Rpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cmdb.resource.Rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateResource",
			Handler:    _Rpc_CreateResource_Handler,
		},
		{
			MethodName: "DescResource",
			Handler:    _Rpc_DescResource_Handler,
		},
		{
			MethodName: "QueryResource",
			Handler:    _Rpc_QueryResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/resource/pb/rpc.proto",
}
