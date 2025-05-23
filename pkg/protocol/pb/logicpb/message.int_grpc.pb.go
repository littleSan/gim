// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: logic/message.int.proto

package logicpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MessageIntService_Sync_FullMethodName       = "/logic.MessageIntService/Sync"
	MessageIntService_MessageACK_FullMethodName = "/logic.MessageIntService/MessageACK"
	MessageIntService_Pushs_FullMethodName      = "/logic.MessageIntService/Pushs"
	MessageIntService_PushAll_FullMethodName    = "/logic.MessageIntService/PushAll"
)

// MessageIntServiceClient is the client API for MessageIntService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageIntServiceClient interface {
	// 消息同步
	Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncReply, error)
	// 设备收到消息回执
	MessageACK(ctx context.Context, in *MessageACKRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 推送
	Pushs(ctx context.Context, in *PushsRequest, opts ...grpc.CallOption) (*PushsReply, error)
	// 全服推送
	PushAll(ctx context.Context, in *PushAllRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type messageIntServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageIntServiceClient(cc grpc.ClientConnInterface) MessageIntServiceClient {
	return &messageIntServiceClient{cc}
}

func (c *messageIntServiceClient) Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SyncReply)
	err := c.cc.Invoke(ctx, MessageIntService_Sync_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageIntServiceClient) MessageACK(ctx context.Context, in *MessageACKRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MessageIntService_MessageACK_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageIntServiceClient) Pushs(ctx context.Context, in *PushsRequest, opts ...grpc.CallOption) (*PushsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PushsReply)
	err := c.cc.Invoke(ctx, MessageIntService_Pushs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageIntServiceClient) PushAll(ctx context.Context, in *PushAllRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MessageIntService_PushAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageIntServiceServer is the server API for MessageIntService service.
// All implementations must embed UnimplementedMessageIntServiceServer
// for forward compatibility.
type MessageIntServiceServer interface {
	// 消息同步
	Sync(context.Context, *SyncRequest) (*SyncReply, error)
	// 设备收到消息回执
	MessageACK(context.Context, *MessageACKRequest) (*emptypb.Empty, error)
	// 推送
	Pushs(context.Context, *PushsRequest) (*PushsReply, error)
	// 全服推送
	PushAll(context.Context, *PushAllRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMessageIntServiceServer()
}

// UnimplementedMessageIntServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMessageIntServiceServer struct{}

func (UnimplementedMessageIntServiceServer) Sync(context.Context, *SyncRequest) (*SyncReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedMessageIntServiceServer) MessageACK(context.Context, *MessageACKRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageACK not implemented")
}
func (UnimplementedMessageIntServiceServer) Pushs(context.Context, *PushsRequest) (*PushsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pushs not implemented")
}
func (UnimplementedMessageIntServiceServer) PushAll(context.Context, *PushAllRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushAll not implemented")
}
func (UnimplementedMessageIntServiceServer) mustEmbedUnimplementedMessageIntServiceServer() {}
func (UnimplementedMessageIntServiceServer) testEmbeddedByValue()                           {}

// UnsafeMessageIntServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageIntServiceServer will
// result in compilation errors.
type UnsafeMessageIntServiceServer interface {
	mustEmbedUnimplementedMessageIntServiceServer()
}

func RegisterMessageIntServiceServer(s grpc.ServiceRegistrar, srv MessageIntServiceServer) {
	// If the following call pancis, it indicates UnimplementedMessageIntServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MessageIntService_ServiceDesc, srv)
}

func _MessageIntService_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageIntServiceServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageIntService_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageIntServiceServer).Sync(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageIntService_MessageACK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageACKRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageIntServiceServer).MessageACK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageIntService_MessageACK_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageIntServiceServer).MessageACK(ctx, req.(*MessageACKRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageIntService_Pushs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageIntServiceServer).Pushs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageIntService_Pushs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageIntServiceServer).Pushs(ctx, req.(*PushsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageIntService_PushAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageIntServiceServer).PushAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageIntService_PushAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageIntServiceServer).PushAll(ctx, req.(*PushAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageIntService_ServiceDesc is the grpc.ServiceDesc for MessageIntService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageIntService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logic.MessageIntService",
	HandlerType: (*MessageIntServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sync",
			Handler:    _MessageIntService_Sync_Handler,
		},
		{
			MethodName: "MessageACK",
			Handler:    _MessageIntService_MessageACK_Handler,
		},
		{
			MethodName: "Pushs",
			Handler:    _MessageIntService_Pushs_Handler,
		},
		{
			MethodName: "PushAll",
			Handler:    _MessageIntService_PushAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logic/message.int.proto",
}
