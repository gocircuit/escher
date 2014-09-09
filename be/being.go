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
	"github.com/gocircuit/escher/see"
)

type Being struct {
	Faculty
}

func (b *Being) MaterializeAddress(addr see.Address) Reflex {
	_, u := b.Faculty.Lookup(addr.Walk()...)
	return b.Materialize(u)
}

func (b *Being) Materialize(x Meaning) Reflex {
	switch t := x.(type) {
	case see.Address:
		return b.MaterializeAddress(t)
	case int, float64, complex128, string:
		return NewNounReflex(t)
	case Gate:
		return t.Materialize()
	// case GateWithMatter:
	// 	??
	case Circuit:
		return b.MaterializeCircuit(t)
	case nil:
		log.Fatalf("Not found")
	default:
		log.Fatalf("Not knowing how to materialize %v/%T", t, t)
	}
	panic(0)
}

func (b *Being) MaterializeCircuit(u Circuit) (super Reflex) {
	symbols := make(map[Name]Reflex)
	var name Name
	for y, m := range u.Images() {
		if _, ok := y.(string); !ok {
			continue // don't materialize non-string symbols
		}
		if _, ok := m.(Super); ok {
			name = y
		} else {
			symbols[y] = b.Materialize(m)
		}
	}
	if name == nil {
		panic("no super")
	}
	super, symbols[name] = make(Reflex), make(Reflex)
	for v, _ := range u.Valves(name) {
		super[v], symbols[name][v] = NewSynapse()
	}
	for _, vx := range u.Reals() {
		for _, x := range vx {
			go Link(symbols[x.Image[0]][x.Valve[0]], symbols[x.Image[1]][x.Valve[1]])
		}
	}
	return super
}
