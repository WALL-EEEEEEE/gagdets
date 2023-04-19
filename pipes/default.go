package pipes

import (
	"github.com/WALL-EEEEEEE/gagdets/core"
	log "github.com/sirupsen/logrus"
)

func StdPipe(collector *core.Collector) {
	for item := range *collector {
		log.Info(item)
	}
}

func init() {
	core.Exec.AddPipe(StdPipe)
}
