// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"strconv"
	. "github.com/gocircuit/escher/image"
)

const MatchingName = "="

func SeeUnion(src *Src) (x Image) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	x = Make()
	m := Make()
	x.Grow(MatchingName, m)
	t := src.Copy()
	t.Match("{")
	Space(t)
	var i int
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
			x.Attach(peer)
		} else {
			m.Grow(strconv.Itoa(i), match)
			i++
		}
	}
	Space(t)
	t.Match("}")
	src.Become(t)
	if m.Len() == 0 { // no matchings
		x.Abandon(MatchingName)
	}
	return
}
