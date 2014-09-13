// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	"log"

	. "github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/circuit"
)

type Being struct {
	Faculty
}

func (b *Being) MaterializeAddress(addr Address) Reflex {
	// log.Printf("addressing %s", string(addr))
	_, u := b.Faculty.Lookup(addr.Strings()...)
	return b.Materialize(nil, u, true) // 
}

func (b *Being) Materialize(matter *Matter, x Meaning, recurse bool) Reflex {
	// log.Printf("materializing (%v) %v/%T", recurse, x, x)
	switch t := x.(type) {
	// Addresses are materialized recursively
	case Address:
		return b.MaterializeAddress(t)
	// Irreducible types are materialized as gates that emit the irreducible values
	case int, float64, complex128, string:
		return NewNounReflex(t)
	// Go-gates are materialized into runtime reflexes
	case MaterializerFunc:
		return t()
	case Materializer:
		return t.Materialize()
	case MaterializerWithMatter:
		return t.Materialize(matter)
	case Gate:
		return MaterializeInterface(t)
	case Circuit:
		if recurse {
			return b.MaterializeCircuit(t)
		}
		return NewNounReflex(t)
	case Super:
		log.Fatal("Cannot materialize super")
	case nil:
		log.Fatalf("Not found")
	default:
		log.Fatalf("Not knowing how to materialize %v/%T", t, t)
	}
	panic(0)
}

func (b *Being) MaterializeCircuit(u Circuit) (super Reflex) {
	images := make(map[Name]Reflex)
	var name Name
	for y, m := range u.Images() {
		if _, ok := y.(string); !ok {
			continue // don't materialize non-string images
		}
		if _, ok := m.(Super); ok {
			name = y
		} else {
			images[y] = b.Materialize(
				&Matter{Design: u},
				m, false,
			)
		}
	}
	if name == nil {
		panic("no super")
	}
	super, images[name] = make(Reflex), make(Reflex)
	for v, _ := range u.Valves(name) {
		super[v], images[name][v] = NewSynapse()
	}
	for _, vx := range u.Reals() {
		for _, x_ := range vx {
			x := x_
			// go func() {
			// 	log.Printf("\t%s:%s <==> %s:%s, %v :: %v", 
			// 		x.Image[0], x.Valve[0], x.Image[1], x.Valve[1],
			// 		images[x.Image[0]], images[x.Image[1]])
			// 	Link(images[x.Image[0]][x.Valve[0]], images[x.Image[1]][x.Valve[1]])
			// }()
			go Link(images[x.Image[0]][x.Valve[0]], images[x.Image[1]][x.Valve[1]])
		}
	}
	return super
}
