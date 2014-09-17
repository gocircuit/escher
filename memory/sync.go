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

// Memory is a synchronized memory.
type Memory struct{
	sync.Mutex
	__ *memory
}

func NewMemory() *Memory {
	return &Memory{__: newMemory()}
}

func (m *Memory) StartHijack() Circuit {
	m.Lock()
	return m.__.root
}

func (m *Memory) EndHijack() {
	m.Unlock()
}

func (m *Memory) Yield() Circuit {
	m.Lock()
	defer m.Unlock()
	return m.__.Yield()
}

// View
func (m *Memory) View() Circuit {
	m.Lock()
	defer m.Unlock()
	return m.__.View()
}

// Restart
func (m *Memory) Restart() Circuit {
	m.Lock()
	defer m.Unlock()
	return m.__.Restart()
}

// Lookup
func (m *Memory) Lookup(name ...Name) Meaning {
	m.Lock()
	defer m.Unlock()
	return m.__.Lookup(name)
}

// Step
func (m *Memory) Step(gate Name) (Circuit, Address) {
	m.Lock()
	defer m.Unlock()
	return m.__.Step(gate)
}

// Jump
func (m *Memory) Jump(gate ...Name) Circuit {
	m.Lock()
	defer m.Unlock()
	return m.__.Jump(gate...)
}

// Include
func (m *Memory) Include(n Name, x Meaning) Circuit {
	m.Lock()
	defer m.Unlock()
	return m.__.Include(n, x)
}

// Exclude
func (m *Memory) Exclude(n Name) Circuit {
	m.Lock()
	defer m.Unlock()
	return m.__.Exclude(n)
}

// Link
func (m *Memory) Link(u, v Vector) Circuit {
	m.Lock()
	defer m.Unlock()
	return m.__.Link(u, v)
}

// Unlink
func (m *Memory) Unlink(u, v Vector) Circuit {
	m.Lock()
	defer m.Unlock()
	return m.__.Unlink(u, v)
}
