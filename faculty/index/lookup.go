// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package index

import (
	// "fmt"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

type Lookup struct{}

func (Lookup) Spark(*be.Eye, Circuit, ...interface{}) Value {
	return nil
}

func (Lookup) CognizeView(eye *be.Eye, v interface{}) {
	u := v.(Circuit)
	x := u.CircuitAt("Index")
	addr := u.VerbAt("Address")
	r := be.AsIndex(x).Recall(addr.Address()...)
	if r == nil {
		eye.Show("NotFound", New().Grow("NotFound", Circuit(addr)).Grow("In", x))
	} else {
		eye.Show("Found", r)
	}
}

func (Lookup) CognizeNotFound(eye *be.Eye, v interface{}) {}

func (Lookup) CognizeFound(eye *be.Eye, v interface{}) {}
