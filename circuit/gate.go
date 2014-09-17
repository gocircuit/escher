// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	// "fmt"
)

// Convenience access

func (u Circuit) IntOrZeroAt(name Name) int {
	i, ok := u.circuit.OptionAt(name)
	if !ok {
		return 0
	}
	return i.(int)
}

//
func (u Circuit) CircuitAt(name Name) Circuit {
	return u.circuit.At(name).(Circuit)
}

func (u Circuit) CircuitOptionAt(name Name) (Circuit, bool) {
	v, ok := u.OptionAt(name)
	if ok {
		return v.(Circuit), ok
	}
	return Circuit{}, false
}

// int
func (u Circuit) IntAt(name Name) int {
	return u.circuit.At(name).(int)
}

func (u Circuit) IntOptionAt(name Name) (int, bool) {
	v, ok := u.OptionAt(name)
	if ok {
		return v.(int), ok
	}
	return 0, false
}

//
func (u Circuit) StringAt(name Name) string {
	return u.circuit.At(name).(string)
}

func (u Circuit) StringOptionAt(name Name) (string, bool) {
	v, ok := u.OptionAt(name)
	if ok {
		return v.(string), ok
	}
	return "", false
}

//

func (u Circuit) AddressAt(name Name) Address {
	return u.circuit.At(name).(Address)
}

func (u Circuit) AddressOptionAt(name Name) (Address, bool) {
	v, ok := u.OptionAt(name)
	if ok {
		return v.(Address), ok
	}
	return nil, false
}

// Series-application methods

func (u Circuit) ReGrow(name string, meaning Meaning) Circuit {
	u.circuit.Include(name, meaning)
	return u
}

func (u Circuit) Grow(name string, meaning Meaning) Circuit {
	if _, over := u.circuit.Include(name, meaning); over {
		panic("over writing")
	}
	return u
}

func (u Circuit) Refine(name string) Circuit {
	r := New()
	u.Grow(name, r)
	return r
}

func (u Circuit) Abandon(name string) Circuit {
	u.circuit.Exclude(name)
	return u
}

func (u Circuit) Rename(x, y Name) Circuit {
	m, ok := u.circuit.Exclude(x)
	if !ok {
		panic("np")
	}
	if _, over := u.circuit.Include(y, m); over {
		panic("over")
	}
	return u
}

func (u Circuit) Lookup(gate ...Name) Meaning {
	x := u
	for i, g := range gate {
		if i+1 == len(gate) {
			return x.At(g)
		}
		x = x.CircuitAt(g)
	}
	return x
}

// Low-level

func (u *circuit) Include(name Name, meaning Meaning) (before Meaning, overwrite bool) {
	before, overwrite = u.gate[name]
	u.gate[name] = meaning
	return
}

func (u *circuit) Exclude(name Name) (meaning Meaning, forgotten bool) {
	meaning, forgotten = u.gate[name]
	delete(u.gate, name)
	return
}

// Len returns the number of gates.
func (u *circuit) Len() int {
	return len(u.gate)
}

func (c *circuit) OptionAt(name Name) (Meaning, bool) {
	v, ok := c.gate[name]
	return v, ok
}

func (c *circuit) At(name Name) Meaning {
	return c.gate[name]
}

func (c *circuit) Super() Name {
	for n, m := range c.gate {
		if _, ok := m.(Super); ok {
			return n
		}
	}
	return nil
}
