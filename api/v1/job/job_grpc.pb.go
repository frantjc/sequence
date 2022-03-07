// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package job

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

// JobClient is the client API for Job service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobClient interface {
	RunJob(ctx context.Context, in *RunJobRequest, opts ...grpc.CallOption) (Job_RunJobClient, error)
}

type jobClient struct {
	cc grpc.ClientConnInterface
}

func NewJobClient(cc grpc.ClientConnInterface) JobClient {
	return &jobClient{cc}
}

func (c *jobClient) RunJob(ctx context.Context, in *RunJobRequest, opts ...grpc.CallOption) (Job_RunJobClient, error) {
	stream, err := c.cc.NewStream(ctx, &Job_ServiceDesc.Streams[0], "/sequence.v1.job.Job/RunJob", opts...)
	if err != nil {
		return nil, err
	}
	x := &jobRunJobClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Job_RunJobClient interface {
	Recv() (*types.Log, error)
	grpc.ClientStream
}

type jobRunJobClient struct {
	grpc.ClientStream
}

func (x *jobRunJobClient) Recv() (*types.Log, error) {
	m := new(types.Log)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// JobServer is the server API for Job service.
// All implementations must embed UnimplementedJobServer
// for forward compatibility
type JobServer interface {
	RunJob(*RunJobRequest, Job_RunJobServer) error
	mustEmbedUnimplementedJobServer()
}

// UnimplementedJobServer must be embedded to have forward compatible implementations.
type UnimplementedJobServer struct {
}

func (UnimplementedJobServer) RunJob(*RunJobRequest, Job_RunJobServer) error {
	return status.Errorf(codes.Unimplemented, "method RunJob not implemented")
}
func (UnimplementedJobServer) mustEmbedUnimplementedJobServer() {}

// UnsafeJobServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobServer will
// result in compilation errors.
type UnsafeJobServer interface {
	mustEmbedUnimplementedJobServer()
}

func RegisterJobServer(s grpc.ServiceRegistrar, srv JobServer) {
	s.RegisterService(&Job_ServiceDesc, srv)
}

func _Job_RunJob_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RunJobRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JobServer).RunJob(m, &jobRunJobServer{stream})
}

type Job_RunJobServer interface {
	Send(*types.Log) error
	grpc.ServerStream
}

type jobRunJobServer struct {
	grpc.ServerStream
}

func (x *jobRunJobServer) Send(m *types.Log) error {
	return x.ServerStream.SendMsg(m)
}

// Job_ServiceDesc is the grpc.ServiceDesc for Job service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Job_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sequence.v1.job.Job",
	HandlerType: (*JobServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RunJob",
			Handler:       _Job_RunJob_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/v1/job/job.proto",
}
