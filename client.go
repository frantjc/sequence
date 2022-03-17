package sequence

import (
	"context"
	"io"

	containerapi "github.com/frantjc/sequence/api/v1/container"
	imageapi "github.com/frantjc/sequence/api/v1/image"
	jobapi "github.com/frantjc/sequence/api/v1/job"
	stepapi "github.com/frantjc/sequence/api/v1/step"
	workflowapi "github.com/frantjc/sequence/api/v1/workflow"
	"github.com/frantjc/sequence/internal/convert"
	"github.com/frantjc/sequence/internal/grpcio"
	"github.com/frantjc/sequence/runtime"
	"github.com/frantjc/sequence/workflow"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client is a wrapper around each of sequence's gRPC clients
type Client struct {
	jobClient      jobapi.JobClient
	stepClient     stepapi.StepClient
	workflowClient workflowapi.WorkflowClient

	containerClient containerapi.ContainerClient
	imageClient     imageapi.ImageClient
}

// New is an alias to NewClient
var New = NewClient

// NewClient returns a new Client
func NewClient(ctx context.Context, addr string, opts ...ClientOpt) (*Client, error) {
	cc, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := &Client{
		jobClient:      jobapi.NewJobClient(cc),
		stepClient:     stepapi.NewStepClient(cc),
		workflowClient: workflowapi.NewWorkflowClient(cc),
	}
	return client, nil
}

// JobClient returns the client's underlying gRPC JobClient
func (c *Client) JobClient() jobapi.JobClient {
	return c.jobClient
}

// StepClient returns the client's underlying gRPC StepClient
func (c *Client) StepClient() stepapi.StepClient {
	return c.stepClient
}

// WorkflowClient returns the client's underlying gRPC WorkflowClient
func (c *Client) WorkflowClient() workflowapi.WorkflowClient {
	return c.workflowClient
}

// ContainerClient returns the client's underlying gRPC WorkflowClient
func (c *Client) ContainerClient() containerapi.ContainerClient {
	return c.containerClient
}

// ImageClient returns the client's underlying gRPC WorkflowClient
func (c *Client) ImageClient() imageapi.ImageClient {
	return c.imageClient
}

// Runtime returns a runtime.Runtime implementation using the underlying clients
func (c *Client) Runtime() runtime.Runtime {
	return NewGRPCRuntime(c.ImageClient(), c.ContainerClient())
}

// RunStep calls the underlying gRPC StepClient's RunStep and
// writes its logs to the given io.Writer
func (c *Client) RunStep(ctx context.Context, step *workflow.Step, w io.Writer) error {
	stream, err := c.StepClient().RunStep(ctx, &stepapi.RunStepRequest{
		Step: convert.StepToProtoStep(step),
	})
	if err != nil {
		return err
	}

	return grpcio.DemultiplexLogStream(stream, w, w)
}

// RunJob calls the underlying gRPC JobClient's RunJob and
// writes its logs to the given io.Writer
func (c *Client) RunJob(ctx context.Context, job *workflow.Job, w io.Writer) error {
	stream, err := c.JobClient().RunJob(ctx, &jobapi.RunJobRequest{
		Job: convert.JobToProtoJob(job),
	})
	if err != nil {
		return err
	}

	return grpcio.DemultiplexLogStream(stream, w, w)
}

// RunWorkflow calls the underlying gRPC WorkflowClient's RunWorkflow and
// writes its logs to the given io.Writer
func (c *Client) RunWorkflow(ctx context.Context, workflow *workflow.Workflow, w io.Writer) error {
	stream, err := c.WorkflowClient().RunWorkflow(ctx, &workflowapi.RunWorkflowRequest{
		Workflow: convert.WorkflowToTypeWorkflow(workflow),
	})
	if err != nil {
		return err
	}

	return grpcio.DemultiplexLogStream(stream, w, w)
}
