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

func (Lookup) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	return nil
}

func (Lookup) CognizeView(eye *be.Eye, v interface{}) {
	u := v.(Circuit)
	r := be.AsIndex(u.CircuitAt("Index")).Recall(u.AddressAt("Address").Path...)
	if r == nil {
		eye.Show("Error", "NotFound")
	} else {
		eye.Show(DefaultValve, r)
	}
}

func (Lookup) CognizeError(eye *be.Eye, v interface{}) {}

func (Lookup) Cognize(eye *be.Eye, v interface{}) {}
