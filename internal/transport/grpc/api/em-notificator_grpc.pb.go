// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: internal/transport/grpc/proto/em-notificator.proto

package api

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

// MailSenderServiceClient is the client API for MailSenderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MailSenderServiceClient interface {
	SendMail(ctx context.Context, in *Mail, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type mailSenderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMailSenderServiceClient(cc grpc.ClientConnInterface) MailSenderServiceClient {
	return &mailSenderServiceClient{cc}
}

func (c *mailSenderServiceClient) SendMail(ctx context.Context, in *Mail, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/idler_email.MailSenderService/SendMail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailSenderServiceServer is the server API for MailSenderService service.
// All implementations must embed UnimplementedMailSenderServiceServer
// for forward compatibility
type MailSenderServiceServer interface {
	SendMail(context.Context, *Mail) (*emptypb.Empty, error)
	mustEmbedUnimplementedMailSenderServiceServer()
}

// UnimplementedMailSenderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMailSenderServiceServer struct {
}

func (UnimplementedMailSenderServiceServer) SendMail(context.Context, *Mail) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMail not implemented")
}
func (UnimplementedMailSenderServiceServer) mustEmbedUnimplementedMailSenderServiceServer() {}

// UnsafeMailSenderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MailSenderServiceServer will
// result in compilation errors.
type UnsafeMailSenderServiceServer interface {
	mustEmbedUnimplementedMailSenderServiceServer()
}

func RegisterMailSenderServiceServer(s grpc.ServiceRegistrar, srv MailSenderServiceServer) {
	s.RegisterService(&MailSenderService_ServiceDesc, srv)
}

func _MailSenderService_SendMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailSenderServiceServer).SendMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idler_email.MailSenderService/SendMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailSenderServiceServer).SendMail(ctx, req.(*Mail))
	}
	return interceptor(ctx, in, info, handler)
}

// MailSenderService_ServiceDesc is the grpc.ServiceDesc for MailSenderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MailSenderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "idler_email.MailSenderService",
	HandlerType: (*MailSenderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMail",
			Handler:    _MailSenderService_SendMail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/transport/grpc/proto/em-notificator.proto",
}
