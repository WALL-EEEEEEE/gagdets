package pipes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

<<<<<<< HEAD
	"github.com/WALL-EEEEEEE/gagdets/core"
=======
	. "github.com/WALL-EEEEEEE/Axiom/core"
>>>>>>> aa185959b593dc2e181ba7238de572c30881c8d8
	"github.com/WALL-EEEEEEE/gagdets/items"
	log "github.com/sirupsen/logrus"
)

type JsonPipe struct {
	Pipe
}

func NewJsonPipe() StdPipe {
	return StdPipe{Pipe: NewPipe("JsonPipe")}
}

func (pipe *JsonPipe) Run(collector *Collector) {
	cnt := 0
	data := []items.Topic{}
	for item := range *collector {
		log.Debugf("JsonPipe: %v -> %v", collector, item.(items.Topic).Content)
		data = append(data, item.(items.Topic))
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
	Reg.Register(&jsonPipe)
}
