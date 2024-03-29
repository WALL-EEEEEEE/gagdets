package pipes

import (
	. "github.com/WALL-EEEEEEE/Axiom/core"
	. "github.com/WALL-EEEEEEE/gagdets/core"
	log "github.com/sirupsen/logrus"
)

type XlsxPipe struct {
	Pipe
}

func NewXlsxPipe() XlsxPipe {
	return XlsxPipe{Pipe: NewPipe("TecentDocPipe")}
}

func (pipe *XlsxPipe) Run() {
	topic_cnt := 1
	output_stream := pipe.Pipe.GetOutputStream()
	for {
		item, ok := output_stream.Read()
		if !ok {
			break
		}
		log.Info(item)
		topic_cnt += 1
	}
	log.Infof("Found %d topics", topic_cnt)

}

func init() {
	xlsxPipe := NewXlsxPipe()
	Reg.Register(&xlsxPipe)
}
