// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	"fmt"
)

// Vector ...
type Vector Circuit

func NewVector(gate, valve Name) Vector {
	return Vector(New().Grow("Gate", gate).Grow("Valve", valve))
}

func (v Vector) Reduce() (gate, valve Name) {
	return Circuit(v).At("Gate"), Circuit(v).At("Valve")
}

func (v Vector) Gate() Name {
	return Circuit(v).At("Gate")
}

func (v Vector) Valve() Name {
	return Circuit(v).At("Valve")
}

func (v Vector) Copy() Reducible {
	return Vector(Circuit(v).Clone())
}

func (v Vector) String() string {
	g, u := v.Reduce()
	return fmt.Sprintf("%v:%v", g, u)
}

func (v Vector) Same(x Reducible) bool {
	w, ok := x.(Vector)
	if !ok {
		return false
	}
	return Same(Circuit(v), Circuit(w))
}

//
func (c *circuit) Link(x, y Vector) {
	xg, xv := x.Reduce()
	yg, yv := y.Reduce()
	if xg == yg && xv == yv {
		panic("self loop")
	}
	xs, ys := c.valves(xg), c.valves(yg)
	if _, ok := xs[xv]; ok {
		panic("dup")
	}
	if _, ok := ys[yv]; ok {
		panic("dup")
	}
	xs[xv], ys[yv] = y, x
}

func (c *circuit) Unlink(x, y Vector) {
	xg, xv := x.Reduce()
	yg, yv := y.Reduce()
	xs, ys := c.flow[xg], c.flow[yg]
	delete(xs, xv)
	delete(ys, yv)
	if len(xs) == 0 {
		delete(c.flow, xg)
	}
	if len(ys) == 0 {
		delete(c.flow, yg)
	}
}

func (c *circuit) valves(p Name) map[Name]Vector {
	if _, ok := c.flow[p]; !ok {
		c.flow[p] = make(map[Name]Vector)
	}
	return c.flow[p]
}

func (u *circuit) Valves(gate Name) map[Name]Vector {
	return u.flow[gate]
}

func (c *circuit) Follow(v Vector) Vector {
	return c.flow[v.Gate()][v.Valve()]
}
