// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	// "log"
	"fmt"

	. "github.com/gocircuit/escher/circuit"
)

func SeeCircuit(src *Src) (v Value) {
	defer func() {
		if r := recover(); r != nil {
			v = nil
		}
	}()
	t := src.Copy()
	t.Match("{")
	if v = SeeChamber(t); v == nil {
		return nil
	}
	Space(t)
	t.Match("}")
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
		println("x")
		q := t.Copy()
		Space(q)
		println("x=", q.String())
		if pn, pm := SeePeer(q); pn != nil { // parse peer
			println("x1", pn)
			if _, ok := pn.(Nameless); ok { // if peer is nameless, give it the next numeric index
				println("x1a")
				u.Include(j, pm)
				j++
			} else {
				fmt.Printf("x1b == %T/%v %T/%v\n", pn, pn, pm, pm) // name parses before address
				u.Include(pn, pm)                                  // record the order of definition in the same namespace but with number keys
			}
			println("x1/")
		} else if x, carry := SeeLink(q, j); x != nil { // parse matching
			println("x2")
			for _, c := range carry { // add carry peers to circuit
				if c != nil {
					u.Include(c.Name, c.Value)
					j++ // one numeric gate index is used in each carry
				}
			}
			u.Link(x[0], x[1])
		} else {
			println("w")
			break
		}
		Space(q)
		t.Become(q)
	}
	Space(t)
	src.Become(t)
	return u
}
