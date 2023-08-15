package sources

import (
	. "github.com/WALL-EEEEEEE/Axiom/core"
	. "github.com/WALL-EEEEEEE/gagdets/items"
	. "github.com/WALL-EEEEEEE/gagdets/pipes"
	log "github.com/sirupsen/logrus"
)

type TestTask struct {
	Task
}

func NewTestTask() TestTask {
	task := TestTask{Task: NewTask("TestTask")}
	task.ITask = &task
	return task
}

func (task *TestTask) Run() {
	var topics []Topic = []Topic{
		{
			Content: "test",
		},
	}
	log.Infof("Start Task %s", task.GetName())
	for _, topic := range topics {
		log.Infof("Find Topic: %+v", topic)
		in := task.GetInputStream()
		(&in).In() <- topic
	}

}

func init() {
	testTask := NewTestTask()
	stdPipe := NewStdPipe()
	testTask.Chain(&stdPipe)
	Reg.Register(&testTask)
}
