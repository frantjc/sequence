package workflow

import (
	"fmt"
	"io"
	"net/url"
	"os"

	"gopkg.in/yaml.v3"
)

func NewJobFromFile(name string) (*Job, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	return NewJobFromReader(f)
}

func NewJobFromReader(r io.Reader) (*Job, error) {
	j := &Job{}
	d := yaml.NewDecoder(r)
	return j, d.Decode(j)
}

type Job struct {
	Name        string            `json:"name,omitempty" yaml:"name,omitempty"`
	Permissions interface{}       `json:"permissions,omitempty" yaml:"permissions,omitempty"`
	Needs       interface{}       `json:"needs,omitempty" yaml:"needs,omitempty"`
	If          interface{}       `json:"if,omitempty" yaml:"if,omitempty"`
	RunsOn      string            `json:"runs-on,omitempty" yaml:"runs-on,omitempty"`
	Environment *JobEnvironment   `json:"environment,omitempty" yaml:"environment,omitempty"`
	Concurrency interface{}       `json:"concurrency,omitempty" yaml:"concurrency,omitempty"`
	Outputs     map[string]string `json:"outputs,omitempty" yaml:"outputs,omitempty"`
	Env         map[string]string `json:"env,omitempty" yaml:"env,omitempty"`
	Container   interface{}       `json:"container,omitempty" yaml:"container,omitempty"`
	Steps       []Step            `json:"steps,omitempty" yaml:"steps,omitempty"`
}

type JobEnvironment struct {
	Name string   `json:"name,omitempty" yaml:"name,omitempty"`
	URL  *url.URL `json:"url,omitempty" yaml:"url,omitempty"`
}

type Concurrency struct {
	Group            string `json:"group,omitempty" yaml:"group,omitempty"`
	CancelInProgress bool   `json:"cancel-in-progress,omitempty" yaml:"cancel-in-progress,omitempty"`
}

type Defaults struct {
	Run *Run `json:"run,omitempty" yaml:"run,omitempty"`
}

type Run struct {
	Shell            string `json:"shell,omitempty" yaml:"shell,omitempty"`
	WorkingDirectory string `json:"working-directory,omitempty" yaml:"working-directory,omitempty"`
}

type Container struct {
	Image string `json:"image,omitempty" yaml:"image,omitempty"`
}

func (j *Job) GetStep(id string) (*Step, error) {
	for _, step := range j.Steps {
		if step.GetID() == id {
			return &step, nil
		}
	}
	return nil, fmt.Errorf("job has no step with name or id %s", id)
}
