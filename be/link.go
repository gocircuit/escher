// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	"log"
	"strings"

	. "github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/circuit"
)

type Being struct {
	Faculty
}

func (b *Being) MaterializeAddress(addr string) Reflex {
	_, u := b.Faculty.Lookup(strings.Split(addr, ".")...)
	return b.Materialize(u)
}

func (b *Being) Materialize(x Meaning) Reflex {
	switch t := x.(type) {
	case int, float64, complex128, string:
		return NewNounReflex(t)
	case Gate:
		return t.Materialize()
	case GateWithMatter:
		?
	case Circuit:
		return MaterializeCircuit(t)
	case nil:
		log.Fatalf("Not found")
	default:
		log.Fatalf("Not knowing how to materialize %v/%T", t, t)
	}
	panic(0)
}

func (b *Being) MaterializeCircuit(u Circuit) {
	??
}
