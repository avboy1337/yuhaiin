// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package app

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

// ConnectionsClient is the client API for Connections service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectionsClient interface {
	Conns(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ConnResp, error)
	CloseConn(ctx context.Context, in *CloseConnsReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Statistic(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Connections_StatisticClient, error)
}

type connectionsClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectionsClient(cc grpc.ClientConnInterface) ConnectionsClient {
	return &connectionsClient{cc}
}

func (c *connectionsClient) Conns(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ConnResp, error) {
	out := new(ConnResp)
	err := c.cc.Invoke(ctx, "/yuhaiin.app.connmanager.connections/conns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionsClient) CloseConn(ctx context.Context, in *CloseConnsReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/yuhaiin.app.connmanager.connections/close_conn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionsClient) Statistic(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Connections_StatisticClient, error) {
	stream, err := c.cc.NewStream(ctx, &Connections_ServiceDesc.Streams[0], "/yuhaiin.app.connmanager.connections/statistic", opts...)
	if err != nil {
		return nil, err
	}
	x := &connectionsStatisticClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Connections_StatisticClient interface {
	Recv() (*RateResp, error)
	grpc.ClientStream
}

type connectionsStatisticClient struct {
	grpc.ClientStream
}

func (x *connectionsStatisticClient) Recv() (*RateResp, error) {
	m := new(RateResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConnectionsServer is the server API for Connections service.
// All implementations must embed UnimplementedConnectionsServer
// for forward compatibility
type ConnectionsServer interface {
	Conns(context.Context, *emptypb.Empty) (*ConnResp, error)
	CloseConn(context.Context, *CloseConnsReq) (*emptypb.Empty, error)
	Statistic(*emptypb.Empty, Connections_StatisticServer) error
	mustEmbedUnimplementedConnectionsServer()
}

// UnimplementedConnectionsServer must be embedded to have forward compatible implementations.
type UnimplementedConnectionsServer struct {
}

func (UnimplementedConnectionsServer) Conns(context.Context, *emptypb.Empty) (*ConnResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Conns not implemented")
}
func (UnimplementedConnectionsServer) CloseConn(context.Context, *CloseConnsReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseConn not implemented")
}
func (UnimplementedConnectionsServer) Statistic(*emptypb.Empty, Connections_StatisticServer) error {
	return status.Errorf(codes.Unimplemented, "method Statistic not implemented")
}
func (UnimplementedConnectionsServer) mustEmbedUnimplementedConnectionsServer() {}

// UnsafeConnectionsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectionsServer will
// result in compilation errors.
type UnsafeConnectionsServer interface {
	mustEmbedUnimplementedConnectionsServer()
}

func RegisterConnectionsServer(s grpc.ServiceRegistrar, srv ConnectionsServer) {
	s.RegisterService(&Connections_ServiceDesc, srv)
}

func _Connections_Conns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionsServer).Conns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yuhaiin.app.connmanager.connections/conns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionsServer).Conns(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connections_CloseConn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseConnsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionsServer).CloseConn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yuhaiin.app.connmanager.connections/close_conn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionsServer).CloseConn(ctx, req.(*CloseConnsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connections_Statistic_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConnectionsServer).Statistic(m, &connectionsStatisticServer{stream})
}

type Connections_StatisticServer interface {
	Send(*RateResp) error
	grpc.ServerStream
}

type connectionsStatisticServer struct {
	grpc.ServerStream
}

func (x *connectionsStatisticServer) Send(m *RateResp) error {
	return x.ServerStream.SendMsg(m)
}

// Connections_ServiceDesc is the grpc.ServiceDesc for Connections service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Connections_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yuhaiin.app.connmanager.connections",
	HandlerType: (*ConnectionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "conns",
			Handler:    _Connections_Conns_Handler,
		},
		{
			MethodName: "close_conn",
			Handler:    _Connections_CloseConn_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "statistic",
			Handler:       _Connections_Statistic_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/app/connmanager.proto",
}