// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	// "fmt"

	. "github.com/gocircuit/escher/circuit"
)

func SeePeer(src *Src) (n Name, m Meaning) {
	defer func() {
		if r := recover(); r != nil {
			n, m = nil, nil
		}
	}()
	t := src.Copy()
	Space(t)
	left := SeeMeaning(t)
	if left == nil {
		panic("peer")
	}
	Whitespace(t)
	right := SeeMeaning(t)
	if !Space(t) { // require newline at end
		return nil, nil
	}
	if right == nil { // one term (a value)
		src.Become(t)
		return Nameless{}, left
	} else { // two terms (name and value)
		src.Become(t)
		if c, ok := right.(Circuit); ok {
			if _, ok := c.At(""); ok {
				panic("two anonymous super")
			}
			c.Change("", Super{})
		}
		return left.(Address).Simple(), right
	}
	panic("peer")
}

type Nameless struct{}
