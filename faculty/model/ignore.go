// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package model

import (
	"github.com/gocircuit/escher/be"
	cir "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Register(be.NewMaterializer(IgnoreValves{}), "model", "IgnoreValves")
}

type IgnoreValves struct{}

func (IgnoreValves) Spark(*be.Eye, cir.Circuit, ...interface{}) cir.Value {
	return nil
}

func (IgnoreValves) CognizeCircuit(eye *be.Eye, v interface{}) {
	u := v.(cir.Circuit).Copy()
	n := u.Unify("ignoreValves")
	u.Gate[n] = cir.NewVerbAddress("*", "Ignore")
	u.Reflow(cir.Super, n)
	eye.Show(cir.DefaultValve, u)
}

func (IgnoreValves) Cognize(eye *be.Eye, v interface{}) {}
