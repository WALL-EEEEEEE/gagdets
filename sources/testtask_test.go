package sources

import (
	"testing"

	. "github.com/WALL-EEEEEEE/Axiom/core"
	"github.com/WALL-EEEEEEE/gagdets/pipes"
)

// TestTestTask calls executor to run TestTask, checking
// for a valid return value.
func TestTestTask(t *testing.T) {
	test_task := NewTestTask()
	std_pipe := pipes.NewStdPipe()
	test_task.Chain(&std_pipe)
	Exec.Add(&test_task)
	Exec.Start()
}
