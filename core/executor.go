package core

type Runnable interface {
	Run()
}

type Executor interface {
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
