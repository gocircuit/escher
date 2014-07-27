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
	t := src.Copy()
	t.Match("{")
	Space(t)
	for i := 0; ; i++ {
		q := t.Copy()
		Space(q)
		fwd, rev, peer := SeePeerOrMatching(q, fmt.Sprintf("$%d", i))
		if peer == nil {
			break
		}
		Space(q)
		q.TryMatch(",")
		Space(q)
		t.Become(q)
		x.Merge(fwd, rev, peer)
	}
	Space(t)
	t.Match("}")
	src.Become(t)
	return x
}
