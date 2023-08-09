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
	return StdPipe{Pipe: NewPipe("StdPipe")}
}

func (pipe *StdPipe) Run() {
	topic_cnt := 0
	out_stream := pipe.GetOutputStream().Out()
	for item := range out_stream {
		log.Debugf("StdPipe: %v -> %v", out_stream, item.(Topic).Content)
		topic_cnt += 1
	}
	log.Infof("StdPipe: %d topics", topic_cnt)
}

func init() {
	stdPipe := NewStdPipe()
	Reg.Register(&stdPipe)
}
