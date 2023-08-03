package core

import (
	. "github.com/WALL-EEEEEEE/Axiom/core"
)

type IPipe interface {
	ITask
}

type Pipe struct {
	Task
	name string
}

func NewPipe(name string) Pipe {
	return Pipe{name: name}
}

func (pipe Pipe) GetName() string {
	return pipe.name
}

func (pipe Pipe) GetType() []ServType {
	return []ServType{PIPE}
}
