package core

import (
	. "github.com/WALL-EEEEEEE/Axiom/core"
)

type Pipe struct {
	Task
}

func NewPipe(name string) Pipe {
	pipe := Pipe{Task: NewTask(name)}
	pipe.ITask = &pipe
	return pipe
}

func (pipe *Pipe) GetName() string {
	return pipe.Task.GetName()
}

func (spider *Pipe) GetType() []ServType {
	return []ServType{PIPE}
}
