// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

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

func (c *circuit) Follow(p, v Name) (q, u Name) {
	x, ok := c.valves(p)[v]
	if !ok {
		return nil, nil
	}
	return x.To()
}
