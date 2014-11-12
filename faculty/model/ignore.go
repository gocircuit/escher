// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package model

import (
	// "fmt"

	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/circuit"
)

func init() {
	faculty.Register("model.IgnoreValves", be.NewNativeMaterializer(IgnoreValves{}))
}

type IgnoreValves struct{}

func (IgnoreValves) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	return nil
}

func (IgnoreValves) CognizeCircuit(eye *be.Eye, v interface{}) {
	u := v.(Circuit).Copy()
	n := u.Unify("ignoreValves")
	u.Gate[n] = NewAddress("Ignore")
	u.Reflow(Super, n)
	eye.Show(DefaultValve, u)
}

func (IgnoreValves) Cognize(eye *be.Eye, v interface{}) {}
