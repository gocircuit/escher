// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	// "log"

	. "github.com/gocircuit/escher/circuit"
)

func See(src *Src) (Name, Circuit) {
	n, v := SeePeer(src)
	if v == nil {
		return nil, Nil
	}
	return n, v.(Circuit)
}

func SeePeer(src *Src) (n Name, m Value) {
	defer func() {
		if r := recover(); r != nil {
			n, m = nil, nil
		}
	}()
	t := src.Copy()
	Space(t)
	left := SeeValue(t)
	if left == nil {
		panic("peer")
	}
	Whitespace(t)
	right := SeeValue(t)
	if !Space(t) { // require newline at end
		return nil, nil
	}
	if right == nil { // one term (a value)
		src.Become(t)
		return Nameless{}, left
	} else { // two terms (name and value)
		src.Become(t)
		return left.(Address).Simplify(), right
	}
	panic("peer")
}

type Nameless struct{}
