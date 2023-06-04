package pipes

import (
<<<<<<< HEAD
	"github.com/WALL-EEEEEEE/gagdets/core"
=======
	. "github.com/WALL-EEEEEEE/Axiom/core"
>>>>>>> aa185959b593dc2e181ba7238de572c30881c8d8
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
