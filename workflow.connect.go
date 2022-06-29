// Code generated by protoc-gen-sqnc. DO NOT EDIT.
//
// Source: workflow.proto

package sequence

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// WorkflowServiceName is the fully-qualified name of the WorkflowService service.
	WorkflowServiceName = "sequence.WorkflowService"
)

// WorkflowServiceClient is a client for the sequence.WorkflowService service.
type WorkflowServiceClient interface {
	RunWorkflow(context.Context, *connect_go.Request[RunWorkflowRequest]) (*connect_go.ServerStreamForClient[RunWorkflowResponse], error)
	RunJob(context.Context, *connect_go.Request[RunJobRequest]) (*connect_go.ServerStreamForClient[RunJobResponse], error)
	RunStep(context.Context, *connect_go.Request[RunStepRequest]) (*connect_go.ServerStreamForClient[RunStepResponse], error)
}

// NewWorkflowServiceClient constructs a client for the sequence.WorkflowService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewWorkflowServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) WorkflowServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &workflowServiceClient{
		runWorkflow: connect_go.NewClient[RunWorkflowRequest, RunWorkflowResponse](
			httpClient,
			baseURL+"/sequence.WorkflowService/RunWorkflow",
			opts...,
		),
		runJob: connect_go.NewClient[RunJobRequest, RunJobResponse](
			httpClient,
			baseURL+"/sequence.WorkflowService/RunJob",
			opts...,
		),
		runStep: connect_go.NewClient[RunStepRequest, RunStepResponse](
			httpClient,
			baseURL+"/sequence.WorkflowService/RunStep",
			opts...,
		),
	}
}

// workflowServiceClient implements WorkflowServiceClient.
type workflowServiceClient struct {
	runWorkflow *connect_go.Client[RunWorkflowRequest, RunWorkflowResponse]
	runJob      *connect_go.Client[RunJobRequest, RunJobResponse]
	runStep     *connect_go.Client[RunStepRequest, RunStepResponse]
}

// RunWorkflow calls sequence.WorkflowService.RunWorkflow.
func (c *workflowServiceClient) RunWorkflow(ctx context.Context, req *connect_go.Request[RunWorkflowRequest]) (*connect_go.ServerStreamForClient[RunWorkflowResponse], error) {
	return c.runWorkflow.CallServerStream(ctx, req)
}

// RunJob calls sequence.WorkflowService.RunJob.
func (c *workflowServiceClient) RunJob(ctx context.Context, req *connect_go.Request[RunJobRequest]) (*connect_go.ServerStreamForClient[RunJobResponse], error) {
	return c.runJob.CallServerStream(ctx, req)
}

// RunStep calls sequence.WorkflowService.RunStep.
func (c *workflowServiceClient) RunStep(ctx context.Context, req *connect_go.Request[RunStepRequest]) (*connect_go.ServerStreamForClient[RunStepResponse], error) {
	return c.runStep.CallServerStream(ctx, req)
}

// WorkflowServiceHandler is an implementation of the sequence.WorkflowService service.
type WorkflowServiceHandler interface {
	RunWorkflow(context.Context, *connect_go.Request[RunWorkflowRequest], *connect_go.ServerStream[RunWorkflowResponse]) error
	RunJob(context.Context, *connect_go.Request[RunJobRequest], *connect_go.ServerStream[RunJobResponse]) error
	RunStep(context.Context, *connect_go.Request[RunStepRequest], *connect_go.ServerStream[RunStepResponse]) error
}

// NewWorkflowServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewWorkflowServiceHandler(svc WorkflowServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/sequence.WorkflowService/RunWorkflow", connect_go.NewServerStreamHandler(
		"/sequence.WorkflowService/RunWorkflow",
		svc.RunWorkflow,
		opts...,
	))
	mux.Handle("/sequence.WorkflowService/RunJob", connect_go.NewServerStreamHandler(
		"/sequence.WorkflowService/RunJob",
		svc.RunJob,
		opts...,
	))
	mux.Handle("/sequence.WorkflowService/RunStep", connect_go.NewServerStreamHandler(
		"/sequence.WorkflowService/RunStep",
		svc.RunStep,
		opts...,
	))
	return "/sequence.WorkflowService/", mux
}

// UnimplementedWorkflowServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedWorkflowServiceHandler struct{}

func (UnimplementedWorkflowServiceHandler) RunWorkflow(context.Context, *connect_go.Request[RunWorkflowRequest], *connect_go.ServerStream[RunWorkflowResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sequence.WorkflowService.RunWorkflow is not implemented"))
}

func (UnimplementedWorkflowServiceHandler) RunJob(context.Context, *connect_go.Request[RunJobRequest], *connect_go.ServerStream[RunJobResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sequence.WorkflowService.RunJob is not implemented"))
}

func (UnimplementedWorkflowServiceHandler) RunStep(context.Context, *connect_go.Request[RunStepRequest], *connect_go.ServerStream[RunStepResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sequence.WorkflowService.RunStep is not implemented"))
}
