// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package memory

import (
	// "container/list"
	"sync"

	. "github.com/gocircuit/escher/circuit"
)

// I see forward. I think back. I see that I think. I think that I see. Thinking and seeing are not apart.

type Memory interface {
	Lookup(Address) Meaning
	Goto(...Name) Memory

	Circuit() Circuit
	Include(Name, Meaning) Meaning
	Exclude(Name) Meaning
	Link(u, v Vector)
	Unlink(u, v Vector)
}

// memory is a hierarchical namespace of circuits.
type memory struct{
	*sync.Mutex
	x Circuit
}

func NewMemory() Memory {
	var lk sync.Mutex
	return &memory{
		Mutex: &lk,
		x: New(),
	}
}

func (m *memory) Lookup(addr Address) Meaning {
	m.Lock()
	defer m.Unlock()
	return Copy(m.x.Goto(addr.Path()...))
}

func (m *memory) Goto(gate ...Name) Memory {
	m.Lock()
	defer m.Unlock()
	return &memory{
		Mutex: m.Mutex,
		x: m.x.Goto(gate...).(Circuit),
	}
}

func (m *memory) Circuit() Circuit {
	m.Lock()
	defer m.Unlock()
	return m.x.Clone()
}

func (m *memory) Include(n Name, v Meaning) Meaning {
	m.Lock()
	defer m.Unlock()
	r, _ := m.x.Include(n, Copy(v))
	return r
}

func (m *memory) Exclude(n Name) Meaning {
	m.Lock()
	defer m.Unlock()
	r, _ := m.x.Exclude(n)
	return r
}

func (m *memory) Link(u, v Vector) {
	m.Lock()
	defer m.Unlock()
	m.x.Link(u, v)
}

func (m *memory) Unlink(u, v Vector) {
	m.Lock()
	defer m.Unlock()
	m.x.Unlink(u, v)
}
