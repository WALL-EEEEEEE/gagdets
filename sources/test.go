package sources

import (
	"encoding/json"
	"os"

	"github.com/WALL-EEEEEEE/gagdets/core"
	log "github.com/sirupsen/logrus"
)

type TestTask struct {
	core.Task
}

func NewTestTask() TestTask {
	return TestTask{Task: core.NewTask("TestTask")}
}

func (task *TestTask) Run(collector *core.Collector) {
	var topics []Topic
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
	core.Reg.Register(&testTask)
}
