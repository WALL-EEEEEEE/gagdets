package pipes

import (
	"github.com/WALL-EEEEEEE/gagdets/core"
	"github.com/WALL-EEEEEEE/gagdets/sources"
	log "github.com/sirupsen/logrus"
)

type StdPipe struct {
	core.Pipe
}

func NewStdPipe() StdPipe {
	return StdPipe{Pipe: core.NewPipe("StdPipe")}
}

func (pipe *StdPipe) Run(collector *core.Collector) {
	topic_cnt := 0
	for item := range *collector {
		log.Debugf("StdPipe: %v -> %v", collector, item.(sources.Topic).Content)
		topic_cnt += 1
	}
	log.Infof("StdPipe: %d topics", topic_cnt)
}

func init() {
	stdPipe := NewStdPipe()
	core.Reg.Register(&stdPipe)
}
