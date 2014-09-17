// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	"log"

	. "github.com/gocircuit/escher/circuit"
)

type Being struct {
	mem Circuit
}

func NewBeing(m Circuit) *Being {
	return &Being{m}
}

func (b *Being) MaterializeAddress(addr Address) Reflex {
	// log.Printf("addressing %s", string(addr))
	x := b.mem.Lookup(addr.Path()...)
	return b.Materialize(nil, x, true) // 
}

func (b *Being) Materialize(matter *Matter, x Meaning, recurse bool) Reflex {
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
	gates := make(map[Name]Reflex)
	var name Name // name of circuit u
	for _, g := range u.Letters() {
		m := u.At(g)
		if _, ok := m.(Super); ok {
			name = g
		} else {
			gates[g] = b.Materialize(
				&Matter{Design: u},
				m, false,
			)
		}
	}
	if name == nil {
		panic("no super")
	}
	super, gates[name] = make(Reflex), make(Reflex)
	for v, _ := range u.Valves(name) {
		super[v], gates[name][v] = NewSynapse()
	}
	for _, g_ := range u.Letters() {
		g := g_
		for v_, t := range u.Valves(g) {
			v := v_
			tg, tv := t.Reduce()
			go Link(gates[g][v], gates[tg][tv])
			// go func() {
			// 	log.Printf("%s:%s -> %s:%s | %v %v", g, v, tg, tv, gates[g][v], gates[tg][tv])
			// 	Link(gates[g][v], gates[tg][tv])
			// }()
		}
	}
	return super
}
