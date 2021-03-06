package sequence_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/frantjc/go-js"
	"github.com/frantjc/sequence"
	"github.com/frantjc/sequence/internal/paths/volumes"
	"github.com/frantjc/sequence/pkg/github/actions"
	"github.com/frantjc/sequence/pkg/github/actions/uses"
	"github.com/frantjc/sequence/runtime"
	"github.com/stretchr/testify/assert"
)

func TestErrOnUnmeetableJobNeeds(t *testing.T) {
	_, err := sequence.NewWorkflowExecutor(ctx, &sequence.Workflow{
		Jobs: map[string]*sequence.Job{
			"needs": {
				Needs: "unmeetable",
			},
			"hello-there": {},
			"general-kenobi": {
				Needs: "hello-there",
			},
		},
	})
	assert.Error(t, err)
	assert.True(t, sequence.IsErrUnmeetableJobNeeds(err))
}

func TestNoErrOnMeetableJobNeeds(t *testing.T) {
	needs := "hello-there"
	_, err := sequence.NewWorkflowExecutor(ctx, &sequence.Workflow{
		Jobs: map[string]*sequence.Job{
			needs: {},
			"general-kenobi": {
				Needs: needs,
			},
		},
	})
	assert.Nil(t, err)
}

func TestWorkflowExecutorNeedsTest(t *testing.T) {
	var (
		value  = "hello there"
		output = "greeting"
		needs  = "hello-there"
		stepID = "test"
	)
	for _, rt := range NewTestRuntimes(t) {
		WorkflowExecutorTest(
			t, rt,
			&sequence.Workflow{
				Jobs: map[string]*sequence.Job{
					"general-kenobi": {
						Needs: needs,
						Steps: []*sequence.Step{
							{
								Image: alpineRef,
								Run:   fmt.Sprintf("echo ::notice::${{ needs.%s.outputs.%s }}", needs, output),
							},
						},
					},
					needs: {
						Steps: []*sequence.Step{
							{
								Id:    stepID,
								Image: alpineRef,
								Run:   fmt.Sprintf("echo ::set-output name=%s::%s", output, value),
							},
						},
						Outputs: map[string]string{
							output: fmt.Sprintf("${{ steps.%s.outputs.%s }}", stepID, output),
						},
					},
				},
			},
			sequence.OnWorkflowCommand(func(event *sequence.Event[*actions.WorkflowCommand]) {
				switch event.Type.Command {
				case actions.CommandSetOutput:
					assert.Equal(t, value, event.Type.Value)
					assert.Equal(t, output, event.Type.GetName())
				case actions.CommandNotice:
					assert.Equal(t, value, event.Type.Value)
				default:
					assert.True(t, false, "unexpected workflow command", event.Type.String())
				}
			}),
		)
	}
}

func WorkflowExecutorTest(t *testing.T, rt runtime.Runtime, workflow *sequence.Workflow, opts ...sequence.ExecutorOpt) {
	t.Helper()

	var (
		imagesPulled           = []runtime.Image{}
		containersCreated      = []runtime.Container{}
		volumesCreated         = []runtime.Volume{}
		workflowCommandsIssued = []*actions.WorkflowCommand{}
	)

	stdout, err := os.CreateTemp("", "")
	assert.NotNil(t, stdout)
	assert.Nil(t, err)
	defer os.Remove(stdout.Name())

	stderr := stdout
	assert.NotNil(t, stderr)

	we, err := NewTestWorkflowExecutor(
		t, rt, workflow,
		append(
			opts,
			sequence.OnImagePull(func(event *sequence.Event[runtime.Image]) {
				imagesPulled = append(imagesPulled, event.Type)
			}),
			sequence.OnContainerCreate(func(event *sequence.Event[runtime.Container]) {
				containersCreated = append(containersCreated, event.Type)
			}),
			sequence.OnVolumeCreate(func(event *sequence.Event[runtime.Volume]) {
				volumesCreated = append(volumesCreated, event.Type)
			}),
			sequence.OnWorkflowCommand(func(event *sequence.Event[*actions.WorkflowCommand]) {
				workflowCommandsIssued = append(workflowCommandsIssued, event.Type)
			}),
			sequence.WithStreams(os.Stdin, stdout, stderr),
		)...,
	)
	assert.NotNil(t, we)
	assert.Nil(t, err)

	assert.Nil(t, we.ExecuteContext(ctx))
	assert.Greater(t, len(imagesPulled), 0)
	assert.Greater(t, len(containersCreated), 0)
	assert.Greater(t, len(volumesCreated), 0)

	for _, job := range workflow.Jobs {
		for _, step := range job.Steps {
			if step.IsGitHubAction() {
				action, err := uses.Parse(step.Uses)
				assert.Nil(t, err)
				assert.True(t, js.Some(volumesCreated, func(v runtime.Volume, _ int, _ []runtime.Volume) bool {
					return volumes.GetActionSource(action) == v.GetSource()
				}))
			}
		}
	}
}

func NewTestWorkflowExecutor(t *testing.T, rt runtime.Runtime, workflow *sequence.Workflow, opts ...sequence.ExecutorOpt) (sequence.Executor, error) {
	t.Helper()

	gc := NewTestGlobalContext(t)

	we, err := sequence.NewWorkflowExecutor(
		ctx, workflow,
		append([]sequence.ExecutorOpt{
			sequence.WithRuntime(rt),
			sequence.WithGlobalContext(gc),
		}, opts...)...,
	)
	assert.NotNil(t, we)
	assert.Nil(t, err)

	return we, err
}
