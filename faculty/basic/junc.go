// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"

	"github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/kit/plumb"
)

func init() {
	faculty.Root.AddTerminal("Junction", Junction{})
}

// Junction
type Junction struct{}

func (Junction) Materialize() Reflex {
	reflex, _ := NewEyeCognizer(cognizeJunction, "X", "Y", "Z")
	return reflex
}

func cognizeJunction(eye *Eye, dvalve string, dvalue interface{}) {
	ch := make(sparkChan, 2)
	switch dvalve {
	case "X":
		go spark(ch, eye, "Y", dvalue)
		go spark(ch, eye, "Z", dvalue)
	case "Y":
		go spark(ch, eye, "X", dvalue)
		go spark(ch, eye, "Z", dvalue)
	case "Z":
		go spark(ch, eye, "X", dvalue)
		go spark(ch, eye, "Y", dvalue)
	}
	<-ch
	<-ch
}

type sparkChan chan struct{}

func spark(ch sparkChan, eye *Eye, dvalve string, dvalue interface{}) {
	eye.Show(dvalve, dvalue)
	ch <- struct{}{}
}
