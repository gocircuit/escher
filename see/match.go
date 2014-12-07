// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	// "fmt"

	. "github.com/gocircuit/escher/a"
	. "github.com/gocircuit/escher/circuit"
)

type Carry struct {
	Name
	Value
}

func SeeLink(src *Src, nsugar int) (x []Vector, carry []*Carry) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	x, carry = make([]Vector, 2), make([]*Carry, 2)
	t := src.Copy()
	Space(t)
	//
	g, p, v, ok := seeEndpoint(t)
	if !ok {
		return nil, nil
	}
	if g != nil {
		carry[0] = &Carry{nsugar, g}
		x[0] = Vector{nsugar, DefaultValve}
		nsugar++
	} else {
		x[0] = Vector{p, v}
	}
	//
	Whitespace(t)
	t.Match("=")
	Whitespace(t)
	//
	g, p, v, ok = seeEndpoint(t)
	if !ok {
		return nil, nil
	}
	if g != nil {
		carry[1] = &Carry{nsugar, g}
		x[1] = Vector{nsugar, DefaultValve}
		nsugar++
	} else {
		x[1] = Vector{p, v}
	}
	//
	if !Space(t) { // require newline at end
		return nil, nil
	}
	src.Become(t)
	return
}

func seeEndpoint(src *Src) (m Value, p, v Name, ok bool) {
	if p, v, ok = seeNameEndpoint(src); ok { // valve (or empty string)
		return
	}
	m, ok = seeValueEndpoint(src) // int, string, etc.
	return
}

func seeValueEndpoint(src *Src) (m Value, ok bool) {
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

func seeNameEndpoint(src *Src) (gate, valve Name, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			ok = false
		}
	}()
	t := src.Copy()
	gate = SeeValue(t)
	t.Match(string(ValveSelector))
	valve = SeeValue(t)
	src.Become(t)
	ok = true
	return
}
