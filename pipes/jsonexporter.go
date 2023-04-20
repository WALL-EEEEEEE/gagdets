package pipes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/WALL-EEEEEEE/gagdets/core"
	log "github.com/sirupsen/logrus"
)

func JsonPipe(collector chan interface{}) {
	cnt := 1
	data := []core.Item{}
	for item := range collector {
		log.Infof("JsonPipe: %v -> %v", collector, item)
		data = append(data, item)
		cnt += 1
	}
	json_data, _ := json.MarshalIndent(data, "", " ")
	output_name := fmt.Sprintf("data-%s.json", time.Now().Format(time.UnixDate))
	_ = ioutil.WriteFile(output_name, json_data, 0644)
	log.Infof("Write %d item into %s", cnt, output_name)
}

func init() {
	core.Exec.AddPipe(JsonPipe)
}