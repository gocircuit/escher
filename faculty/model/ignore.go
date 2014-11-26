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
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Register(be.NewMaterializer(IgnoreValves{}), "model", "IgnoreValves")
}

type IgnoreValves struct{}

func (IgnoreValves) Spark(*be.Eye, Circuit, ...interface{}) Value {
	return nil
}

func (IgnoreValves) CognizeCircuit(eye *be.Eye, v interface{}) {
	u := v.(Circuit).Copy()
	n := u.Unify("ignoreValves")
	u.Gate[n] = NewVerbAddress("*", "Ignore")
	u.Reflow(Super, n)
	eye.Show(DefaultValve, u)
}

func (IgnoreValves) Cognize(eye *be.Eye, v interface{}) {}
