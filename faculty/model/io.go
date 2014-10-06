// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package model

import (
	// "fmt"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

/*
	{ In *, Out *, _ * }
*/
type IO struct{}

func (IO) Spark(*be.Matter, ...interface{}) Value {
	return IO{}
}

func (IO) Cognize(eye *be.Eye, v interface{}) {
	eye.Show("In", v)
}

func (IO) CognizeIn(eye *be.Eye, v interface{}) {}

func (IO) CognizeOut(eye *be.Eye, v interface{}) {
	eye.Show(DefaultValve, v)
}
