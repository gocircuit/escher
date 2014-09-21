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
	return b.Materialize(nil, design, true)
}

type Being struct {
	mem Circuit
}

func NewBeing(m Circuit) *Being {
	return &Being{m}
}

func (b *Being) MaterializeAddress(addr Address) Reflex {
	x := b.mem.Lookup(addr.Path()...)
	if x == nil {
		log.Fatalf("Address %v is dangling", addr)
	}
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
	case nil:
		panic("report error")
	default:
		log.Fatalf("Not knowing how to materialize %v/%T", t, t)
	}
	panic(0)
}

func (b *Being) MaterializeCircuit(u Circuit) (super Reflex) {
	gates := make(map[Name]Reflex)
	for _, g := range u.Letters() {
		if g == Super {
			log.Fatalf("Circuit design overwrites the %s gate. In:\n%v\n", Super, u)
		}
		m := u.At(g)
		if a, ok := m.(Address); ok && len(a.Path()) == 1 && a.Path()[0] == "Fork" { // Generate circuit partition gates on the fly
			var arm []string
			var defaultUsed bool
			for vlv, _ := range u.Valves(g) {
				if vlv == "" { // 
					defaultUsed = true
				} else {
					arm = append(arm, vlv.(string))
				}
			}
			if !defaultUsed || len(arm) == 0 {
				log.Fatalf("Fork gate's default valve not linked or has no partition valves. In:\n%v\n", u)
			}
			gates[g] = MaterializeUnion(arm...)
		} else {
			gates[g] = b.Materialize(
				&Matter{Design: u},
				m, false,
			)
		}
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
