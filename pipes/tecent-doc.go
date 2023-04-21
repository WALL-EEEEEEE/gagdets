package pipes

import (
	"github.com/WALL-EEEEEEE/gagdets/core"
	"github.com/WALL-EEEEEEE/gagdets/sources"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

var tecent_link_api = "https://api.hiflow.tencent.com/engine/webhook/31/1646403589125844994"

func TecentDocPipe(collector chan interface{}) {

	cnt := 1
	data := []sources.Topic{}
	for item := range collector {
		log.Debugf("JsonPipe: %v -> %v", collector, item.(sources.Topic).Content)
		data = append(data, item.(sources.Topic))
		cnt += 1
	}
	client := resty.New()
	resp, err := client.R().SetBody(data).
		Post(tecent_link_api)
	if err != nil {
		log.Error(err)
	}

	log.Infof("Sync %d topics to tencent doc!", cnt)
}

func init() {
	core.Exec.AddPipe(TecentDocPipe)
}
