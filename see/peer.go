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

func SeePeerOrMatching(src *Src) (name string, peer, match Image) {
	if name, peer = SeePeer(src); peer.Lit() {
		return
	}
	if match = SeeMatching(src); match.Lit() {
		return
	}
	return "", NoImage, NoImage
}

func SeePeer(src *Src) (name string, x Image) {
	defer func() {
		if r := recover(); r != nil {
			x = NoImage
		}
	}()
	t := src.Copy()
	Space(t)
	name = Identifier(t)
	if len(name) == 0 {
		panic("peer name")
	}
	Whitespace(t)
	if x = SeeArithmeticOrNameOrImage(t); x == nil { // composite
		return "", NoImage
	}
	if !Space(t) { // require newline at end
		return "", NoImage
	}
	src.Become(t)
	return
}
