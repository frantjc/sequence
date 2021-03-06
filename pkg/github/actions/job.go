package actions

import "net/url"

type Job struct {
	Name        string      `json:"name,omitempty" yaml:"name,omitempty"`
	Permissions interface{} `json:"permissions,omitempty" yaml:"permissions,omitempty"`
	Needs       interface{} `json:"needs,omitempty" yaml:"needs,omitempty"`
	If          interface{} `json:"if,omitempty" yaml:"if,omitempty"`
	RunsOn      string      `json:"runs-on,omitempty" yaml:"runs-on,omitempty"`
	Environment *struct {
		Name string   `json:"name,omitempty" yaml:"name,omitempty"`
		URL  *url.URL `json:"url,omitempty" yaml:"url,omitempty"`
	} `json:"environment,omitempty" yaml:"environment,omitempty"`
	Concurrency interface{}       `json:"concurrency,omitempty" yaml:"concurrency,omitempty"`
	Outputs     map[string]string `json:"outputs,omitempty" yaml:"outputs,omitempty"`
	Env         map[string]string `json:"env,omitempty" yaml:"env,omitempty"`
	Container   interface{}       `json:"container,omitempty" yaml:"container,omitempty"`
	Steps       []*Step           `json:"steps,omitempty" yaml:"steps,omitempty"`
}
