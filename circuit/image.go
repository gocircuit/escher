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
		w.symbol[n] = y // copyMeaning(y)
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
	switch t := x.(type) {
	case Circuit:
		return t.Copy()
	}
	return x
}
