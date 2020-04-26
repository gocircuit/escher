// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"github.com/hoijui/escher/a"
	cir "github.com/hoijui/escher/circuit"
)

type Carry struct {
	cir.Name
	cir.Value
}

func SeeLink(src *a.Src, nsugar int) (x []cir.Vector, carry []*Carry) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	x, carry = make([]cir.Vector, 2), make([]*Carry, 2)
	t := src.Copy()
	Space(t)
	//
	g, p, v, ok := seeEndpoint(t)
	if !ok {
		return nil, nil
	}
	if g != nil {
		carry[0] = &Carry{nsugar, g}
		x[0] = cir.Vector{nsugar, cir.DefaultValve}
		nsugar++
	} else {
		x[0] = cir.Vector{p, v}
	}
	//
	a.Whitespace(t)
	t.Match("=")
	a.Whitespace(t)
	//
	g, p, v, ok = seeEndpoint(t)
	if !ok {
		return nil, nil
	}
	if g != nil {
		carry[1] = &Carry{nsugar, g}
		x[1] = cir.Vector{nsugar, cir.DefaultValve}
		nsugar++
	} else {
		x[1] = cir.Vector{p, v}
	}
	//
	if !Space(t) { // require newline at end
		return nil, nil
	}
	src.Become(t)
	return
}

func seeEndpoint(src *a.Src) (m cir.Value, p, v cir.Name, ok bool) {
	if p, v, ok = seeNameEndpoint(src); ok { // valve (or empty string)
		return
	}
	m, ok = seeValueEndpoint(src) // int, string, etc.
	return
}

func seeValueEndpoint(src *a.Src) (m cir.Value, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			ok = false
		}
	}()
	t := src.Copy()
	m = SeeValue(t)
	src.Become(t)
	return m, true
}

func seeNameEndpoint(src *a.Src) (gate, valve cir.Name, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			ok = false
		}
	}()
	t := src.Copy()
	gate = SeeValue(t)
	t.Match(string(a.ValveSelector))
	valve = SeeValue(t)
	src.Become(t)
	ok = true
	return
}
