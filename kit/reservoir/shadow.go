// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package reservoir

import (
	. "github.com/gocircuit/escher/circuit"
)

func Shadow(r Reservoir, sh Circuit) Reservoir {
	return &shadow{r, NewReservoir(sh)}
}

type shadow struct {
	r, s Reservoir
}

func (r *shadow) Put(addr Address, value Value) Value {
	return r.s.Put(addr, value)
}

func (r *shadow) Get(addr Address) Value {
	v := r.s.Get(addr)
	if v != nil {
		return v
	}
	return r.r.Get(addr)
}

func (r *shadow) Forget(addr Address) Value {
	return r.s.Forget(addr)
}
