package workflow

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/frantjc/sequence/github/actions"
	"github.com/frantjc/sequence/log"
	"github.com/frantjc/sequence/runtime"
	"github.com/google/uuid"
	"github.com/opencontainers/runtime-spec/specs-go"
)

type githubActionStep struct {
	ID   string
	Name string
	Env  map[string]string
	Uses string
	With map[string]string
	If   interface{}

	Privileged bool
}

var _ executable = &githubActionStep{}

func (e *githubActionStep) id() string {
	if e.ID != "" {
		return e.ID
	} else if e.Name != "" {
		return e.Name
	}

	return e.Uses
}

func (e *githubActionStep) execute(ctx context.Context, ex *jobExecutor) error {
	logout := log.New(ex.stdout).SetVerbose(ex.verbose)

	logout.Debugf("[%sSQNC:DBG%s] parsing 'uses: %s'", log.ColorDebug, log.ColorNone, e.Uses)
	action, err := actions.ParseReference(e.Uses)
	if err != nil {
		return err
	}

	logout.Infof("[%sSQNC%s] setting up action '%s'", log.ColorInfo, log.ColorNone, action.String())
	spec := &runtime.Spec{
		Image:      ex.runnerImage,
		Entrypoint: []string{containerShim, action.String(), ex.globalContext.GitHubContext.ActionPath},
		Cmd:        []string{},
		Mounts: []specs.Mount{
			{
				// actions are global because each step that uses
				// actions/checkout@v2 expects it to function the same
				Source:      ex.actionPath(action),
				Destination: ex.globalContext.GitHubContext.ActionPath,
				Type:        runtime.MountTypeVolume,
			},
		},
	}

	logout.Infof("[%sSQNC%s] pulling image '%s'", log.ColorInfo, log.ColorNone, spec.Image)
	image, err := ex.runtime.PullImage(ctx, spec.Image)
	if err != nil {
		return err
	}
	logout.Debugf("[%sSQNC:DBG%s] finished pulling image '%s'", log.ColorDebug, log.ColorNone, image.Ref())

	logout.Debugf("[%sSQNC:DBG%s] getting or creating volumes", log.ColorDebug, log.ColorNone)
	for _, mount := range spec.Mounts {
		if mount.Type == runtime.MountTypeVolume {
			vol, err := ex.runtime.CreateVolume(ctx, mount.Source)
			if err != nil {
				vol, _ = ex.runtime.GetVolume(ctx, mount.Source)
			}
			mount.Source = vol.Source()
		}
	}
	logout.Debugf("[%sSQNC:DBG%s] finished setting up volumes", log.ColorDebug, log.ColorNone)

	logout.Debugf("[%sSQNC:DBG%s] creating container", log.ColorDebug, log.ColorNone)
	container, err := ex.runtime.CreateContainer(ctx, spec)
	if err != nil {
		return err
	}

	logout.Debugf("[%sSQNC:DBG%s] copying shim to container", log.ColorDebug, log.ColorNone)
	sqncshim, err := shimUsesTarArchive()
	if err != nil {
		return err
	}

	if err = container.CopyTo(ctx, sqncshim, containerShimDir); err != nil {
		return err
	}

	outbuf := new(bytes.Buffer)
	if err = container.Exec(ctx, runtime.ExecStreams(os.Stdin, outbuf, ex.stderr)); err != nil {
		return err
	}

	resp := &StepOut{}
	if err = json.NewDecoder(outbuf).Decode(resp); err != nil {
		return err
	}

	if actionMetadataJSON := []byte(resp.GetActionMetadata()); len(actionMetadataJSON) != 0 {
		logout.Debugf("[%sSQNC:DBG%s] parsing metadata for action '%s'", log.ColorDebug, log.ColorNone, action.String())
		actionMetadata := &actions.Metadata{}
		if err = json.Unmarshal(actionMetadataJSON, actionMetadata); err != nil {
			return err
		}

		if actionMetadata.IsComposite() {
			steps, err := NewStepsFromMetadata(actionMetadata, ex.globalContext.GitHubContext.ActionPath)
			if err != nil {
				return err
			}

			for _, step := range steps {
				if step.IsGitHubAction() {
					githubAction := &githubActionStep{
						ID:         step.ID,
						Name:       step.Name,
						Env:        step.Env,
						Uses:       step.Uses,
						With:       step.With,
						If:         step.If,
						Privileged: step.Privileged,
					}

					if err = githubAction.execute(ctx, ex); err != nil {
						return err
					}
				}
			}

			return nil
		}
		// pre, main and post steps must remain connected via their state
		// but should not share that state with other steps
		// see https://docs.github.com/en/actions/using-workflows/workflow-commands-for-github-actions#sending-values-to-the-pre-and-post-actions
		stateKey := uuid.NewString()
		ex.states[stateKey] = map[string]string{}
		specOpts := []runtime.SpecOpt{
			runtime.WithMounts(spec.Mounts...),
		}
		if preStep, err := NewPreStepFromMetadata(actionMetadata, ex.globalContext.GitHubContext.ActionPath); err != nil {
			return err
		} else if preStep != nil {
			regularStep := &regularStep{
				Env:   preStep.Env,
				Shell: preStep.Shell,
				Run:   preStep.Run,
				If:    preStep.If,
				With:  preStep.With,

				Image:      preStep.Image,
				Entrypoint: preStep.Entrypoint,
				Cmd:        preStep.Cmd,
				Privileged: preStep.Privileged,

				stateKey: stateKey,
				specOpts: specOpts,
			}

			for k, v := range e.With {
				regularStep.With[k] = v
			}

			for k, v := range e.Env {
				regularStep.Env[k] = v
			}

			switch {
			case e.ID != "":
				regularStep.ID = fmt.Sprintf("Pre %s", e.ID)
			case e.Name != "":
				regularStep.Name = fmt.Sprintf("Pre %s", e.Name)
			default:
				regularStep.Name = fmt.Sprintf("Pre %s", e.Uses)
			}

			ex.pre = append(ex.pre, regularStep)
		}

		if mainStep, err := NewMainStepFromMetadata(actionMetadata, ex.globalContext.GitHubContext.ActionPath); err != nil {
			return err
		} else if mainStep != nil {
			regularStep := &regularStep{
				ID:    e.ID,
				Name:  e.Name,
				Env:   mainStep.Env,
				Shell: mainStep.Shell,
				Run:   mainStep.Run,
				If:    mainStep.If,
				With:  mainStep.With,

				Image:      mainStep.Image,
				Entrypoint: mainStep.Entrypoint,
				Cmd:        mainStep.Cmd,
				Privileged: mainStep.Privileged,

				stateKey: stateKey,
				specOpts: specOpts,
			}

			if e.Name == "" {
				regularStep.Name = e.Uses
			}

			for k, v := range e.Env {
				regularStep.Env[k] = v
			}

			for k, v := range e.With {
				regularStep.With[k] = v
			}

			ex.main = append(ex.main, regularStep)
		} else {
			// every non-composite action must have a main step
			return actions.ErrNotAnAction
		}

		if postStep, err := NewPostStepFromMetadata(actionMetadata, ex.globalContext.GitHubContext.ActionPath); err != nil {
			return err
		} else if postStep != nil {
			regularStep := &regularStep{
				Env:   postStep.Env,
				Shell: postStep.Shell,
				Run:   postStep.Run,
				If:    postStep.If,
				With:  postStep.With,

				Image:      postStep.Image,
				Entrypoint: postStep.Entrypoint,
				Cmd:        postStep.Cmd,
				Privileged: postStep.Privileged,

				stateKey: stateKey,
				specOpts: specOpts,
			}

			for k, v := range e.With {
				regularStep.With[k] = v
			}

			for k, v := range e.Env {
				regularStep.Env[k] = v
			}

			if e.ID != "" {
				regularStep.ID = fmt.Sprintf("Post %s", e.ID)
			} else if e.Name != "" {
				regularStep.Name = fmt.Sprintf("Post %s", e.Name)
			} else {
				regularStep.Name = fmt.Sprintf("Post %s", e.Uses)
			}

			ex.post = append([]executable{regularStep}, ex.post...)
		}
		return nil
	}

	logout.Infof("[%sSQNC:ERR%s] not an action '%s'", log.ColorError, log.ColorNone, action.String())
	return actions.ErrNotAnAction
}
