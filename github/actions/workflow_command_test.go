package actions_test

import (
	"testing"

	"github.com/frantjc/sequence/github/actions"
	"github.com/stretchr/testify/assert"
)

func TestWorkflowCommandToString(t *testing.T) {
	var (
		command = &actions.WorkflowCommand{
			Command: "set-output",
			Parameters: map[string]string{
				"name":       "var",
				"otherParam": "param",
			},
			Value: "value",
		}
		expected = "::set-output name=var,otherParam=param::value"
		actual   = command.String()
	)
	assert.Equal(t, expected, actual)
}