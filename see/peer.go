// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	// "fmt"
	"github.com/gocircuit/escher/star"
)

func SeePeerOrMatching(src *Src) (name string, peer, match *star.Star) {
	if name, peer = SeePeer(src); peer != nil {
		return
	}
	if match = SeeMatching(src); match != nil {
		return
	}
	return "", nil, nil
}

func SeePeer(src *Src) (name string, x *star.Star) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	t := src.Copy()
	Space(t)
	name = Identifier(t)
	if len(name) == 0 {
		panic("peer name")
	}
	SpaceNoNewline(t)
	if x = SeeArithmeticOrNameOrImage(t); x == nil {
		return "", nil
	}
	if !Space(t) { // require newline at end
		return "", nil
	}
	src.Become(t)
	return
}
