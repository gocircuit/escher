// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package index

import (
	"github.com/gocircuit/escher/be"
	cir "github.com/gocircuit/escher/circuit"
)

type Lookup struct{}

func (Lookup) Spark(*be.Eye, cir.Circuit, ...interface{}) cir.Value {
	return nil
}

func (Lookup) CognizeView(eye *be.Eye, v interface{}) {
	u := v.(cir.Circuit)
	x := u.CircuitAt("Index")
	addr := u.VerbAt("Address")
	r := be.AsIndex(x).Recall(addr.Address()...)
	if r == nil {
		eye.Show("NotFound", cir.New().Grow("NotFound", cir.Circuit(addr)).Grow("In", x))
	} else {
		eye.Show("Found", r)
	}
}

func (Lookup) CognizeNotFound(eye *be.Eye, v interface{}) {}

func (Lookup) CognizeFound(eye *be.Eye, v interface{}) {}
