// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: app/Adapter/In/ApiGrcp/proto/call.proto

package call

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

// CallServiceClient is the client API for CallService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CallServiceClient interface {
	CallUnary(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error)
	CallServerStream(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (CallService_CallServerStreamClient, error)
	CallClientStream(ctx context.Context, opts ...grpc.CallOption) (CallService_CallClientStreamClient, error)
	CallBidirectional(ctx context.Context, opts ...grpc.CallOption) (CallService_CallBidirectionalClient, error)
}

type callServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCallServiceClient(cc grpc.ClientConnInterface) CallServiceClient {
	return &callServiceClient{cc}
}

func (c *callServiceClient) CallUnary(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error) {
	out := new(CallResponse)
	err := c.cc.Invoke(ctx, "/call.CallService/CallUnary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callServiceClient) CallServerStream(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (CallService_CallServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CallService_ServiceDesc.Streams[0], "/call.CallService/CallServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &callServiceCallServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CallService_CallServerStreamClient interface {
	Recv() (*CallResponse, error)
	grpc.ClientStream
}

type callServiceCallServerStreamClient struct {
	grpc.ClientStream
}

func (x *callServiceCallServerStreamClient) Recv() (*CallResponse, error) {
	m := new(CallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *callServiceClient) CallClientStream(ctx context.Context, opts ...grpc.CallOption) (CallService_CallClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CallService_ServiceDesc.Streams[1], "/call.CallService/CallClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &callServiceCallClientStreamClient{stream}
	return x, nil
}

type CallService_CallClientStreamClient interface {
	Send(*CallRequest) error
	CloseAndRecv() (*CallResponse, error)
	grpc.ClientStream
}

type callServiceCallClientStreamClient struct {
	grpc.ClientStream
}

func (x *callServiceCallClientStreamClient) Send(m *CallRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *callServiceCallClientStreamClient) CloseAndRecv() (*CallResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *callServiceClient) CallBidirectional(ctx context.Context, opts ...grpc.CallOption) (CallService_CallBidirectionalClient, error) {
	stream, err := c.cc.NewStream(ctx, &CallService_ServiceDesc.Streams[2], "/call.CallService/CallBidirectional", opts...)
	if err != nil {
		return nil, err
	}
	x := &callServiceCallBidirectionalClient{stream}
	return x, nil
}

type CallService_CallBidirectionalClient interface {
	Send(*CallRequest) error
	Recv() (*CallResponse, error)
	grpc.ClientStream
}

type callServiceCallBidirectionalClient struct {
	grpc.ClientStream
}

func (x *callServiceCallBidirectionalClient) Send(m *CallRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *callServiceCallBidirectionalClient) Recv() (*CallResponse, error) {
	m := new(CallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CallServiceServer is the server API for CallService service.
// All implementations must embed UnimplementedCallServiceServer
// for forward compatibility
type CallServiceServer interface {
	CallUnary(context.Context, *CallRequest) (*CallResponse, error)
	CallServerStream(*CallRequest, CallService_CallServerStreamServer) error
	CallClientStream(CallService_CallClientStreamServer) error
	CallBidirectional(CallService_CallBidirectionalServer) error
	mustEmbedUnimplementedCallServiceServer()
}

// UnimplementedCallServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCallServiceServer struct {
}

func (UnimplementedCallServiceServer) CallUnary(context.Context, *CallRequest) (*CallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallUnary not implemented")
}
func (UnimplementedCallServiceServer) CallServerStream(*CallRequest, CallService_CallServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CallServerStream not implemented")
}
func (UnimplementedCallServiceServer) CallClientStream(CallService_CallClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CallClientStream not implemented")
}
func (UnimplementedCallServiceServer) CallBidirectional(CallService_CallBidirectionalServer) error {
	return status.Errorf(codes.Unimplemented, "method CallBidirectional not implemented")
}
func (UnimplementedCallServiceServer) mustEmbedUnimplementedCallServiceServer() {}

// UnsafeCallServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CallServiceServer will
// result in compilation errors.
type UnsafeCallServiceServer interface {
	mustEmbedUnimplementedCallServiceServer()
}

func RegisterCallServiceServer(s grpc.ServiceRegistrar, srv CallServiceServer) {
	s.RegisterService(&CallService_ServiceDesc, srv)
}

func _CallService_CallUnary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceServer).CallUnary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/call.CallService/CallUnary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceServer).CallUnary(ctx, req.(*CallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallService_CallServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CallRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CallServiceServer).CallServerStream(m, &callServiceCallServerStreamServer{stream})
}

type CallService_CallServerStreamServer interface {
	Send(*CallResponse) error
	grpc.ServerStream
}

type callServiceCallServerStreamServer struct {
	grpc.ServerStream
}

func (x *callServiceCallServerStreamServer) Send(m *CallResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CallService_CallClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CallServiceServer).CallClientStream(&callServiceCallClientStreamServer{stream})
}

type CallService_CallClientStreamServer interface {
	SendAndClose(*CallResponse) error
	Recv() (*CallRequest, error)
	grpc.ServerStream
}

type callServiceCallClientStreamServer struct {
	grpc.ServerStream
}

func (x *callServiceCallClientStreamServer) SendAndClose(m *CallResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *callServiceCallClientStreamServer) Recv() (*CallRequest, error) {
	m := new(CallRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CallService_CallBidirectional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CallServiceServer).CallBidirectional(&callServiceCallBidirectionalServer{stream})
}

type CallService_CallBidirectionalServer interface {
	Send(*CallResponse) error
	Recv() (*CallRequest, error)
	grpc.ServerStream
}

type callServiceCallBidirectionalServer struct {
	grpc.ServerStream
}

func (x *callServiceCallBidirectionalServer) Send(m *CallResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *callServiceCallBidirectionalServer) Recv() (*CallRequest, error) {
	m := new(CallRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CallService_ServiceDesc is the grpc.ServiceDesc for CallService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CallService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "call.CallService",
	HandlerType: (*CallServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CallUnary",
			Handler:    _CallService_CallUnary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CallServerStream",
			Handler:       _CallService_CallServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CallClientStream",
			Handler:       _CallService_CallClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CallBidirectional",
			Handler:       _CallService_CallBidirectional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "app/Adapter/In/ApiGrcp/proto/call.proto",
}
