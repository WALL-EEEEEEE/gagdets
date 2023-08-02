package sources

import (
	"encoding/json"
	"os"

	. "github.com/WALL-EEEEEEE/Axiom/core"
	"github.com/WALL-EEEEEEE/gagdets/items"
	log "github.com/sirupsen/logrus"
)

type TestTask struct {
	Task
}

func NewTestTask() TestTask {
	return TestTask{Task: NewTask("TestTask")}
}

func (task *TestTask) Run(collector *Collector) {
	var topics []items.Topic
	log.Infof("Start Task %s", task.GetName())
	json_file := "/mnt/d/Project/go/gagdets/data-Fri Apr 21 17:03:33 CST 2023.json"
	json_data, err := os.ReadFile(json_file)
	if err != nil {
		log.Errorf("Failed to load %s (%s)!", json_file, err)
	}
	err = json.Unmarshal(json_data, &topics)
	if err != nil {
		log.Errorf("Failed to load %s (Invalid json content: %s)!", json_file, err)
	}
	for _, topic := range topics {
		log.Infof("Find Topic: %+v", topic)
		*collector <- topic
	}

}

func init() {
	testTask := NewTestTask()
	Reg.Register(&testTask)
}
