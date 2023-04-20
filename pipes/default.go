package pipes

import (
	"github.com/WALL-EEEEEEE/gagdets/core"
	log "github.com/sirupsen/logrus"
)

func StdPipe(collector *core.Collector) {
	topic_cnt := 1
	for item := range *collector {
		log.Infof("StdPipe: %v -> %v", *collector, item)
		log.Debug(item)
		topic_cnt += 1
	}
	log.Infof("Found %d topics", topic_cnt)
}

func init() {
	//core.Exec.AddPipe(StdPipe)
}
