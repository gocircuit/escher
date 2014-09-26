// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package memory

import (
	// "fmt"
	// "container/list"
	"sync"

	. "github.com/gocircuit/escher/circuit"
)

// I see forward. I think back. I see that I think. I think that I see. Thinking and seeing are not apart.

type Memory interface {
	Lookup(Address) Value
	Goto(...Name) Memory
	Refine(Name) Memory

	Circuit() Circuit
	Include(Name, Value) Value
	IncludeIfNot(Name, Value) Value
	Exclude(Name) Value
	Link(u, v Vector)
	Unlink(u, v Vector)

	T(func(Circuit))
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

func (m *memory) Print(prefix, indent string) string {
	return "(Memory)"
}

func (m *memory) T(t func(Circuit))  {
	m.Lock()
	defer m.Unlock()
	t(m.x)
}

func (m *memory) Lookup(addr Address) Value {
	m.Lock()
	defer m.Unlock()
	names := addr.Path()
	k := len(names)-1
	return Copy(m.goto_(names[:k]...).(*memory).x.At(names[k]))
}

func (m *memory) Goto(gate ...Name) Memory {
	m.Lock()
	defer m.Unlock()
	return m.goto_(gate...)
}

func (m *memory) goto_(gate ...Name) Memory {
	for _, name := range gate {
		m = m.x.At(name).(*memory)
	}
	return m
}

func (m *memory) Refine(n Name) (k Memory) {
	m.Lock()
	defer m.Unlock()
	v, ok := m.x.OptionAt(n)
	if !ok {
		k = &memory{
			Mutex: m.Mutex,
			x: New(),
		}
		m.x.Include(n, k)
		return k
	}
	return v.(Memory)
}

func (m *memory) Circuit() Circuit {
	m.Lock()
	defer m.Unlock()
	return m.x.Clone()
}

func (m *memory) Include(n Name, v Value) Value {
	m.Lock()
	defer m.Unlock()
	r, _ := m.x.Include(n, Copy(v))
	return r
}

func (m *memory) IncludeIfNot(n Name, v Value) Value {
	m.Lock()
	defer m.Unlock()
	u, ok := m.x.OptionAt(n)
	if ok {
		return Copy(u)
	}
	m.x.Include(n, Copy(v))
	return v
}

func (m *memory) Exclude(n Name) Value {
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
