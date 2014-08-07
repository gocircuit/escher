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

func SeePeerOrMatching(src *Src) (peer, match Image) {
	if peer = SeePeer(src); peer != nil {
		return
	}
	if match = SeeMatching(src); match != nil {
		return
	}
	return
}

func SeePeer(src *Src) (x Image) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	t := src.Copy()
	Space(t)
	name := Identifier(t)
	if len(name) == 0 {
		panic("peer name")
	}
	Whitespace(t)
	p := SeeArithmeticOrNameOrUnion(t)
	if p == nil {
		return nil
	}
	if !Space(t) { // require newline at end
		return nil
	}
	src.Become(t)
	return Image{name: p}
}
