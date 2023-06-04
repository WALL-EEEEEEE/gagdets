<<<<<<< HEAD
package sources

import (
	"testing"

	"github.com/WALL-EEEEEEE/gagdets/core"
	"github.com/WALL-EEEEEEE/gagdets/pipes"
)

// TestSixtone calls executor to run Sixtone, checking
// for a valid return value.
func TestSixtone(t *testing.T) {
	/*
		name := "Gladys"
		want := regexp.MustCompile(`\b` + name + `\b`)
		msg, err := Hello("Gladys")
		if !want.MatchString(msg) || err != nil {
			t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
		}
	*/
	sixtone_spider := NewSixtoneSpider()
	std_pipe := pipes.NewStdPipe()
	core.Exec.Add(&sixtone_spider)
	core.Exec.AddPipe(&std_pipe)
	core.Exec.Start()
}
=======
package sources

import (
	"testing"

	. "github.com/WALL-EEEEEEE/Axiom/core"
	"github.com/WALL-EEEEEEE/gagdets/pipes"
)

// TestSixtone calls executor to run Sixtone, checking
// for a valid return value.
func TestSixtone(t *testing.T) {
	/*
		name := "Gladys"
		want := regexp.MustCompile(`\b` + name + `\b`)
		msg, err := Hello("Gladys")
		if !want.MatchString(msg) || err != nil {
			t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
		}
	*/
	sixtone_spider := NewSixtoneSpider()
	std_pipe := pipes.NewStdPipe()
	Exec.Add(&sixtone_spider)
	Exec.AddPipe(&std_pipe)
	Exec.Start()
}
>>>>>>> aa185959b593dc2e181ba7238de572c30881c8d8
