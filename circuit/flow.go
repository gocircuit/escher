// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	// "fmt"
	"log"
)

// Link connects two different, yet unconnected valves (by vector),
// potentially from the same gate
func (u Circuit) Link(x, y Vector) {
	if x.Gate == y.Gate && x.Valve == y.Valve {
		panic("self loop")
	}
	xs, ys := u.valves(x.Gate), u.valves(y.Gate)
	if z, ok := xs[x.Valve]; ok && !Same(z, y) {
		log.Fatalf("%v:%v already connected to %v, not %v", x.Gate, x.Valve, z, y)
		panic("contra")
	}
	if z, ok := ys[y.Valve]; ok && !Same(z, x) {
		log.Fatalf("%v:%v already connected to %v, not %v", y.Gate, y.Valve, z, x)
		panic("contra")
	}
	xs[x.Valve], ys[y.Valve] = y, x
}

func (u Circuit) valves(p Name) map[Name]Vector {
	if _, ok := u.Flow[p]; !ok {
		u.Flow[p] = make(map[Name]Vector)
	}
	return u.Flow[p]
}

// Unlink removes the link between two Vectors
func (u Circuit) Unlink(x, y Vector) {
	xs, ys := u.Flow[x.Gate], u.Flow[y.Gate]
	delete(xs, x.Valve)
	delete(ys, y.Valve)
	if len(xs) == 0 {
		delete(u.Flow, x.Gate)
	}
	if len(ys) == 0 {
		delete(u.Flow, y.Gate)
	}
}

// Valves returns the list of connected valve-name
// of the gate with the supplied name,
// and their connected vectors
func (u Circuit) Valves(gate Name) map[Name]Vector {
	return u.Flow[gate]
}

// ValveNames returns the list of connected valve-names
// of the gate with the supplied name
func (u Circuit) ValveNames(gate Name) []Name {
	var r []Name
	for n := range u.Flow[gate] {
		r = append(r, n)
	}
	return r
}

// Degree returns the number of connected valves
// of the gate with the supplied name
func (u Circuit) Degree(gate Name) int {
	return len(u.Flow[gate])
}

// View returns a copy of this circuit reduced to the supplied gate
func (u Circuit) View(gate Name) Circuit {
	x := New()
	for vlv, vec := range u.Flow[gate] {
		x.Include(vlv, u.At(vec.Gate))
	}
	return x
}

// Follow returns the vector linked up with the supplied vector
func (u Circuit) Follow(v Vector) Vector {
	return u.Flow[v.Gate][v.Valve]
}

func (u Circuit) Flows() (r [][2]Vector) {
	for xname, xview := range u.Flow {
		for xvalve, xvec := range xview {
			r = append(r, [2]Vector{{xname, xvalve}, xvec})
		}
	}
	return
}

// Vol counts the number of connected valves within this circuit
func (u Circuit) Vol() (vol int) {
	for _, view := range u.Flow {
		for range view {
			vol++
		}
	}
	return
}
