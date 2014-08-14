// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	"fmt"

	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/faculty"
)

func Init(n string) {
	name = n
	faculty.Root.AddTerminal("ignore", think.Ignore{})
	faculty.Root.AddTerminal("show", Println{})
	faculty.Root.AddTerminal("see", Scanln{})
	faculty.Root.AddTerminal("name", think.NewNounReflex(name))
}

var name string

// Name returns the name assigned to this program execution
func Name() string {
	return name
}

// Scanln
type Scanln struct{}

func (Scanln) Materialize() think.Reflex {
	s, t := think.NewSynapse()
	go func() {
		r := s.Focus(think.DontCognize)
		go func() {
			for {
				var em string
				fmt.Scanln(&em)
				r.ReCognize(em)
			}
		}()
	}()
	return think.Reflex{"Question": t}
}

// Println
type Println struct{}

func (Println) Materialize() think.Reflex {
	s, t := think.NewSynapse()
	go func() {
		s.Focus(
			func(v interface{}) { 
				fmt.Printf("%v\n", v)
			},
		)
	}()
	return think.Reflex{"Answer": t}
}
