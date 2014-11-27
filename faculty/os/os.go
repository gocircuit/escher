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

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/faculty"
)

func Init(arg []string) {
	faculty.Register(be.NewSource(argCircuit(arg)), "os", "Arg")
	faculty.Register(be.NewSource(argCircuit(os.Getenv())), "os", "Env")
	faculty.Register(be.NewMaterializer(Exit{}), "os", "Exit")
	faculty.Register(be.NewMaterializer(Fatal{}), "os", "Fatal")
	faculty.Register(be.NewMaterializer(LookPath{}), "os", "LookPath")
	faculty.Register(Stdin{}, "os", "Stdin")
	faculty.Register(Stdout{}, "os", "Stdout")
	faculty.Register(Stderr{}, "os", "Stderr")
	faculty.Register(Process{}, "os", "Process")
}

func argCircuit(arg []string) Circuit {
	r := New()
	for i, a := range arg {
		r.Include(i, a)
	}
	return r
}

// Exit
type Exit struct{}

func (Exit) Spark(*be.Eye, Circuit, ...interface{}) Value {
	return nil
}

func (Exit) OverCognize(eye *be.Eye, name Name, value interface{}) {
	switch t := value.(type) {
	case int:
		os.Exit(t)
	default:
		os.Exit(0)
	}
}

// Fatal
type Fatal struct{}

func (Fatal) Spark(*be.Eye, Circuit, ...interface{}) Value {
	return nil
}

func (Fatal) OverCognize(eye *be.Eye, name Name, value interface{}) {
	log.Fatalf("%v", value)
}

// LookPath
type LookPath struct{}

func (LookPath) Spark(*be.Eye, Circuit, ...interface{}) Value {
	return nil
}

func (LookPath) CognizeName(eye *be.Eye, value interface{}) {
	p, err := exec.LookPath(value.(string))
	if err != nil {
		log.Fatalf("no file path to %s", value.(string))
	}
	eye.Show(DefaultValve, p)
}

func (LookPath) Cognize(eye *be.Eye, value interface{}) {}
