// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	// "fmt"
	. "github.com/gocircuit/escher/image"
)

func SeeMatching(src *Src) (x Image) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	x = Make()
	t := src.Copy()
	Space(t)
	j0 := SeeJoin(t)
	x.Grow(0, j0)
	Whitespace(t)
	t.Match("=")
	Whitespace(t)
	j1 := SeeJoin(t)
	x.Grow(1, j1)
	if !Space(t) { // require newline at end
		return nil
	}
	src.Become(t)
	return
}

func SeeJoin(src *Src) (x Image) {
	if x = seePeerValveJoin(src); x != nil { // valve (or empty string)
		return x
	}
	if x = seeDesignJoin(src); x != nil { // int, string, etc.
		return x
	}
	panic(1)
}

func seeDesignJoin(src *Src) (x Image) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	t := src.Copy()
	dsgn := SeeSymbol(t)
	switch dsgn.(type) {
	case nil, Name, Address:
		return nil
	}
	src.Become(t)
	return Image{
		"Design":  dsgn,
	}
}

func seePeerValveJoin(src *Src) (x Image) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	t := src.Copy()
	peer := SeeName(t)
	if peer == nil {
		return nil
	}
	t.Match(string(ValveSelector))
	valve := SeeName(t)
	if valve == nil {
		return nil
	}
	src.Become(t)
	return Image{
		"Peer":  string(peer.(Name)),
		"Valve": string(valve.(Name)),
	}
}
