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
	left := SeeSymbol(t)
	if left == nil {
		panic("peer")
	}
	Whitespace(t)
	right := SeeSymbol(t)
	if !Space(t) { // require newline at end
		return nil
	}
	if right == nil { // one term (a value in a union)
		src.Become(t)
		return Image{nameless{}: left}
	} else { // two terms (name and value)
		src.Become(t)
		return Image{string(left.(Name)): right}
	}
	panic("peer")
}

type nameless struct{}
