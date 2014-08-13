// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	// "fmt"
	"strconv"
	. "github.com/gocircuit/escher/image"
)

const MatchingName = "="

func SeeUnion(src *Src) (x interface{}) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	y := Make()
	m := Make()
	y.Grow(MatchingName, m)
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
			keys := peer.Sort()
			if keys[0] == "" { // if peer is nameless, this is a slice element
				y.Grow(strconv.Itoa(j), peer[""])
				j++
			} else {
				y.Attach(peer)
			}
		} else {
			m.Grow(strconv.Itoa(i), match)
			i++
		}
	}
	Space(t)
	t.Match("}")
	src.Become(t)
	if m.Len() == 0 { // no matchings
		y.Abandon(MatchingName)
	}
	return y
}
