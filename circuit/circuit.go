// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	// "log"
)

// Name is one of: int or string
type Name interface{}

// Meaning is one of: see.Address, string, int, float64, complex128, Circuit
type Meaning interface{}

// circuit ...
type circuit struct {
	gate map[Name]Meaning
	flow map[Name]map[Name]Vector // gate -> valve -> opposing gate and valve
}

type Circuit struct {
	*circuit
}

func New() Circuit {
	return Circuit{newCircuit()}
}

func newCircuit() *circuit {
	return &circuit{
		gate: make(map[Name]Meaning),
		flow: make(map[Name]map[Name]Vector),
	}
}

var Nil Circuit // the nil circuit

func (u *circuit) IsNil() bool {
	return u == nil
}

func (u *circuit) IsEmpty() bool {
	if u == nil {
		return true
	}
	return len(u.gate) == 0 && len(u.flow) == 0
}

func (c *circuit) Letters() []string {
	var l []string
	for key, _ := range c.gate {
		if s, ok := key.(string); ok {
			l = append(l, s)
		}
	}
	return l
}

func (c *circuit) Numbers() []int {
	var l []int
	for key, _ := range c.gate {
		if i, ok := key.(int); ok {
			l = append(l, i)
		}
	}
	return l
}

func (u *circuit) Gates() map[Name]Meaning {
	return u.gate
}

func (u *circuit) String() string {
	return u.Print("", "\t")
}
