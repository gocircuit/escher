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
	u := value.(Circuit)
	for name, _ := range u.SortedNames() {
		eye.Show(DefaultValve, New().Grow("Name", name).Grow("Value", u.At(name)))
	}
	eye.Show(DefaultValve, Term)
}

func (Yield) Cognize(eye *be.Eye, value interface{}) {}
