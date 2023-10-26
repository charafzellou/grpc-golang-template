// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/main.proto

package grpc_golang_template

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

// RequestServiceClient is the client API for RequestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RequestServiceClient interface {
	RequestMethod(ctx context.Context, in *InputRequest, opts ...grpc.CallOption) (*OutputRequest, error)
}

type requestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRequestServiceClient(cc grpc.ClientConnInterface) RequestServiceClient {
	return &requestServiceClient{cc}
}

func (c *requestServiceClient) RequestMethod(ctx context.Context, in *InputRequest, opts ...grpc.CallOption) (*OutputRequest, error) {
	out := new(OutputRequest)
	err := c.cc.Invoke(ctx, "/proto.RequestService/RequestMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RequestServiceServer is the server API for RequestService service.
// All implementations must embed UnimplementedRequestServiceServer
// for forward compatibility
type RequestServiceServer interface {
	RequestMethod(context.Context, *InputRequest) (*OutputRequest, error)
	mustEmbedUnimplementedRequestServiceServer()
}

// UnimplementedRequestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRequestServiceServer struct {
}

func (UnimplementedRequestServiceServer) RequestMethod(context.Context, *InputRequest) (*OutputRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestMethod not implemented")
}
func (UnimplementedRequestServiceServer) mustEmbedUnimplementedRequestServiceServer() {}

// UnsafeRequestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RequestServiceServer will
// result in compilation errors.
type UnsafeRequestServiceServer interface {
	mustEmbedUnimplementedRequestServiceServer()
}

func RegisterRequestServiceServer(s grpc.ServiceRegistrar, srv RequestServiceServer) {
	s.RegisterService(&RequestService_ServiceDesc, srv)
}

func _RequestService_RequestMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).RequestMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RequestService/RequestMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).RequestMethod(ctx, req.(*InputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RequestService_ServiceDesc is the grpc.ServiceDesc for RequestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RequestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RequestService",
	HandlerType: (*RequestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestMethod",
			Handler:    _RequestService_RequestMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/main.proto",
}
