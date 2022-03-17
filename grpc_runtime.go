package sequence

import (
	"context"

	containerapi "github.com/frantjc/sequence/api/v1/container"
	imageapi "github.com/frantjc/sequence/api/v1/image"
	"github.com/frantjc/sequence/internal/convert"
	"github.com/frantjc/sequence/internal/grpcio"
	"github.com/frantjc/sequence/runtime"
	"github.com/frantjc/sequence/sio"
)

func NewGRPCRuntime(i imageapi.ImageClient, c containerapi.ContainerClient) runtime.Runtime {
	return &runtimeClient{
		containerClient: c,
		imageClient:     i,
	}
}

var (
	_ runtime.Container = &runtimeContainer{}
	_ runtime.Image     = &runtimeImage{}
	_ runtime.Runtime   = &runtimeClient{}
)

type runtimeContainer struct {
	id     string
	client containerapi.ContainerClient
}

func (c *runtimeContainer) ID() string {
	return c.id
}

func (c *runtimeContainer) Exec(ctx context.Context, opts ...runtime.ExecOpt) error {
	var (
		stdout = sio.NewNoOpWriter()
		stderr = stdout
	)
	if exec, err := runtime.NewExec(opts...); err != nil {
		return err
	} else {
		if exec.Stdout != nil {
			stdout = exec.Stdout
		}
		if exec.Stderr != nil {
			stderr = exec.Stderr
		}
	}

	stream, err := c.client.ExecContainer(ctx, &containerapi.ExecContainerRequest{
		Id: c.ID(),
	})
	if err != nil {
		return err
	}

	return grpcio.DemultiplexLogStream(stream, stdout, stderr)
}

type runtimeImage struct {
	ref string
}

func (i *runtimeImage) Ref() string {
	return i.ref
}

type runtimeClient struct {
	containerClient containerapi.ContainerClient
	imageClient     imageapi.ImageClient
}

func (r *runtimeClient) ContainerClient() containerapi.ContainerClient {
	return r.containerClient
}

func (r *runtimeClient) ImageClient() imageapi.ImageClient {
	return r.imageClient
}

func (r *runtimeClient) PullImage(ctx context.Context, ref string) (runtime.Image, error) {
	res, err := r.imageClient.PullImage(ctx, &imageapi.PullImageRequest{
		Ref: ref,
	})
	if err != nil {
		return nil, err
	}

	return &runtimeImage{
		ref: res.Image.Ref,
	}, nil
}

func (r *runtimeClient) CreateContainer(ctx context.Context, s *runtime.Spec) (runtime.Container, error) {
	res, err := r.containerClient.CreateContainer(ctx, &containerapi.CreateContainerRequest{
		Spec: convert.RuntimeSpecToProtoSpec(s),
	})
	if err != nil {
		return nil, err
	}

	return &runtimeContainer{
		id:     res.Container.Id,
		client: r.ContainerClient(),
	}, nil
}

func (r *runtimeClient) GetContainer(ctx context.Context, id string) (runtime.Container, error) {
	res, err := r.containerClient.GetContainer(ctx, &containerapi.GetContainerRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return &runtimeContainer{
		id:     res.Container.Id,
		client: r.ContainerClient(),
	}, nil
}
