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

type Yield struct{}

func (Yield) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	return nil
}

func (Yield) CognizeCircuit(eye *be.Eye, value interface{}) {
	for n, v := range value.(Circuit).Gate {
		eye.Show(DefaultValve, New().Grow("Name", n).Grow("Value", v))
	}
}

func (Yield) Cognize(eye *be.Eye, value interface{}) {}
