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

// Matchings is a name type for the image of matchings within a circuit design image.
type Matchings struct{}

func SeeUnion(src *Src) (x interface{}) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	y := Make()
	m := Make()
	y.Grow(Matchings{}, m)
	t := src.Copy()
	t.Match("{")
	Space(t)
	var i, j int
	for {
		q := t.Copy()
		Space(q)
		peer, match := SeePeerOrMatching(q)
		if peer == nil && match == nil {
			break
		}
		Space(q)
		t.Become(q)
		if peer != nil {
			keys := peer.Names()
			if _, noname := keys[0].(nameless); noname { // if peer is nameless, this is a slice element
				y.Grow(Number(j), peer[nameless{}])
				j++
			} else {
				y.Attach(peer)
			}
		} else {
			m.Grow(i, match)
			i++
		}
	}
	Space(t)
	t.Match("}")
	src.Become(t)
	if m.Len() == 0 { // no matchings
		y.Abandon(Matchings{})
	}
	return y
}
