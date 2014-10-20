// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

import (
	"strings"
	"sync"

	. "github.com/gocircuit/escher/circuit"
	. "github.com/gocircuit/escher/kit/memory"
	"github.com/gocircuit/escher/see"
)

type Reservoir struct {
	sync.Mutex
	Memory
}

func NewReservoir() *Reservoir {
	return &Reservoir{Memory: Memory(New())}
}

func (r *Reservoir) Dump() Circuit {
	r.Lock()
	defer r.Unlock()
	return r.Memory.Dump()
}

func (r *Reservoir) Fetch(addr Address) Value {
	r.Lock()
	defer r.Unlock()
	return r.Memory.Lookup(addr)
}

func (r *Reservoir) Save(addr Address, value Value) Value {
	r.Lock()
	defer r.Unlock()
	return r.Memory.Save(addr, value)
}

func (r *Reservoir) Register(name string, v interface{}) {
	r.Lock()
	defer r.Unlock()
	r.Memory.Save(see.ParseAddr(name), v)
}
