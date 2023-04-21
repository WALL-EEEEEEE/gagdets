package pipes

import (
	"github.com/WALL-EEEEEEE/gagdets/core"
	"github.com/WALL-EEEEEEE/gagdets/sources"
	log "github.com/sirupsen/logrus"
)

func StdPipe(collector chan interface{}) {
	topic_cnt := 0
	for item := range collector {
		log.Debugf("StdPipe: %v -> %v", collector, item.(sources.Topic).Content)
		topic_cnt += 1
	}
	log.Infof("StdPipe: %d topics", topic_cnt)
}

func init() {
	core.Exec.AddPipe(StdPipe)
}
