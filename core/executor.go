package core

import (
	"sync"
)

type Collector chan interface{}
type IPipe func(ch chan interface{})

type IExecutor interface {
	Start()
	GetName() string
	Add(task IRunnable)
	AddPipe(pipe IPipe)
	List() []string
}

type IRunnable interface {
	Run(collector chan interface{})
	GetName() string
}

type DefaultExecutor struct {
	name      string
	tasks     []IRunnable
	pipes     []IPipe
	collector chan interface{}
	broker    *Broker[interface{}]
}

func NewDFExecutor(name string) DefaultExecutor {
	broker := NewBroker[interface{}]()
	exec := DefaultExecutor{name: name, collector: broker.publishCh, broker: broker}
	return exec
}

func (executor *DefaultExecutor) startTask(wg *sync.WaitGroup) {
	_start := func(task IRunnable) {
		defer func() {
			wg.Done()
		}()
		task.Run(executor.broker.publishCh)
	}
	for _, task := range executor.tasks {
		wg.Add(1)
		go _start(task)
	}
}

func (executor *DefaultExecutor) startPipe(wg *sync.WaitGroup) {
	_start := func(pipe IPipe) {
		defer func() {
			wg.Done()
		}()
		collector := executor.broker.Subscribe()
		pipe(collector)
	}
	for _, pipe := range executor.pipes {
		wg.Add(1)
		go _start(pipe)
	}
}

func (executor *DefaultExecutor) Start() {
	go executor.broker.Start()
	var task_wg sync.WaitGroup
	var pipe_wg sync.WaitGroup
	executor.startTask(&task_wg)
	executor.startPipe(&pipe_wg)
	task_wg.Wait()
	executor.broker.Close()
	pipe_wg.Wait()
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

var Exec = NewDFExecutor("default")
