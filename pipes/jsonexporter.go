package pipes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	. "github.com/WALL-EEEEEEE/Axiom/core"
	. "github.com/WALL-EEEEEEE/gagdets/core"
	. "github.com/WALL-EEEEEEE/gagdets/items"
	log "github.com/sirupsen/logrus"
)

type JsonPipe struct {
	Pipe
}

func NewJsonPipe() StdPipe {
	return StdPipe{Pipe: NewPipe("JsonPipe")}
}

func (pipe *JsonPipe) Run() {
	cnt := 0
	data := []Topic{}
	output_stream := pipe.GetOutputStream().Out()
	for item := range output_stream {
		log.Debugf("JsonPipe: %v -> %v", output_stream, item.(Topic).Content)
		data = append(data, item.(Topic))
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
