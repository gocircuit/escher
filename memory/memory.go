// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package memory

import (
	// "container/list"

	. "github.com/gocircuit/escher/circuit"
)

// memory is a hierarchical namespace of circuits.
type memory struct{
	root Circuit
	seeing Circuit // output and current focus
}

func newMemory() *memory {
	root := New()
	return &memory{
		root: root,
		seeing: root,
	}
}

// Restart
func (m *memory) Restart() Circuit {
	m.seeing = m.root
	return m.seeing
}

// Step
func (m *memory) Step(gate Name) Circuit {
	a, ok := m.seeing.At(gate).(Address)
	if !ok {
		panic("cannot enter non-addresses")
	}
	return m.Jump(a.Path()...)
}

// Lookup
func (m *memory) Lookup(gate ...Name) Meaning {
	var x Circuit = m.root
	for i, g := range gate {
		if i+1 == len(gate) {
			return x.At(g)
		}
		x = x.CircuitAt(g)
	}
	return x
}

// Jump
func (m *memory) Jump(gate ...Name) Circuit {
	m.seeing = m.Lookup(gate...).(Circuit)
	return m.seeing
}

// Plumbing

func (m *memory) Refine(n Name) Circuit {
	m.Include(n, New())
	return m.Step(n)
}

// Include
func (m *memory) Include(n Name, x Meaning) Circuit {
	if _, over := m.seeing.Include(n, x); over {
		panic("over including")
	}
	return m.seeing
}

// Exclude
func (m *memory) Exclude(n Name) Circuit {
	if _, forgotten := m.seeing.Exclude(n); !forgotten {
		panic("nothing to exclude")
	}
	return m.seeing
}

// Link
func (m *memory) Link(u, v Vector) Circuit {
	m.seeing.Link(u, v)
	return m.seeing
}

// Unlink
func (m *memory) Unlink(u, v Vector) Circuit {
	m.seeing.Unlink(u, v)
	return m.seeing
}
