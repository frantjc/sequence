package step

import (
	"context"

	api "github.com/frantjc/sequence/api/v1/step"
	"github.com/frantjc/sequence/conf"
	"github.com/frantjc/sequence/internal/convert"
	"github.com/frantjc/sequence/internal/grpcio"
	"github.com/frantjc/sequence/runtime"
	"github.com/frantjc/sequence/workflow"
	"google.golang.org/grpc"
)

type stepClient struct {
	runtime runtime.Runtime
}

var _ api.StepClient = &stepClient{}

func (c *stepClient) RunStep(ctx context.Context, in *api.RunStepRequest, _ ...grpc.CallOption) (api.Step_RunStepClient, error) {
	var (
		conf, err = conf.NewFromFlagsWithRepository(in.Repository)
		stream    = grpcio.NewLogStream(ctx)
		opts      = []workflow.ExecOpt{
			workflow.WithRuntime(c.runtime),
			workflow.WithGitHubToken(conf.GitHub.Token),
			workflow.WithRepository(in.Repository),
			workflow.WithStdout(grpcio.NewLogOutStreamWriter(stream)),
			workflow.WithStderr(grpcio.NewLogErrStreamWriter(stream)),
		}
	)
	if err != nil {
		return nil, err
	}

	if in.Verbose || conf.Verbose {
		opts = append(opts, workflow.WithVerbose)
	}

	if in.RunnerImage != "" {
		opts = append(opts, workflow.WithRunnerImage(in.RunnerImage))
	} else {
		in.RunnerImage = conf.Runtime.RunnerImage
	}

	if in.Job != nil {
		opts = append(opts, workflow.WithJob(
			convert.ProtoJobToJob(in.Job),
		))
	}

	executor, err := workflow.NewStepExecutor(convert.ProtoStepToStep(in.Step), opts...)
	if err != nil {
		return nil, err
	}

	go func() {
		defer stream.CloseSend()
		if err = executor.Start(ctx); err != nil {
			stream.SendErr(err)
		}
	}()

	return stream, nil
}
