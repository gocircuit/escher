// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"fmt"
	"github.com/gocircuit/escher/star"
)

func SeeStar(src *Src) (x *star.Star) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	x = star.Make()
	m := star.Make()
	x.Merge("$", m)
	t := src.Copy()
	t.Match("{")
	Space(t)
	var i int
	for {
		q := t.Copy()
		Space(q)
		name, peer, match := SeePeerOrMatching(q)
		if peer == nil && match == nil {
			break
		}
		Space(q)
		t.Become(q)
		if peer != nil {
			x.Merge(name, peer)
		} else {
			k := fmt.Sprintf("%d", i)
			m.Merge(k, match)
			i++
		}
	}
	Space(t)
	t.Match("}")
	src.Become(t)
	return x
}
