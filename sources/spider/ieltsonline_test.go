package spider

import (
	"testing"

	. "github.com/WALL-EEEEEEE/Axiom/core"
	"github.com/WALL-EEEEEEE/gagdets/pipes"
)

// TestSixtone calls executor to run Sixtone, checking
// for a valid return value.
func TestIELTSONLINE(t *testing.T) {
	/*
		name := "Gladys"
		want := regexp.MustCompile(`\b` + name + `\b`)
		msg, err := Hello("Gladys")
		if !want.MatchString(msg) || err != nil {
			t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
		}
	*/
	iets_spider := NewIETSSpider()
	std_pipe := pipes.NewStdPipe()
	iets_spider.Chain(&std_pipe)
	Exec.Add(&iets_spider)
	Exec.Start()
}
