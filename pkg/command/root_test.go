package command_test

import (
	"testing"

	"github.com/frantjc/sequence/pkg/command"
	"github.com/stretchr/testify/assert"
)

func TestNewRootCommand(t *testing.T) {
	cmd, err := command.NewRootCmd()
	assert.Nil(t, err)
	assert.NotNil(t, cmd)
}
