// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	"log"
)

// Name is one of: int or string
type Name interface{}

// Meaning is one of: see.Address, string, int, float64, complex128, Circuit
type Meaning interface{}

// Super is a placeholder meaning for the super image
type Super struct{}

func (Super) String() string {
	return "*"
}

// circuit ...
type circuit struct {
	image map[Name]Meaning
	real map[Name]map[Name]Real // image -> valve -> opposing image and valve
}

type Circuit struct {
	*circuit
}

func New() Circuit {
	return Circuit{newCircuit()}
}

func newCircuit() *circuit {
	return &circuit{
		image: make(map[Name]Meaning),
		real: make(map[Name]map[Name]Real),
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
	return len(u.image) == 0 && len(u.real) == 0
}

func (c *circuit) Letters() []string {
	var l []string
	for key, _ := range c.image {
		if s, ok := key.(string); ok {
			l = append(l, s)
		}
	}
	return l
}

func (c *circuit) Numbers() []int {
	var l []int
	for key, _ := range c.image {
		if i, ok := key.(int); ok {
			l = append(l, i)
		}
	}
	return l
}

func (u *circuit) Images() map[Name]Meaning {
	return u.image
}

func (u *circuit) Reals() map[Name]map[Name]Real {
	return u.real
}

func (u *circuit) String() string {
	return u.Print("", "\t")
}

func (u *circuit) Seal(name Name) {
	if _, ok := u.Include(name, Super{}); ok {
		panic("overwriting super")
	}
	for nm, y := range u.Images() {
		if y == nil {
			log.Fatalf("nil peer: %v", nm)
		}
	}
}
