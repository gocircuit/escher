// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package basic

import (
	"fmt"

	"github.com/gocircuit/escher/tree"
	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Root.AddTerminal("ignore", think.Ignore{})
	faculty.Root.AddTerminal("show", Println{})
}

// Println
type Println struct{}

func (Println) Materialize() think.Reflex {
	s, t := think.NewSynapse()
	go func() {
		s.Attach(func(v interface{}) { println(fmt.Sprintf("%v", v)) })
	}()
	return think.Reflex{"": t}
}
