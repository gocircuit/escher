// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

type Reducible interface {
	Copy() Reducible
	Same(Reducible) bool
}

func Copy(x Meaning) Meaning {
	switch t := x.(type) {
	case Reducible:
		return t.Copy()
	}
	return x
}

func Same(x, y Meaning) bool {
	xr, x_ := x.(Reducible)
	yr, y_ := y.(Reducible)
	if x_ && y_ {
		return xr.Same(yr)
	}
	return x == y
}

// Circuit is reducible

func (u Circuit) Copy() Circuit {
	if u.circuit == nil {
		return Circuit{}
	}
	return Circuit{u.circuit.copy()}
}

func (u *circuit) copy() *circuit {
	w := newCircuit()
	for n, m := range u.gate {
		w.gate[n] = Copy(m)
	}
	for g, h := range u.flow {
		x := make(map[Name]Vector)
		w.flow[g] = x
		for a, b := range h {
			x[a] = b
		}
	}
	return w
}

func (x Circuit) Same(y Circuit) bool {
	if x.circuit == nil && y.circuit == nil {
		return true
	}
	if x.circuit == nil || y.circuit == nil {
		return false
	}
	return x.circuit.isWithin(y.circuit) && y.circuit.isWithin(x.circuit)
}

func (u *circuit) isWithin(w *circuit) bool {
	// gate
	for g, y := range u.gate {
		yy, ok := w.gate[g]
		if !ok {
			return false
		}
		if !Same(y, yy) {
			return false
		}
	}
	// flow
	for g, h := range u.flow {
		hh, ok := w.flow[g]
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
