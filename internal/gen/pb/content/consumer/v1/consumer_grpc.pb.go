// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: content/consumer/v1/consumer.proto

package v1

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

// ConsumerServiceClient is the client API for ConsumerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConsumerServiceClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PingResponse, error)
	// GetTemplate is designed to allow consumers of the platform to serve the
	// user with a template they can start from. This is more important for
	// languages that require selective formatting or a main function. An example
	// of these languages would be C++, and C.
	GetTemplate(ctx context.Context, in *GetTemplateRequest, opts ...grpc.CallOption) (*GetTemplateResponse, error)
	// GetSupportedLanguages will return a list of languages that can be exposed
	// to the user. This response contains a display name for the language that
	// will contain compiler information if important and will also return the
	// code. The code is the value sent to the server when requesting to compile
	// and run.
	GetSupportedLanguages(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetSupportedLanguagesResponse, error)
	// CompileQueueRequest is the core compile request endpoint. Calling into this
	// will trigger the flow to run the user-submitted code.
	CreateCompileRequest(ctx context.Context, in *CompileRequest, opts ...grpc.CallOption) (*CompileQueueResponse, error)
	// GetCompileResultRequest is required to be called after requesting to
	// compile, all details about the running state and the final output
	// of the compiling and execution are from this.
	GetCompileResultRequest(ctx context.Context, in *CompileResultRequest, opts ...grpc.CallOption) (*CompileResultResponse, error)
}

type consumerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConsumerServiceClient(cc grpc.ClientConnInterface) ConsumerServiceClient {
	return &consumerServiceClient{cc}
}

func (c *consumerServiceClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/content.consumer.v1.ConsumerService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) GetTemplate(ctx context.Context, in *GetTemplateRequest, opts ...grpc.CallOption) (*GetTemplateResponse, error) {
	out := new(GetTemplateResponse)
	err := c.cc.Invoke(ctx, "/content.consumer.v1.ConsumerService/GetTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) GetSupportedLanguages(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetSupportedLanguagesResponse, error) {
	out := new(GetSupportedLanguagesResponse)
	err := c.cc.Invoke(ctx, "/content.consumer.v1.ConsumerService/GetSupportedLanguages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) CreateCompileRequest(ctx context.Context, in *CompileRequest, opts ...grpc.CallOption) (*CompileQueueResponse, error) {
	out := new(CompileQueueResponse)
	err := c.cc.Invoke(ctx, "/content.consumer.v1.ConsumerService/CreateCompileRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) GetCompileResultRequest(ctx context.Context, in *CompileResultRequest, opts ...grpc.CallOption) (*CompileResultResponse, error) {
	out := new(CompileResultResponse)
	err := c.cc.Invoke(ctx, "/content.consumer.v1.ConsumerService/GetCompileResultRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsumerServiceServer is the server API for ConsumerService service.
// All implementations should embed UnimplementedConsumerServiceServer
// for forward compatibility
type ConsumerServiceServer interface {
	Ping(context.Context, *emptypb.Empty) (*PingResponse, error)
	// GetTemplate is designed to allow consumers of the platform to serve the
	// user with a template they can start from. This is more important for
	// languages that require selective formatting or a main function. An example
	// of these languages would be C++, and C.
	GetTemplate(context.Context, *GetTemplateRequest) (*GetTemplateResponse, error)
	// GetSupportedLanguages will return a list of languages that can be exposed
	// to the user. This response contains a display name for the language that
	// will contain compiler information if important and will also return the
	// code. The code is the value sent to the server when requesting to compile
	// and run.
	GetSupportedLanguages(context.Context, *emptypb.Empty) (*GetSupportedLanguagesResponse, error)
	// CompileQueueRequest is the core compile request endpoint. Calling into this
	// will trigger the flow to run the user-submitted code.
	CreateCompileRequest(context.Context, *CompileRequest) (*CompileQueueResponse, error)
	// GetCompileResultRequest is required to be called after requesting to
	// compile, all details about the running state and the final output
	// of the compiling and execution are from this.
	GetCompileResultRequest(context.Context, *CompileResultRequest) (*CompileResultResponse, error)
}

// UnimplementedConsumerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedConsumerServiceServer struct {
}

func (UnimplementedConsumerServiceServer) Ping(context.Context, *emptypb.Empty) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedConsumerServiceServer) GetTemplate(context.Context, *GetTemplateRequest) (*GetTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplate not implemented")
}
func (UnimplementedConsumerServiceServer) GetSupportedLanguages(context.Context, *emptypb.Empty) (*GetSupportedLanguagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupportedLanguages not implemented")
}
func (UnimplementedConsumerServiceServer) CreateCompileRequest(context.Context, *CompileRequest) (*CompileQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompileRequest not implemented")
}
func (UnimplementedConsumerServiceServer) GetCompileResultRequest(context.Context, *CompileResultRequest) (*CompileResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompileResultRequest not implemented")
}

// UnsafeConsumerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsumerServiceServer will
// result in compilation errors.
type UnsafeConsumerServiceServer interface {
	mustEmbedUnimplementedConsumerServiceServer()
}

func RegisterConsumerServiceServer(s grpc.ServiceRegistrar, srv ConsumerServiceServer) {
	s.RegisterService(&ConsumerService_ServiceDesc, srv)
}

func _ConsumerService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.consumer.v1.ConsumerService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_GetTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).GetTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.consumer.v1.ConsumerService/GetTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).GetTemplate(ctx, req.(*GetTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_GetSupportedLanguages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).GetSupportedLanguages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.consumer.v1.ConsumerService/GetSupportedLanguages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).GetSupportedLanguages(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_CreateCompileRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).CreateCompileRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.consumer.v1.ConsumerService/CreateCompileRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).CreateCompileRequest(ctx, req.(*CompileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_GetCompileResultRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompileResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).GetCompileResultRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.consumer.v1.ConsumerService/GetCompileResultRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).GetCompileResultRequest(ctx, req.(*CompileResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConsumerService_ServiceDesc is the grpc.ServiceDesc for ConsumerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConsumerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content.consumer.v1.ConsumerService",
	HandlerType: (*ConsumerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ConsumerService_Ping_Handler,
		},
		{
			MethodName: "GetTemplate",
			Handler:    _ConsumerService_GetTemplate_Handler,
		},
		{
			MethodName: "GetSupportedLanguages",
			Handler:    _ConsumerService_GetSupportedLanguages_Handler,
		},
		{
			MethodName: "CreateCompileRequest",
			Handler:    _ConsumerService_CreateCompileRequest_Handler,
		},
		{
			MethodName: "GetCompileResultRequest",
			Handler:    _ConsumerService_GetCompileResultRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content/consumer/v1/consumer.proto",
}
