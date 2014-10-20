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

func SeeCircuit(src *Src) (v Value) {
	defer func() {
		if r := recover(); r != nil {
			v = nil
		}
	}()
	t := src.Copy()
	Space(t)
	t.Match("{")
	if v = SeeChamber(t); v == nil {
		return nil
	}
	t.Match("}")
	Space(t)
	src.Become(t)
	return
}

func SeeChamber(src *Src) (v Value) {
	defer func() {
		if r := recover(); r != nil {
			v = nil
		}
	}()
	u := New()
	t := src.Copy()
	Space(t)
	var j int
	for {
		q := t.Copy()
		Space(q)
		if pn, pm := SeePeer(q); pn != nil { // parse peer
			if _, ok := pn.(Nameless); ok { // if peer is nameless, give it the next numeric index
				u.Include(j, pm)
				j++
			} else {
				u.Include(pn, pm) // record the order of definition in the same namespace but with number keys
			}
		} else if x, carry := SeeLink(q, j); x != nil { // parse matching
			for _, c := range carry { // add carry peers to circuit
				if c != nil {
					u.Include(c.Name, c.Value)
					j++ // one numeric gate index is used in each carry
				}
			}
			u.Link(x[0], x[1])
		} else {
			break
		}
		Space(q)
		t.Become(q)
	}
	Space(t)
	src.Become(t)
	return u
}
