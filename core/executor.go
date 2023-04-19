package core

type Collector chan interface{}

type IExecutor interface {
	Run()
	GetName() string
	Add(task IRunnable)
	List() []string
	Output(callback func(collector *Collector))
}

type IRunnable interface {
	Run(collector *Collector)
	GetName() string
}

type DefaultExecutor struct {
	name      string
	tasks     []IRunnable
	collector *Collector
}

func NewDFExecutor(name string) DefaultExecutor {
	collector := make(Collector, 100)
	exec := DefaultExecutor{name: name, collector: &collector}
	return exec
}

func (executor *DefaultExecutor) Run() {
	defer close(*executor.collector)
	for _, task := range executor.tasks {
		task.Run(executor.collector)
	}
}

func (executor *DefaultExecutor) Add(task IRunnable) {
	executor.tasks = append(executor.tasks, task)
}

func (executor *DefaultExecutor) List() []string {
	var names []string
	for _, task := range executor.tasks {
		names = append(names, task.GetName())
	}
	return names
}
func (executor *DefaultExecutor) Output(callback func(collector *Collector)) {
	callback(executor.collector)
}

var Exec = NewDFExecutor("default")
