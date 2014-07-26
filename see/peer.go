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

func SeePeerOrMatching(src *Src, anon string) (name string, x *star.Star) {
	if name, x = SeePeer(src); x != nil {
		return
	}
	if x = SeeMatching(src); x != nil {
		return anon, x
	}
	return "", nil
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
	if len(name) > 0 {
		SpaceNoNewline(t)
	}
	arithmeticOrNameOrStar := SeeArithmeticOrNameOrStar(t)
	if arithmeticOrNameOrStar == nil {
		return "", nil
	}
	if !Space(t) { // require newline at end
		return "", nil
	}
	src.Become(t)
	return name, arithmeticOrNameOrStar
}
