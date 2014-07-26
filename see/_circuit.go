// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"github.com/gocircuit/escher/star"
)

func SeeCircuit(src *Src) (name string, x *star.Star) {
	name, x = SeePeer(src)
	if x == nil {
		return "", nil
	}
	// XXX: process?
}

func SeeCircuit(src *Src) (name string, x *star.Star) {
	cir = &Circuit{}
	t := src.Copy()
	Space(t)
	if named {
		cir.Name = Identifier(t) // empty-string identifier ok
		Space(t)
	}
	if !t.TryMatch("{") {
		return nil
	}
	Space(t)
	for {
		switch q := SeePeerOrMatching(t).(type) {
		case *Matching:
			cir.Match = append(cir.Match, q)
			continue
		case *Peer:
			cir.Peer = append(cir.Peer, q)
			continue
		}
		break
	}
	Space(t)
	t.Match("}")
	if !Space(t) { // require newline at end
		return nil
	}
	src.Become(t)
	return DesugarCircuit(cir)
}
