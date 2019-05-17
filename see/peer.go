// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"github.com/gocircuit/escher/a"
	cir "github.com/gocircuit/escher/circuit"
)

func SeePeer(src *a.Src) (n cir.Name, m cir.Value) {
	if n, m = seeNameGate(src); n != nil {
		return n, m
	}
	return seeNamelessGate(src)
}

func seeNameGate(src *a.Src) (n cir.Name, m cir.Value) {
	defer func() {
		if r := recover(); r != nil {
			n, m = nil, nil
		}
	}()
	t := src.Copy()
	a.Whitespace(t)
	left := SeeValue(t)
	if len(a.Whitespace(t)) == 0 {
		panic("no whitespace after name")
	}
	right := SeeValue(t)
	if !Space(t) { // require newline at end
		return nil, nil
	}
	if right == "" {
		panic("no gate value")
	}
	src.Become(t)
	return left, right
}

func seeNamelessGate(src *a.Src) (n cir.Name, m cir.Value) {
	defer func() {
		if r := recover(); r != nil {
			n, m = nil, nil
		}
	}()
	t := src.Copy()
	a.Whitespace(t)
	value := SeeValue(t)
	if !Space(t) { // require newline at end
		return nil, nil
	}
	if value == "" {
		panic("nameless empty-string value implicit")
	}
	src.Become(t)
	return Nameless{}, value
}

type Nameless struct{}
