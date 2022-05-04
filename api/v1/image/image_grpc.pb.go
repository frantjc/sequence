// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package image

import (
	context "context"
	types "github.com/frantjc/sequence/api/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ImageClient is the client API for Image service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageClient interface {
	PullImage(ctx context.Context, in *PullImageRequest, opts ...grpc.CallOption) (*PullImageResponse, error)
	PruneImages(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error)
}

type imageClient struct {
	cc grpc.ClientConnInterface
}

func NewImageClient(cc grpc.ClientConnInterface) ImageClient {
	return &imageClient{cc}
}

func (c *imageClient) PullImage(ctx context.Context, in *PullImageRequest, opts ...grpc.CallOption) (*PullImageResponse, error) {
	out := new(PullImageResponse)
	err := c.cc.Invoke(ctx, "/sequence.v1.image.Image/PullImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageClient) PruneImages(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/sequence.v1.image.Image/PruneImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageServer is the server API for Image service.
// All implementations must embed UnimplementedImageServer
// for forward compatibility
type ImageServer interface {
	PullImage(context.Context, *PullImageRequest) (*PullImageResponse, error)
	PruneImages(context.Context, *types.Empty) (*types.Empty, error)
	mustEmbedUnimplementedImageServer()
}

// UnimplementedImageServer must be embedded to have forward compatible implementations.
type UnimplementedImageServer struct {
}

func (UnimplementedImageServer) PullImage(context.Context, *PullImageRequest) (*PullImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullImage not implemented")
}
func (UnimplementedImageServer) PruneImages(context.Context, *types.Empty) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PruneImages not implemented")
}
func (UnimplementedImageServer) mustEmbedUnimplementedImageServer() {}

// UnsafeImageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageServer will
// result in compilation errors.
type UnsafeImageServer interface {
	mustEmbedUnimplementedImageServer()
}

func RegisterImageServer(s grpc.ServiceRegistrar, srv ImageServer) {
	s.RegisterService(&Image_ServiceDesc, srv)
}

func _Image_PullImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).PullImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sequence.v1.image.Image/PullImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).PullImage(ctx, req.(*PullImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Image_PruneImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).PruneImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sequence.v1.image.Image/PruneImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).PruneImages(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Image_ServiceDesc is the grpc.ServiceDesc for Image service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Image_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sequence.v1.image.Image",
	HandlerType: (*ImageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PullImage",
			Handler:    _Image_PullImage_Handler,
		},
		{
			MethodName: "PruneImages",
			Handler:    _Image_PruneImages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/image/image.proto",
}
