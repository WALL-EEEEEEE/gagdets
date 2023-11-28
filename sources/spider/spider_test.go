package spider

import (
	"testing"

	. "github.com/WALL-EEEEEEE/Axiom/core"
	. "github.com/WALL-EEEEEEE/gagdets/core"
	"github.com/WALL-EEEEEEE/gagdets/pipes"
	log "github.com/sirupsen/logrus"
)

type ExampleSpider struct {
	Spider
}

func NewExampleSpider() ExampleSpider {
	spider := ExampleSpider{Spider: NewSpider("ExampleSpider", []string{})}
	spider.ITask = &spider
	return spider
}

func (spider *ExampleSpider) Run() {
	log.Info("ExampleSpider")
}

func TestSpider(t *testing.T) {
	example := NewExampleSpider()
	std_pipe := pipes.NewStdPipe()
	example.Chain(&std_pipe)
	Exec.Add(&example)
	Exec.Start()
}
