package sequence

import (
	"context"
	"os"

	"github.com/frantjc/go-js"
	"github.com/frantjc/sequence/internal/paths"
	"github.com/frantjc/sequence/internal/paths/volumes"
	"github.com/frantjc/sequence/pkg/github/actions"
	"github.com/frantjc/sequence/pkg/github/actions/uses"
	"github.com/frantjc/sequence/runtime"
)

type stepsExecutor struct {
	*executor
	steps            []*Step
	preStepWrappers  []*stepWrapper
	mainStepWrappers []*stepWrapper
	postStepWrappers []*stepWrapper
}

func NewStepsExecutor(ctx context.Context, steps []*Step, opts ...ExecutorOpt) (Executor, error) {
	var (
		gc, err = actions.NewContext(paths.GlobalContextOpts()...)
		se      = &stepsExecutor{
			executor: &executor{
				StreamIn:      os.Stdin,
				StreamOut:     os.Stdout,
				StreamErr:     os.Stderr,
				RunnerImage:   DefaultRunnerImage,
				GlobalContext: gc,
			},
			steps: steps,
		}
	)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		if err := opt(se.executor); err != nil {
			return nil, err
		}
	}

	return se, nil
}

func (e *stepsExecutor) Execute() error {
	return e.ExecuteContext(context.Background())
}

func (e *stepsExecutor) ExecuteContext(ctx context.Context) error {
	for _, step := range e.steps {
		if step.IsGitHubAction() {
			action, err := uses.Parse(step.Uses)
			if err != nil {
				return err
			}

			actionMetadata, err := e.SetupAction(ctx, action)
			if err != nil {
				return err
			}

			if actionMetadata.IsComposite() {
				steps, err := NewStepsFromCompositeActionMetadata(actionMetadata, paths.Action)
				if err != nil {
					return err
				}

				rse := &stepsExecutor{
					executor: e.executor,
					steps:    steps,
				}

				if err := rse.ExecuteContext(ctx); err != nil {
					return err
				}

				e.preStepWrappers = append(e.preStepWrappers, rse.preStepWrappers...)
				e.mainStepWrappers = append(e.mainStepWrappers, rse.mainStepWrappers...)
				e.postStepWrappers = append(js.Reverse(rse.postStepWrappers), e.postStepWrappers...)
			} else {
				preStep, mainStep, postStep, err := NewStepsFromNonCompositeMetadata(actionMetadata, paths.Action, step)
				if err != nil {
					return err
				}

				var (
					extraMounts = []*runtime.Mount{
						{
							Source:      volumes.GetActionSource(action),
							Destination: e.GlobalContext.GitHubContext.ActionPath,
							Type:        runtime.MountTypeVolume,
						},
					}
					state = map[string]string{}
				)

				if preStep != nil {
					e.preStepWrappers = append(e.preStepWrappers, &stepWrapper{
						step:        preStep,
						extraMounts: extraMounts,
						state:       state,
						id:          e.GetID(),
					})
				}

				if mainStep != nil {
					e.mainStepWrappers = append(e.mainStepWrappers, &stepWrapper{
						step:        mainStep,
						extraMounts: extraMounts,
						state:       state,
						id:          e.GetID(),
					})
				}

				if postStep != nil {
					e.postStepWrappers = append([]*stepWrapper{
						{
							step:        postStep,
							extraMounts: extraMounts,
							state:       state,
							id:          e.GetID(),
						},
					}, e.postStepWrappers...)
				}
			}
		} else {
			e.mainStepWrappers = append(e.mainStepWrappers, &stepWrapper{
				step: step,
				id:   e.GetID(),
			})
		}
	}

	for _, stepWrapper := range append(
		append(e.preStepWrappers, e.mainStepWrappers...),
		e.postStepWrappers...,
	) {
		e.OnStepStart.Invoke(&Event[*Step]{
			Type:          stepWrapper.step,
			GlobalContext: e.GlobalContext,
		})

		swe := &stepWrapperExecutor{
			executor:           e.executor,
			stepWrapper:        stepWrapper,
			stopCommandsTokens: map[string]bool{},
		}

		if err := swe.Execute(ctx); err != nil {
			return err
		}

		e.executor = swe.executor

		e.OnStepFinish.Invoke(&Event[*Step]{
			Type:          stepWrapper.step,
			GlobalContext: e.GlobalContext,
		})
	}

	return nil
}
