// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	// "log"
	"sync"

	. "github.com/gocircuit/escher/circuit"
)

// Defer is a materializer which holds on to a reflex and connects it on materialization.
// Defers can be materialized only once. Repeat materializations will return a nil reflex and value.
type Defer struct {
	sync.Mutex
	Reflex
	Value
}

func NewDefer(r Reflex, v Value) Materializer {
	return &Defer{Reflex: r, Value: v}
}

func (d *Defer) Materialize(*Matter) (r Reflex, v Value) {
	d.Lock()
	defer d.Unlock()
	r, v = d.Reflex, d.Value
	d.Reflex, d.Value = nil, nil
	return
}
