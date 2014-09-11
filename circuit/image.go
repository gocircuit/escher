// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

// Convenience access

func (u Circuit) OptionalIntAt(name string) int {
	i, ok := u.circuit.At(name)
	if !ok {
		return 0
	}
	return i.(int)
}

func (u Circuit) CircuitAt(name string) Circuit {
	return u.circuit.AtNil(name).(Circuit)
}

func (u Circuit) StringAt(name string) string {
	return u.circuit.AtNil(name).(string)
}

func (u Circuit) AddressAt(name string) Address {
	return u.circuit.AtNil(name).(Address)
}

func (u Circuit) Grow(name string, meaning Meaning) Circuit {
	u.circuit.Include(name, meaning)
	return u
}

func (u Circuit) Abandon(name string) Circuit {
	u.circuit.Exclude(name)
	return u
}

// Low-level

func (u *circuit) Include(name Name, meaning Meaning) (before Meaning, overwrite bool) {
	before, overwrite = u.image[name]
	u.image[name] = meaning
	return
}

func (u *circuit) Exclude(name Name) Meaning {
	forgotten := u.image[name]
	delete(u.image, name)
	return forgotten
}

// Len returns the number of images.
func (u *circuit) Len() int {
	return len(u.image)
}

func (c *circuit) At(name Name) (Meaning, bool) {
	v, ok := c.image[name]
	return v, ok
}

func (c *circuit) AtNil(name Name) Meaning {
	return c.image[name]
}

func (c *circuit) MeaningAt(name Name) Meaning {
	v, ok := c.image[name]
	if !ok {
		panic(1)
	}
	return v
}
