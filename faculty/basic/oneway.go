// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

type OneWay struct{}

func (OneWay) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	return nil
}

func (OneWay) CognizeFrom(eye *be.Eye, value interface{}) {
	eye.Show("To", value)
}

func (OneWay) CognizeTo(eye *be.Eye, value interface{}) {}
