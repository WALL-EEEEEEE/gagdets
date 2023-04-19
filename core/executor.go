package core

type Runnable interface {
	Run()
}

type Executor interface {
	Run()
	GetName() string
	Add(task IRunnable)
	List() []string
}

type IRunnable interface {
	Run()
	GetName() string
}

type DefaultExecutor struct {
	name  string
	tasks []IRunnable
}

func NewDFExecutor(name string) DefaultExecutor {
	exec := DefaultExecutor{name: name}
	return exec
}

func (executor *DefaultExecutor) Run() {
	for _, task := range executor.tasks {
		task.Run()
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

var Exec = NewDFExecutor("default")
