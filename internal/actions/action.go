package actions

import (
	"io"

	"gopkg.in/yaml.v2"
)

func NewFromBytes(b []byte) (*Action, error) {
	a := &Action{}
	return a, yaml.Unmarshal(b, a)
}

func NewFromReader(r io.Reader) (*Action, error) {
	a := &Action{}
	d := yaml.NewDecoder(r)
	return a, d.Decode(a)
}

func NewFromString(s string) (*Action, error) {
	return NewFromBytes([]byte(s))
}

type Action struct {
	Name        string   `json:",omitempty"`
	Author      string   `json:",omitempty"`
	Description string   `json:",omitempty"`
	Inputs      *Inputs  `json:",omitempty"`
	Outputs     *Outputs `json:",omitempty"`
	Runs        *Runs    `json:",omitempty"`
}

type Inputs map[string]struct {
	Description        string `json:",omitempty"`
	Required           bool   `json:",omitempty"`
	Default            string `json:",omitempty"`
	DeprecationMessage string `json:",omitempty"`
}

type Outputs map[string]struct {
	Description string `json:",omitempty"`
}

type Runs struct {
	Plugin     string            `json:",omitempty"`
	Using      string            `json:",omitempty"`
	Main       string            `json:",omitempty"`
	Image      string            `json:",omitempty"`
	Entrypoint string            `json:",omitempty"`
	Args       []string          `json:",omitempty"`
	Env        map[string]string `json:",omitempty"`
}