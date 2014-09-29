// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	// "fmt"
)

// Circuit is irreducible

func (u Circuit) Clone() Circuit {
	w := New()
	for n, m := range u.Gate {
		w.Gate[n] = Copy(m) // deep copy
	}
	for g, h := range u.Flow {
		x := make(map[Name]Vector)
		w.Flow[g] = x
		for a, b := range h {
			x[a] = b
		}
	}
	return w
}

func (u Circuit) Copy() Irreducible {
	return u.Clone()
}

func (x Circuit) Same(r Irreducible) bool {
	y, ok := r.(Circuit)
	if !ok {
		return false
	}
	return x.IsContainedIn(y) && y.IsContainedIn(x)
}

func (u Circuit) IsContainedIn(w Circuit) bool {
	// gate
	for g, y := range u.Gate {
		yy, ok := w.Gate[g]
		if !ok {
			return false
		}
		if !Same(y, yy) {
			return false
		}
	}
	// flow
	for g, h := range u.Flow {
		hh, ok := w.Flow[g]
		if !ok {
			return false
		}
		for a, b := range h {
			bb, ok := hh[a]
			if !ok {
				return false
			}
			if !Same(b, bb) {
				return false
			}
		}
	}
	return true
}
