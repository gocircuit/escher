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

func SeePeerOrMatching(src *Src, name string) (fwd, rev string, x *star.Star) {
	if fwd, rev, x = SeePeer(src); x != nil {
		return
	}
	if fwd, rev, x = SeeMatching(src, name); x != nil {
		return
	}
	return "", "", nil
}

func SeePeer(src *Src) (fwd, rev string, x *star.Star) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	t := src.Copy()
	Space(t)
	fwd = Identifier(t)
	t.Match("\\")
	rev = Identifier(t)
	SpaceNoNewline(t)
	if x = SeeArithmeticOrNameOrStar(t); x == nil {
		return "", "", nil
	}
	if !Space(t) { // require newline at end
		return "", "", nil
	}
	src.Become(t)
	return
}
