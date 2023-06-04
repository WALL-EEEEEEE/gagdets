package pipes

import (
	. "github.com/WALL-EEEEEEE/Axiom/core"
	log "github.com/sirupsen/logrus"
)

type XlsxPipe struct {
	Pipe
}

func NewXlsxPipe() StdPipe {
	return StdPipe{Pipe: NewPipe("TecentDocPipe")}
}

func (pipe *XlsxPipe) Run(collector *Collector) {
	topic_cnt := 1
	for item := range *collector {
		log.Info(item)
		topic_cnt += 1
	}
	log.Infof("Found %d topics", topic_cnt)

}

func init() {
	xlsxPipe := NewXlsxPipe()
	Reg.Register(&xlsxPipe)
}
