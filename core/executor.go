package core

import "sync"

type Collector chan interface{}
type IPipe func(collector *Collector)

type IExecutor interface {
	Start()
	GetName() string
	Add(task IRunnable)
	AddPipe(pipe IPipe)
	List() []string
}

type IRunnable interface {
	Run(collector *Collector)
	GetName() string
}

type DefaultExecutor struct {
	name      string
	tasks     []IRunnable
	pipes     []IPipe
	collector *Collector
	broker    *Broker[interface{}]
}

func NewDFExecutor(name string) DefaultExecutor {
	broker := NewBroker[interface{}]()
	collector := Collector(broker.publishCh)
	exec := DefaultExecutor{name: name, collector: &collector, broker: broker}
	return exec
}

func (executor *DefaultExecutor) Start() {
	go executor.broker.Start()
	defer executor.broker.Close()
	var wg sync.WaitGroup
	_start := func(task IRunnable) {
		defer wg.Done()
		task.Run(executor.collector)
	}
	for _, task := range executor.tasks {
		wg.Add(1)
		go _start(task)
	}
	executor.Output()
	wg.Wait()
}

func (executor *DefaultExecutor) Add(task IRunnable) {
	executor.tasks = append(executor.tasks, task)
}

func (executor *DefaultExecutor) AddPipe(pipe IPipe) {
	executor.pipes = append(executor.pipes, pipe)
}
func (executor *DefaultExecutor) List() []string {
	var names []string
	for _, task := range executor.tasks {
		names = append(names, task.GetName())
	}
	return names
}
func (executor *DefaultExecutor) Output() {
	for _, pipe := range executor.pipes {
		collector := Collector(executor.broker.Subscribe())
		go pipe(&collector)
	}
}

var Exec = NewDFExecutor("default")
