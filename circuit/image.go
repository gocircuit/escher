// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

// import (
// 	"bytes"
// 	"fmt"
// )

func (u Circuit) Copy() Circuit {
	return Circuit{u.circuit.Copy()}
}

func (u *circuit) Copy() *circuit {
	w := newCircuit()
	// symbols
	for n, y := range u.symbol {
		w.symbol[n] = copyMeaning(y)
	}
	// matchings
	for n, z := range u.match {
		x := make(map[Name]Matching)
		w.match[n] = x
		for a, b := range z {
			x[a] = b
		}
	}
	return w
}

func copyMeaning(x Meaning) Meaning {
	return x
	// switch t := x.(type) {
	// case Circuit:
	// 	return t.Copy()
	// }
	// return x
}

func Same(x, y Circuit) bool {
	if x.circuit == nil && y.circuit == nil {
		return true
	}
	if x.circuit == nil || y.circuit == nil {
		return false
	}
	return x.circuit.Contains(y.circuit) && y.circuit.Contains(x.circuit)
}

func (u *circuit) Contains(w *circuit) bool {
	// symbol
	for n, y := range u.symbol {
		yy, ok := w.symbol[n]
		if !ok {
			return false
		}
		if y != yy {
			return false
		}
	}
	// match
	for n, z := range u.match {
		zz, ok := w.match[n]
		if !ok {
			return false
		}
		for v, m := range z {
			mm, ok := zz[v]
			if !ok {
				return false
			}
			if !SameMatching(m, mm) {
				return false
			}
		}
	}
	return true
}
