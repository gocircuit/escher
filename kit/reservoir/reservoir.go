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

func NewReservoir(u ...Circuit) Reservoir {
	if len(u) == 0 {
		return &reservoir{Memory: Memory(New())}
	}
	if len(u) == 1 {
		return &reservoir{Memory: Memory(DeepCopy(u[0]).(Circuit))}
	}
	panic(1)
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
