// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package reservoir

import (
	"sync"

	. "github.com/gocircuit/escher/circuit"
	. "github.com/gocircuit/escher/kit/memory"
)

type Reservoir interface {
	Put(addr Address, value Value) Value
	Get(addr Address) Value
	Forget(Address) Value
}

type reservoir struct {
	sync.Mutex
	Memory
}

func NewReservoir() Reservoir {
	return &reservoir{Memory: Memory(New())}
}

func (r *reservoir) Get(addr Address) Value {
	r.Lock()
	defer r.Unlock()
	return r.Memory.Get(addr)
}

func (r *reservoir) Forget(addr Address) Value {
	r.Lock()
	defer r.Unlock()
	return r.Memory.Forget(addr)
}

func (r *reservoir) Put(addr Address, value Value) Value {
	r.Lock()
	defer r.Unlock()
	return r.Memory.Put(addr, value)
}
