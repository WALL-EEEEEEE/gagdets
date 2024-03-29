package core

import (
	. "github.com/WALL-EEEEEEE/Axiom/core"
)

type Spider struct {
	Task
	Urls []string
}

func NewSpider(name string, urls []string) Spider {
	spider := Spider{Task: NewTask(name), Urls: urls}
	spider.ITask = &spider
	return spider
}

func (spider *Spider) GetName() string {
	return spider.Task.GetName()
}

func (spider *Spider) GetType() []ServType {
	return []ServType{SPIDER}
}
