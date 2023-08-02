package pipes

import (
	. "github.com/WALL-EEEEEEE/Axiom/core"
	"github.com/WALL-EEEEEEE/gagdets/items"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

var tecent_link_api = "https://api.hiflow.tencent.com/engine/webhook/31/1646403589125844994"

type TecentDocPipe struct {
	Pipe
}

func NewTecentDocPipe() StdPipe {
	return StdPipe{Pipe: NewPipe("TecentDocPipe")}
}

func (pipe *TecentDocPipe) Run(collector *Collector) {
	cnt := 1
	data := []items.Topic{}
	for item := range *collector {
		log.Debugf("TecentDocPipe: %v -> %v", collector, item.(items.Topic).Content)
		data = append(data, item.(items.Topic))
		cnt += 1
	}
	client := resty.New()
	var result struct {
	}
	_, err := client.R().SetBody(data).SetResult(result).
		Post(tecent_link_api)
	if err != nil {
		log.Error(err)
	}

	log.Infof("Sync %d topics to tencent doc!", cnt)
}

func init() {
	tecentDocPipe := NewTecentDocPipe()
	Reg.Register(&tecentDocPipe)
}
