package pipes

import (
	. "github.com/WALL-EEEEEEE/Axiom/core"
	. "github.com/WALL-EEEEEEE/gagdets/core"
	. "github.com/WALL-EEEEEEE/gagdets/items"
	log "github.com/sirupsen/logrus"
)

type StdPipe struct {
	Pipe
}

func NewStdPipe() StdPipe {
	std_pipe := StdPipe{Pipe: NewPipe("StdPipe")}
	std_pipe.ITask = &std_pipe
	return std_pipe
}

func (pipe *StdPipe) Run() {
	topic_cnt := 0
	in_stream := pipe.GetOutputStream()
	for {
		item, ok := in_stream.Read()
		if !ok {
			break
		}
		log.Infof("StdPipe: %v -> %v", in_stream, item.(Topic).Content)
		topic_cnt += 1
	}
	log.Infof("StdPipe: %d topics", topic_cnt)
}

func init() {
	stdPipe := NewStdPipe()
	Reg.Register(&stdPipe)
}
