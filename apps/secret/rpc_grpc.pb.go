// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: apps/secret/pb/rpc.proto

package secret

import (
	context "context"
	resource "github.com/qiaogy91/cmdb/apps/resource"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Rpc_CreateSecret_FullMethodName = "/cmdb.secret.Rpc/CreateSecret"
	Rpc_DeleteSecret_FullMethodName = "/cmdb.secret.Rpc/DeleteSecret"
	Rpc_QuerySecret_FullMethodName  = "/cmdb.secret.Rpc/QuerySecret"
	Rpc_DescSecret_FullMethodName   = "/cmdb.secret.Rpc/DescSecret"
	Rpc_SyncResource_FullMethodName = "/cmdb.secret.Rpc/SyncResource"
)

// RpcClient is the client API for Rpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcClient interface {
	CreateSecret(ctx context.Context, in *Spec, opts ...grpc.CallOption) (*Secret, error)
	DeleteSecret(ctx context.Context, in *DeleteSecretRequest, opts ...grpc.CallOption) (*Secret, error)
	QuerySecret(ctx context.Context, in *QuerySecretRequest, opts ...grpc.CallOption) (*SecretSet, error)
	DescSecret(ctx context.Context, in *DescSecretRequest, opts ...grpc.CallOption) (*Secret, error)
	SyncResource(ctx context.Context, in *SyncResourceRequest, opts ...grpc.CallOption) (Rpc_SyncResourceClient, error)
}

type rpcClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcClient(cc grpc.ClientConnInterface) RpcClient {
	return &rpcClient{cc}
}

func (c *rpcClient) CreateSecret(ctx context.Context, in *Spec, opts ...grpc.CallOption) (*Secret, error) {
	out := new(Secret)
	err := c.cc.Invoke(ctx, Rpc_CreateSecret_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) DeleteSecret(ctx context.Context, in *DeleteSecretRequest, opts ...grpc.CallOption) (*Secret, error) {
	out := new(Secret)
	err := c.cc.Invoke(ctx, Rpc_DeleteSecret_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) QuerySecret(ctx context.Context, in *QuerySecretRequest, opts ...grpc.CallOption) (*SecretSet, error) {
	out := new(SecretSet)
	err := c.cc.Invoke(ctx, Rpc_QuerySecret_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) DescSecret(ctx context.Context, in *DescSecretRequest, opts ...grpc.CallOption) (*Secret, error) {
	out := new(Secret)
	err := c.cc.Invoke(ctx, Rpc_DescSecret_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) SyncResource(ctx context.Context, in *SyncResourceRequest, opts ...grpc.CallOption) (Rpc_SyncResourceClient, error) {
	stream, err := c.cc.NewStream(ctx, &Rpc_ServiceDesc.Streams[0], Rpc_SyncResource_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &rpcSyncResourceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Rpc_SyncResourceClient interface {
	Recv() (*resource.Resource, error)
	grpc.ClientStream
}

type rpcSyncResourceClient struct {
	grpc.ClientStream
}

func (x *rpcSyncResourceClient) Recv() (*resource.Resource, error) {
	m := new(resource.Resource)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RpcServer is the server API for Rpc service.
// All implementations must embed UnimplementedRpcServer
// for forward compatibility
type RpcServer interface {
	CreateSecret(context.Context, *Spec) (*Secret, error)
	DeleteSecret(context.Context, *DeleteSecretRequest) (*Secret, error)
	QuerySecret(context.Context, *QuerySecretRequest) (*SecretSet, error)
	DescSecret(context.Context, *DescSecretRequest) (*Secret, error)
	SyncResource(*SyncResourceRequest, Rpc_SyncResourceServer) error
	mustEmbedUnimplementedRpcServer()
}

// UnimplementedRpcServer must be embedded to have forward compatible implementations.
type UnimplementedRpcServer struct {
}

func (UnimplementedRpcServer) CreateSecret(context.Context, *Spec) (*Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSecret not implemented")
}
func (UnimplementedRpcServer) DeleteSecret(context.Context, *DeleteSecretRequest) (*Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecret not implemented")
}
func (UnimplementedRpcServer) QuerySecret(context.Context, *QuerySecretRequest) (*SecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySecret not implemented")
}
func (UnimplementedRpcServer) DescSecret(context.Context, *DescSecretRequest) (*Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescSecret not implemented")
}
func (UnimplementedRpcServer) SyncResource(*SyncResourceRequest, Rpc_SyncResourceServer) error {
	return status.Errorf(codes.Unimplemented, "method SyncResource not implemented")
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

func _Rpc_CreateSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Spec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).CreateSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_CreateSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).CreateSecret(ctx, req.(*Spec))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_DeleteSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).DeleteSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_DeleteSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).DeleteSecret(ctx, req.(*DeleteSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_QuerySecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).QuerySecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_QuerySecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).QuerySecret(ctx, req.(*QuerySecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_DescSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).DescSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_DescSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).DescSecret(ctx, req.(*DescSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_SyncResource_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SyncResourceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RpcServer).SyncResource(m, &rpcSyncResourceServer{stream})
}

type Rpc_SyncResourceServer interface {
	Send(*resource.Resource) error
	grpc.ServerStream
}

type rpcSyncResourceServer struct {
	grpc.ServerStream
}

func (x *rpcSyncResourceServer) Send(m *resource.Resource) error {
	return x.ServerStream.SendMsg(m)
}

// Rpc_ServiceDesc is the grpc.ServiceDesc for Rpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cmdb.secret.Rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSecret",
			Handler:    _Rpc_CreateSecret_Handler,
		},
		{
			MethodName: "DeleteSecret",
			Handler:    _Rpc_DeleteSecret_Handler,
		},
		{
			MethodName: "QuerySecret",
			Handler:    _Rpc_QuerySecret_Handler,
		},
		{
			MethodName: "DescSecret",
			Handler:    _Rpc_DescSecret_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SyncResource",
			Handler:       _Rpc_SyncResource_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "apps/secret/pb/rpc.proto",
}
