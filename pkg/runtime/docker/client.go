package docker

import (
	"github.com/docker/docker/client"
)

type dockerRuntime struct {
	client *client.Client
}