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
	Meaning
}

const DefaultValve = "_"

func SeeReal(src *Src, nsugar int) (x *Real, carry []*Carry) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	x, carry = &Real{}, make([]*Carry, 2)
	t := src.Copy()
	Space(t)
	//
	g, p, v, ok := seeJoin(t)
	if !ok {
		return nil, nil
	}
	if g != nil {
		sugar := fmt.Sprintf("sugar#%d", nsugar)
		carry[0] = &Carry{sugar, g}
		x.Image[0], x.Valve[0] = sugar, DefaultValve
	} else {
		x.Image[0], x.Valve[0] = p, v
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
		sugar := fmt.Sprintf("sugar#%d", nsugar+1)
		carry[1] = &Carry{sugar, g}
		x.Image[1], x.Valve[1] = sugar, "_"
	} else {
		x.Image[1], x.Valve[1] = p, v
	}
	//
	if !Space(t) { // require newline at end
		return nil,nil
	}
	src.Become(t)
	return
}

func seeJoin(src *Src) (m Meaning, p, v Name, ok bool) {
	if p, v, ok = seeJoinAddress(src); ok { // valve (or empty string)
		return
	}
	m, ok = seeJoinMeaning(src) // int, string, etc.
	return
}

func seeJoinMeaning(src *Src) (m Meaning, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			ok = false
		}
	}()
	t := src.Copy()
	m = SeeMeaning(t)
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
	p := SeeAddress(t).(Address)
	t.Match(string(ValveSelector))
	v := SeeAddress(t).(Address)
	src.Become(t)
	return p.Simple(), v.Simple(), true
}
