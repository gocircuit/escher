// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	"log"

	. "github.com/gocircuit/escher/circuit"
	. "github.com/gocircuit/escher/memory"
)

func Materialize(m *Memory, design Meaning) Reflex {
	b := &Being{m.StartHijack()}
	defer m.EndHijack()
	matter := &Matter{
		Design: design,
		Valve: nil,
		Path: []Name{},
		Super: nil,
	}
	return b.Materialize(matter, design, true)
}

type Being struct {
	mem Circuit
}

func NewBeing(m Circuit) *Being {
	return &Being{m}
}

func (b *Being) MaterializeAddress(addr Address) Reflex {
	matter := &Matter{
		Valve: nil,
		Path: []Name{},
		Super: nil,
	}
	return b.materializeAddress(matter, addr)
}

func (b *Being) materializeAddress(matter *Matter, addr Address) Reflex {
	val := b.mem.Lookup(addr.Path()...)
	if val == nil {
		log.Fatalf("Address %v is dangling", addr)
	}
	matter.Address = addr
	matter.Design = val
	return b.Materialize(matter, val, true)
}

func (b *Being) Materialize(matter *Matter, x Meaning, recurse bool) Reflex {
	switch t := x.(type) {
	// Addresses are materialized recursively
	case Address:
		return b.materializeAddress(matter, t)
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
			return b.MaterializeCircuit(matter, t)
		}
		return NewNounReflex(t)
	case nil:
		panic("report error")
	default:
		log.Fatalf("Not knowing how to materialize %v/%T", t, t)
	}
	panic(0)
}

func (b *Being) MaterializeCircuit(matter *Matter, u Circuit) (super Reflex) {
	gates := make(map[Name]Reflex)
	for _, g := range u.Letters() {
		if g == Super {
			log.Fatalf("Circuit design overwrites the %s gate. In:\n%v\n", Super, u)
		}
		m := u.At(g)
		gates[g] = b.Materialize(
			&Matter{
				Address: Address{},
				Design: m,
				Valve: valveSet(u.Valves(g)),
				Path: append(matter.Path, g),
				Super: matter,
			},
			m, false,
		)
	}
	super, gates[Super] = make(Reflex), make(Reflex)
	for v, _ := range u.Valves(Super) {
		super[v], gates[Super][v] = NewSynapse()
	}
	for _, g_ := range append(u.Letters(), DefaultValve) {
		g := g_
		for v_, t := range u.Valves(g) {
			v := v_
			tg, tv := t.Reduce()
			checkLink(u, gates, g, v, tg, tv)
			go Link(gates[g][v], gates[tg][tv])
			// go func() {
			// 	log.Printf("%s:%s -> %s:%s | %v %v", g, v, tg, tv, gates[g][v], gates[tg][tv])
			// 	Link(gates[g][v], gates[tg][tv])
			// }()
		}
	}
	return super
}

func checkLink(u Circuit, gates map[Name]Reflex, sg, sv, tg, tv Name) {
	// log.Printf(" %v:%v <=> %v:%v", sg, sv, tg, tv)
	if _, ok := gates[sg]; !ok {
		log.Fatalf("Unknown gate %v in circuit:\n%v\n", sg, u)
	}
	if _, ok := gates[tg]; !ok {
		log.Fatalf("Unknown gate %v in circuit:\n%v\n", tg, u)
	}
	if _, ok := gates[sg][sv]; !ok {
		log.Fatalf("Unknown valve %v:%v in circuit:\n%v\n", sg, sv, u)
	}
	if _, ok := gates[tg][tv]; !ok {
		log.Fatalf("Unknown valve %v:%v in circuit:\n%v\n", tg, tv, u)
	}
}

func valveSet(v map[Name]Vector) map[Name]struct{} {
	w := make(map[Name]struct{})
	for vlv, _ := range v {
		w[vlv] = struct{}{}
	}
	return w
}