package core

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
}

func NewDFExecutor(name string) DefaultExecutor {
	collector := make(Collector, 100)
	exec := DefaultExecutor{name: name, collector: &collector}
	return exec
}

func (executor *DefaultExecutor) Start() {
	defer close(*executor.collector)
	for _, task := range executor.tasks {
		go task.Run(executor.collector)
	}
	executor.Output()
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
		go pipe(executor.collector)
	}
}

var Exec = NewDFExecutor("default")
