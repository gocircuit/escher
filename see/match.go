// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"fmt"

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
	g, p, v, ok := seeJoin(t)
	if !ok {
		return nil, nil
	}
	if g != nil {
		sugar := fmt.Sprintf("#%d", nsugar)
		carry[0] = &Carry{sugar, g}
		x[0] = Vector{sugar, DefaultValve}
	} else {
		x[0] = Vector{p, v}
	}
	//
	Whitespace(t)
	t.Match("=")
	Whitespace(t)
	//
	g, p, v, ok = seeJoin(t)
	if !ok {
		return nil, nil
	}
	if g != nil {
		sugar := fmt.Sprintf("#%d", nsugar+1)
		carry[1] = &Carry{sugar, g}
		x[1] = Vector{sugar, DefaultValve}
	} else {
		x[1] = Vector{p, v}
	}
	//
	if !Space(t) { // require newline at end
		return nil,nil
	}
	src.Become(t)
	return
}

func seeJoin(src *Src) (m Value, p, v Name, ok bool) {
	if p, v, ok = seeJoinAddress(src); ok { // valve (or empty string)
		return
	}
	m, ok = seeJoinValue(src) // int, string, etc.
	return
}

func seeJoinValue(src *Src) (m Value, ok bool) {
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

func seeJoinAddress(src *Src) (peer, valve Name, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			ok = false
		}
	}()
	t := src.Copy()
	p := SeeAddressOrEmpty(t).(Address)
	t.Match(string(ValveSelector))
	v := SeeAddressOrEmpty(t).(Address)
	src.Become(t)
	return p.Simplify(), v.Simplify(), true
}
