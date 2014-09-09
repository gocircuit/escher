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

// Real ...
type Real struct {
	Image [2]Name
	Valve [2]Name
}

func SameReal(x, y Real) bool {
	return x == y
}

func (x Real) To() (image, valve Name) {
	return x.Image[1], x.Valve[1]
}

func (x Real) Reverse() Real {
	x.Image[0], x.Image[1] = x.Image[1], x.Image[0]
	x.Valve[0], x.Valve[1] = x.Valve[1], x.Valve[0]
	return x
}

// New ...
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
	return len(u.image) == 0 && len(u.real) == 0
}

// Change adds a image to this circuit.
func (u *circuit) Change(name Name, meaning Meaning) (before Meaning, overwrite bool) {
	before, overwrite = u.image[name]
	u.image[name] = meaning
	return
}

func (u *circuit) ChangeExclusive(name Name, meaning Meaning) {
	if _, over := u.Change(name, meaning); over {
		panic(1)
	}
}

// At ...
func (c *circuit) At(name Name) (Meaning, bool) {
	v, ok := c.image[name]
	return v, ok
}

func (c *circuit) AtNil(name Name) Meaning {
	return c.image[name]
}

func (u *circuit) Forget(name Name) Meaning {
	forgotten := u.image[name]
	delete(u.image, name)
	return forgotten
}

// Form ...
func (c *circuit) Form(x Real) {
	if x.Image[0] == x.Image[1] && x.Valve[0] == x.Valve[1] {
		panic("misreal")
	}
	p := []map[Name]Real{
		c.valves(x.Image[0]), 
		c.valves(x.Image[1]),
	}
	v := x.Valve
	if _, ok := p[0][v[0]]; ok {
		panic("dup")
	}
	if _, ok := p[1][v[1]]; ok {
		panic("dup")
	}
	p[0][v[0]], p[1][v[1]] = x, x.Reverse()
}

func (c *circuit) valves(p Name) map[Name]Real {
	if _, ok := c.real[p]; !ok {
		c.real[p] = make(map[Name]Real)
	}
	return c.real[p]
}

func (u *circuit) Valves(image Name) map[Name]Real {
	return u.real[image]
}

// Real ...
func (c *circuit) Follow(p, v Name) (q, u Name) {
	x, ok := c.valves(p)[v]
	if !ok {
		return nil, nil
	}
	return x.Image[1], x.Valve[1]
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
	u.ChangeExclusive(name, Super{})
	for nm, y := range u.Images() {
		if y == nil {
			log.Fatalf("nil peer: %v", nm)
		}
	}
}
