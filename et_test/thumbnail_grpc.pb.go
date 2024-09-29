// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: proto/thumbnail.proto

package et_test

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ThumbnailService_GetThumbnail_FullMethodName = "/proto.ThumbnailService/GetThumbnail"
)

// ThumbnailServiceClient is the client API for ThumbnailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThumbnailServiceClient interface {
	GetThumbnail(ctx context.Context, in *GetThumbnailRequest, opts ...grpc.CallOption) (*GetThumbnailResponse, error)
}

type thumbnailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewThumbnailServiceClient(cc grpc.ClientConnInterface) ThumbnailServiceClient {
	return &thumbnailServiceClient{cc}
}

func (c *thumbnailServiceClient) GetThumbnail(ctx context.Context, in *GetThumbnailRequest, opts ...grpc.CallOption) (*GetThumbnailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetThumbnailResponse)
	err := c.cc.Invoke(ctx, ThumbnailService_GetThumbnail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThumbnailServiceServer is the server API for ThumbnailService service.
// All implementations must embed UnimplementedThumbnailServiceServer
// for forward compatibility.
type ThumbnailServiceServer interface {
	GetThumbnail(context.Context, *GetThumbnailRequest) (*GetThumbnailResponse, error)
	mustEmbedUnimplementedThumbnailServiceServer()
}

// UnimplementedThumbnailServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedThumbnailServiceServer struct{}

func (UnimplementedThumbnailServiceServer) GetThumbnail(context.Context, *GetThumbnailRequest) (*GetThumbnailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThumbnail not implemented")
}
func (UnimplementedThumbnailServiceServer) mustEmbedUnimplementedThumbnailServiceServer() {}
func (UnimplementedThumbnailServiceServer) testEmbeddedByValue()                          {}

// UnsafeThumbnailServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThumbnailServiceServer will
// result in compilation errors.
type UnsafeThumbnailServiceServer interface {
	mustEmbedUnimplementedThumbnailServiceServer()
}

func RegisterThumbnailServiceServer(s grpc.ServiceRegistrar, srv ThumbnailServiceServer) {
	// If the following call pancis, it indicates UnimplementedThumbnailServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ThumbnailService_ServiceDesc, srv)
}

func _ThumbnailService_GetThumbnail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetThumbnailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThumbnailServiceServer).GetThumbnail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThumbnailService_GetThumbnail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThumbnailServiceServer).GetThumbnail(ctx, req.(*GetThumbnailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ThumbnailService_ServiceDesc is the grpc.ServiceDesc for ThumbnailService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ThumbnailService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ThumbnailService",
	HandlerType: (*ThumbnailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetThumbnail",
			Handler:    _ThumbnailService_GetThumbnail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/thumbnail.proto",
}
