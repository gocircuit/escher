// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package reservoir

import (
	. "github.com/gocircuit/escher/circuit"
)

func Restrict(r Reservoir, addr Address) Reservoir {
	return &restrict{r, addr}
}

type restrict struct {
	r Reservoir
	a Address
}

func (r *restrict) Put(addr Address, value Value) Value {
	return r.r.Put(r.a.Append(addr), value)
}

func (r *restrict) Get(addr Address) Value {
	return r.r.Get(r.a.Append(addr))
}

func (r *restrict) Forget(addr Address) Value {
	return r.r.Forget(r.a.Append(addr))
}
