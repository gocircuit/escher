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

	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

func Init(sourceDir string, arg []string) {
	faculty.Register("os.Arg", be.NewNoun(argCircuit(arg)))
	faculty.Register("os.SourceDir", be.NewNoun(sourceDir))
	faculty.Register("os.Env", Env{})
	faculty.Register("os.Exit", Exit{})
	faculty.Register("os.Fatal", Fatal{})
	faculty.Register("os.Stdin", Stdin{})
	faculty.Register("os.Stdout", Stdout{})
	faculty.Register("os.Stderr", Stderr{})
	//
	faculty.Register("os.LookPath", LookPath{})
	faculty.Register("os.Process", Process{})
}

func argCircuit(arg []string) Circuit {
	r := New()
	for i, a := range arg {
		r.Include(i, a)
	}
	return r
}

// Env
type Env struct{}

func (Env) Materialize() (be.Reflex, Value) {
	reflex, _ := be.NewEyeCognizer(
		func(eye *be.Eye, valve Name, value interface{}) {
			if valve != "Name" {
				return
			}
			n, ok := value.(string)
			if !ok {
				panic("non-string name perceived by os.env")
			}
			ev := os.Getenv(n)
			log.Printf("Environment %s=%s", n, ev)
			eye.Show("Value", ev)
		},
		"Name", "Value",
	)
	return reflex, Env{}
}

// Exit
type Exit struct{}

func (Exit) Materialize() (be.Reflex, Value) {
	reflex, _ := be.NewEyeCognizer(
		func(eye *be.Eye, valve Name, value interface{}) {
			switch t := value.(type) {
			case int:
				os.Exit(t)
			default:
				os.Exit(0)
			}
		}, 
		DefaultValve,
	)
	return reflex, Exit{}
}

// Fatal
type Fatal struct{}

func (Fatal) Materialize() (be.Reflex, Value) {
	reflex, _ := be.NewEyeCognizer(
		func(eye *be.Eye, valve Name, value interface{}) {
			log.Fatalln(value)
		}, 
		DefaultValve,
	)
	return reflex, Fatal{}
}

// LookPath
type LookPath struct{}

func (LookPath) Materialize() (be.Reflex, Value) {
	reflex, _ := be.NewEyeCognizer(
		func(eye *be.Eye, valve Name, value interface{}) {
			if valve != "Name" {
				return
			}
			p, err := exec.LookPath(value.(string))
			if err != nil {
				log.Fatalf("no file path to %s", value.(string))
			}
			eye.Show(DefaultValve, p)
		},
		"Name", DefaultValve,
	)
	return reflex, LookPath{}
}
