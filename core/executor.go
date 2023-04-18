package core

<<<<<<< HEAD
type Runnable interface {
	Run()
}

type Executor interface {
=======
type IExecutor interface {
>>>>>>> f24c43d52ed967623debc4a5d7d96eba06784f0c
	Run()
	Add(task Runnable)
}

type DefaultExecutor struct {
	tasks []*Runnable
	Executor
}

func (executor *DefaultExecutor) New() *DefaultExecutor {
	executor.tasks = []*Runnable{}
	return executor
}

func (executor *DefaultExecutor) Run() {
}
func (executor *DefaultExecutor) Add(task *Runnable) {
	executor.tasks = append(executor.tasks, task)
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
		names = append(names, task.(Runnable).GetName())
	}
	return names
}

var Exec = NewDFExecutor("default")
