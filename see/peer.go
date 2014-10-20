// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	// "log"

	. "github.com/gocircuit/escher/circuit"
)

func Parse(src *Src) (Name, Value) {
	return SeePeer(src)
}

func SeePeer(src *Src) (n Name, m Value) {
	if n, m = seeNameGate(src); n != nil {
		return n, m
	}
	return seeNamelessGate(src)
}

func seeNameGate(src *Src) (n Name, m Value) {
	defer func() {
		if r := recover(); r != nil {
			n, m = nil, nil
		}
	}()
	t := src.Copy()
	Space(t)
	left := SeeName(t)
	if left == "" {
		panic("no gate name")
	}
	Whitespace(t)
	right := SeeValue(t)
	if !Space(t) { // require newline at end
		return nil, nil
	}
	if right == nil {
		panic("no gate value")
	}
	src.Become(t)
	return left, right
}

func seeNamelessGate(src *Src) (n Name, m Value) {
	defer func() {
		if r := recover(); r != nil {
			n, m = nil, nil
		}
	}()
	t := src.Copy()
	Space(t)
	right := SeeValue(t)
	if !Space(t) { // require newline at end
		return nil, nil
	}
	if right == nil {
		panic("no gate value")
	}
	src.Become(t)
	return Nameless{}, right
}

type Nameless struct{}
