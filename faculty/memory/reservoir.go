// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package memory

import (
	"sync"

	. "github.com/gocircuit/escher/circuit"
)

type Memory interface {
	Put(addr Address, value Value) Value
	Get(addr Address) Value
	Forget(Address) Value
}

type memory struct {
	sync.Mutex
	Circuit
}

func NewMemory(u ...Circuit) Memory {
	if len(u) == 0 {
		return &memory{Circuit: New()}
	}
	if len(u) == 1 {
		return &memory{Circuit: DeepCopy(u[0]).(Circuit)}
	}
	panic(1)
}

func (r *memory) Get(addr Address) Value {
	r.Lock()
	defer r.Unlock()
	return r.Circuit.Lookup(addr)
}

func (r *memory) Forget(addr Address) Value {
	r.Lock()
	defer r.Unlock()
	return r.Circuit.Forget(addr)
}

func (r *memory) Put(addr Address, value Value) Value {
	r.Lock()
	defer r.Unlock()
	return r.Circuit.Place(addr, value)
}
