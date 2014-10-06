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

func (u Circuit) Copy() Irreducible {
	return u.Clone()
}

func (x Circuit) Same(r Irreducible) bool {
	y, ok := r.(Circuit)
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
		if uv == wv { // shallow comparison of gates is default Go-level comparison
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
			_ugate, _uvlv := uvec.Reduce()
			_wgate, _wvlv := wvec.Reduce()
			if _ugate != _wgate || _uvlv != _wvlv { // shallow comparison, at Go value level
				return false
			}
		}
	}
	return true
}
