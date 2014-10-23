// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	. "github.com/gocircuit/escher/circuit"
)

// CleanUp removes nil-valued gates and their incident edges.
// CleanUp never returns nil.
func CleanUp(u Circuit) Value {
	for n, g := range u.Gate {
		if g != nil {
			continue
		}
		delete(u.Gate, n)
		for vlv, vec := range u.Flow[n] {
			u.Unlink(Vector{n, vlv}, vec)
		}
	}
	return u
}
