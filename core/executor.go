package core

type IExecutor interface {
	Run()
}

type IRunnable interface {
	Run()
}

type Runnable struct {
	IRunnable
	name string
}

func (runnable Runnable) GetName() string {
	return runnable.name
}
func (runnable Runnable) Run() {
}

type DefaultExecutor struct {
	IExecutor
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
