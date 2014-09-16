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

// I see forward. I think back. I see that I think. I think that I see. Thinking and seeing are not apart.

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

// View
func (m *memory) View() Circuit {
	return m.seeing
}

// Restart
func (m *memory) Restart() Circuit {
	m.seeing = m.root
	return m.seeing
}

// Step
func (m *memory) Step(gate Name) (Circuit, Address) {
	a := m.seeing.At(gate).(Address)
	return m.Jump(a.Path()...), a
}

// Lookup
func (m *memory) Lookup(gate ...Name) Meaning {
	return m.root.Lookup(gate...)
}

// Jump
func (m *memory) Jump(gate ...Name) Circuit {
	m.seeing = m.root.Lookup(gate...).(Circuit)
	return m.seeing
}

// Plumbing

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
