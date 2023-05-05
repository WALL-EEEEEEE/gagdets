package pipes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/WALL-EEEEEEE/gagdets/core"
	"github.com/WALL-EEEEEEE/gagdets/sources"
	log "github.com/sirupsen/logrus"
)

type JsonPipe struct {
	core.Pipe
}

func NewJsonPipe() StdPipe {
	return StdPipe{Pipe: core.NewPipe("JsonPipe")}
}

func (pipe *JsonPipe) Run(collector *core.Collector) {
	cnt := 0
	data := []sources.Topic{}
	for item := range *collector {
		log.Debugf("JsonPipe: %v -> %v", collector, item.(sources.Topic).Content)
		data = append(data, item.(sources.Topic))
		cnt += 1
	}
	if len(data) < 1 {
		log.Infof("No any items.")
		return
	}
	json_data, _ := json.MarshalIndent(data, "", " ")
	output_name := fmt.Sprintf("data-%s.json", time.Now().Format(time.UnixDate))
	_ = ioutil.WriteFile(output_name, json_data, 0644)
	log.Infof("Write %d item into %s", cnt, output_name)

}

func init() {
	jsonPipe := NewJsonPipe()
	core.Reg.Register(&jsonPipe)
}
