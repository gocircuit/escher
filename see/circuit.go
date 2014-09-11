// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	. "github.com/gocircuit/escher/circuit"
)

func SeeCircuit(src *Src) (u Circuit) {
	defer func() {
		if r := recover(); r != nil {
			u = Nil
		}
	}()
	u = New()
	t := src.Copy()
	t.Form("{")
	Space(t)
	var i, j int
	for {
		q := t.Copy()
		Space(q)
		if pn, pm := SeePeer(q); pn != nil { // parse peer
			if _, ok := pn.(Nameless); ok { // if peer is nameless, this is a slice element
				u.Include(j, pm)
				j++
			} else {
				u.Include(pn, pm) // record the order of definition in the same namespace but with number keys
			}
		} else if x, carry := SeeReal(q, 2*i); x != nil { // parse matching
			i++
			for _, c := range carry { // add carry peers to circuit
				if c != nil {
					u.Include(c.Name, c.Meaning)
				}
			}
			u.Form(*x)
		} else {
			break
		}
		Space(q)
		t.Become(q)
	}
	Space(t)
	t.Form("}")
	src.Become(t)
	return
}
