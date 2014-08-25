// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	"fmt"

	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/be"
)

func Init(n string) {
	name = n
	faculty.Root.AddTerminal("Ignore", be.Ignore{})
	faculty.Root.AddTerminal("Show", Println{})
	faculty.Root.AddTerminal("See", Scanln{})
	faculty.Root.AddTerminal("Name", be.NewNounReflex(name))
}

var name string

// Name returns the name assigned to this program execution
func Name() string {
	return name
}

// Scanln
type Scanln struct{}

func (Scanln) Materialize() be.Reflex {
	s, t := be.NewSynapse()
	go func() {
		r := s.Focus(be.DontCognize)
		go func() {
			for {
				var em string
				fmt.Scanln(&em)
				r.ReCognize(em)
			}
		}()
	}()
	return be.Reflex{"_": t}
}

// Println
type Println struct{}

func (Println) Materialize() be.Reflex {
	s, t := be.NewSynapse()
	go func() {
		s.Focus(
			func(v interface{}) {
				fmt.Printf("%v\n", v)
			},
		)
	}()
	return be.Reflex{"_": t}
}
