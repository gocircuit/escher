// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"log"

	. "github.com/gocircuit/escher/union"
)

func SeeUnion(src *Src) (u Union) {
	defer func() {
		if r := recover(); r != nil {
			u = Nil
		}
	}()
	u = New()
	t := src.Copy()
	t.Match("{")
	Space(t)
	var i, j int
	for {
		q := t.Copy()
		Space(q)
		if pn, pm := SeePeer(q); pn != nil { // parse peer
			if _, ok := pn.(Nameless); ok { // if peer is nameless, this is a slice element
				u.Add(j, pm)
				j++
			} else {
				u.Add(pn, pm) // record the order of definition in the same namespace but with number keys
			}
		} else if x, carry := SeeMatching(q, 2*i); x != nil { // parse matching
			i++
			for _, c := range carry { // add carry peers to union
				if c != nil {
					u.Add(c.Name, c.Meaning)
				}
			}
			u.Match(*x)
		} else {
			break
		}
		Space(q)
		t.Become(q)
	}
	Space(t)
	t.Match("}")
	src.Become(t)
	return
}
