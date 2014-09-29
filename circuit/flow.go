// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	// "fmt"
)

func (u Circuit) Link(x, y Vector) {
	xg, xv := x.Reduce()
	yg, yv := y.Reduce()
	if xg == yg && xv == yv {
		panic("self loop")
	}
	xs, ys := u.valves(xg), u.valves(yg)
	if _, ok := xs[xv]; ok {
		panic("dup")
	}
	if _, ok := ys[yv]; ok {
		panic("dup")
	}
	xs[xv], ys[yv] = y, x
}

func (u Circuit) valves(p Name) map[Name]Vector {
	if _, ok := u.Flow[p]; !ok {
		u.Flow[p] = make(map[Name]Vector)
	}
	return u.Flow[p]
}

func (u Circuit) Unlink(x, y Vector) {
	xg, xv := x.Reduce()
	yg, yv := y.Reduce()
	xs, ys := u.Flow[xg], u.Flow[yg]
	delete(xs, xv)
	delete(ys, yv)
	if len(xs) == 0 {
		delete(u.Flow, xg)
	}
	if len(ys) == 0 {
		delete(u.Flow, yg)
	}
}

func (u Circuit) Valves(gate Name) map[Name]Vector {
	return u.Flow[gate]
}

func (u Circuit) Follow(v Vector) Vector {
	g, h := v.Reduce()
	return u.Flow[g][h]
}
