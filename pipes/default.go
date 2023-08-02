package pipes

import (
	. "github.com/WALL-EEEEEEE/Axiom/core"
	"github.com/WALL-EEEEEEE/gagdets/items"
	log "github.com/sirupsen/logrus"
)

type StdPipe struct {
	Pipe
}

func NewStdPipe() StdPipe {
	return StdPipe{Pipe: NewPipe("StdPipe")}
}

func (pipe *StdPipe) Run(collector *Collector) {
	topic_cnt := 0
	for item := range *collector {
		log.Debugf("StdPipe: %v -> %v", collector, item.(items.Topic).Content)
		topic_cnt += 1
	}
	log.Infof("StdPipe: %d topics", topic_cnt)
}

func init() {
	stdPipe := NewStdPipe()
	Reg.Register(&stdPipe)
}
