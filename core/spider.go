package core

import (
	. "github.com/WALL-EEEEEEE/Axiom/core"
)

type Spider struct {
	Task
	Urls []string
}

func (*Spider) Run(collector *Collector) {
}

func NewSpider(name string, urls []string) Spider {
	spider := Spider{Task: NewTask(name), Urls: urls}
	return spider
}

func (spider *Spider) GetName() string {
	return spider.Task.GetName()
}

func (spider *Spider) GetType() []ServType {
	return []ServType{SPIDER}
}
