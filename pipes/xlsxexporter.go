package pipes

import (
	"github.com/WALL-EEEEEEE/gagdets/core"
	log "github.com/sirupsen/logrus"
)

type XlsxPipe struct {
	core.Pipe
}

func NewXlsxPipe() StdPipe {
	return StdPipe{Pipe: core.NewPipe("TecentDocPipe")}
}

func (pipe *XlsxPipe) Run(collector *core.Collector) {
	topic_cnt := 1
	for item := range *collector {
		log.Info(item)
		topic_cnt += 1
	}
	log.Infof("Found %d topics", topic_cnt)

}

func init() {
	xlsxPipe := NewXlsxPipe()
	core.Reg.Register(&xlsxPipe)
}
