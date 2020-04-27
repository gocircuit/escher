// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package os

import (
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/hoijui/escher/be"
	cir "github.com/hoijui/escher/circuit"
	"github.com/hoijui/escher/faculty"
	fio "github.com/hoijui/escher/faculty/io"
)

func Init(arg []string) {
	faculty.Register(be.NewSource(argCircuit(arg)), "os", "Arg")
	faculty.Register(be.NewSource(argCircuit(os.Environ())), "os", "Env")
	faculty.Register(be.NewMaterializer(Exit{}), "os", "Exit")
	faculty.Register(be.NewMaterializer(Fatal{}), "os", "Fatal")
	faculty.Register(be.NewMaterializer(LookPath{}), "os", "LookPath")
	faculty.Register(be.NewMaterializer(Join{}), "os", "Join")
	faculty.Register(be.NewMaterializer(Process{}), "os", "Process")
	faculty.Register(fio.NewReaderMaterializer(os.Stdin), "os", "Stdin")
	faculty.Register(fio.NewWriterMaterializer(os.Stdout), "os", "Stdout")
	faculty.Register(fio.NewWriterMaterializer(os.Stderr), "os", "Stderr")
}

func argCircuit(arg []string) cir.Circuit {
	r := cir.New()
	for i, a := range arg {
		r.Include(i, a)
	}
	return r
}

// Exit
type Exit struct{ be.Sparkless }

func (Exit) OverCognize(eye *be.Eye, name cir.Name, value interface{}) {
	switch t := value.(type) {
	case int:
		os.Exit(t)
	default:
		os.Exit(0)
	}
}

// Fatal
type Fatal struct{ be.Sparkless }

func (Fatal) OverCognize(eye *be.Eye, name cir.Name, value interface{}) {
	log.Fatalf("%v", value)
}

// LookPath
type LookPath struct{ be.Sparkless }

func (LookPath) CognizeName(eye *be.Eye, value interface{}) {
	p, err := exec.LookPath(value.(string))
	if err != nil {
		log.Fatalf("no file path to %s", value.(string))
	}
	eye.Show(cir.DefaultValve, p)
}

func (LookPath) Cognize(eye *be.Eye, value interface{}) {}

// Join
type Join struct{ be.Sparkless }

func (Join) CognizeView(eye *be.Eye, v interface{}) {
	u := v.(cir.Circuit)
	var s []string
	for _, n := range u.SortedNames() {
		s = append(s, u.Gate[n].(string))
	}
	eye.Show(cir.DefaultValve, path.Join(s...))
}

func (Join) Cognize(*be.Eye, interface{}) {}
