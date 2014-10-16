// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	// "fmt"
)

// Circuit is an irreducible. An irreducible is an object with a Copy and Same methods.

func (u Circuit) Copy() Circuit {
	w := New()
	for n, v := range u.Gate {
		w.Gate[n] = v // shallow Go-value level copy of gate values
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

func (u Circuit) DeepCopy() Circuit {
	u = u.Copy()
	for n, v := range u.Gate {
		u.Gate[n] = DeepCopy(v)
	}
	return u
}

func (x Circuit) Same(v Value) bool {
	y, ok := v.(Circuit)
	if !ok {
		return false
	}
	return x.isWithin(y) && y.isWithin(x)
}

func (u Circuit) isWithin(w Circuit) bool {
	// gate
	for ug, uv := range u.Gate {
		wv, ok := w.Gate[ug]
		if !ok {
			return false
		}
		if !Same(uv, wv) {
			return false
		}
	}
	// flow
	for ugate, ufan := range u.Flow {
		wfan, ok := w.Flow[ugate]
		if !ok {
			return false
		}
		for uvlv, uvec := range ufan {
			wvec, ok := wfan[uvlv]
			if !ok {
				return false
			}
			if uvec.Gate != wvec.Gate || uvec.Valve != wvec.Valve { // shallow comparison, at Go value level
				return false
			}
		}
	}
	return true
}
