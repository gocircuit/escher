// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

func CopyMeaning(x Meaning) Meaning {
	switch t := x.(type) {
	case Circuit:
		return t.Copy()
	}
	return x
}

func (u Circuit) Copy() Circuit {
	if u.circuit == nil {
		return Circuit{}
	}
	return Circuit{u.circuit.Copy()}
}

func (u *circuit) Copy() *circuit {
	w := newCircuit()
	// images
	for n, y := range u.image {
		w.image[n] = CopyMeaning(y)
	}
	// reals
	for n, z := range u.real {
		x := make(map[Name]Real)
		w.real[n] = x
		for a, b := range z {
			x[a] = b
		}
	}
	return w
}

func SameMeaning(x, y Meaning) bool {
	xc, x_ := x.(Circuit)
	yc, y_ := y.(Circuit)
	if x_ && y_ {
		return Same(xc, yc)
	}
	return x == y
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
	// image
	for n, y := range u.image {
		yy, ok := w.image[n]
		if !ok {
			return false
		}
		if !SameMeaning(y, yy) {
			return false
		}
	}
	// real
	for n, z := range u.real {
		zz, ok := w.real[n]
		if !ok {
			return false
		}
		for v, m := range z {
			mm, ok := zz[v]
			if !ok {
				return false
			}
			if !SameReal(m, mm) {
				return false
			}
		}
	}
	return true
}
