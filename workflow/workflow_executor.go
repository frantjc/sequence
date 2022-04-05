package workflow

import "context"

func NewWorkflowExecutor(w *Workflow, opts ...ExecOpt) (Executor, error) {
	return &workflowExecutor{w, append(opts, WithWorkflow(w))}, nil
}

type workflowExecutor struct {
	workflow *Workflow
	opts     []ExecOpt
}

var _ Executor = &workflowExecutor{}

func (e *workflowExecutor) Start(ctx context.Context) error {
	// TODO ordering, job outputs, needs, etc
	for _, job := range e.workflow.Jobs {
		ex, err := NewJobExecutor(job, e.opts...)
		if err != nil {
			return err
		}

		if err = ex.Start(ctx); err != nil {
			return err
		}
	}

	return nil
}
