package pipes

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"

	. "github.com/WALL-EEEEEEE/Axiom/core"
	"github.com/WALL-EEEEEEE/gagdets/items"
	log "github.com/sirupsen/logrus"
)

type NewsPipe struct {
	Pipe
}

var default_outdir = "out"

func NewNewsPipe() NewsPipe {
	return NewsPipe{Pipe: NewPipe("NewsPipe")}
}

func (pipe *NewsPipe) Run(collector *Collector) {
	cnt := 0
	for item := range *collector {
		title := item.(items.NewsItem).Title
		log.Debugf("FilePipe: %v -> %v", collector, title)
		output_item_file := path.Join(default_outdir, fmt.Sprintf("%s.json", title))
		json_data, _ := json.MarshalIndent(item, "", " ")
		_, err := os.Stat(default_outdir)
		if os.IsNotExist(err) {
			err = os.MkdirAll(default_outdir, fs.FileMode(os.O_RDWR))
			if err != nil {
				log.Errorln(err)
				continue
			}
		}
		err = ioutil.WriteFile(output_item_file, json_data, 0644)
		if err != nil {
			log.Errorln(err)
			continue
		}
		log.Infof("Pipe news item %s into file %s", title, output_item_file)
		cnt += 1
	}
	log.Infof("NewsPipe Write news items: %d into directory %s", cnt, default_outdir)

}

func init() {
	newsPipe := NewNewsPipe()
	Reg.Register(&newsPipe)
}
