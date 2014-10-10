// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	"log"

	. "github.com/gocircuit/escher/circuit"
	. "github.com/gocircuit/escher/kit/memory"
)

func Materialize(memory Circuit, design Value) (reflex Reflex, residual Value) {
	b := NewRenderer(memory)
	matter := &Matter{
		Design: design,
		View: Circuit{},
		Path: []Name{},
		Super: nil,
	}
	return b.Materialize(matter, design, true)
}

type Renderer struct {
	mem Memory
}

func NewRenderer(mem Circuit) *Renderer {
	return &Renderer{Memory(mem)}
}

func (b *Renderer) MaterializeAddress(addr Address) (Reflex, Value) {
	matter := &Matter{
		View: Circuit{},
		Path: []Name{},
		Super: nil,
	}
	return b.materializeAddress(matter, addr)
}

func (b *Renderer) materializeAddress(matter *Matter, addr Address) (Reflex, Value) {
	val := b.mem.Lookup(addr)
	if val == nil {
		panicf("Address %v is dangling", addr)
	}
	matter.Address = addr
	matter.Design = val
	return b.Materialize(matter, val, true)
}

func (b *Renderer) Materialize(matter *Matter, x Value, recurse bool) (Reflex, Value) {
	switch t := x.(type) {
	// Addresses are materialized recursively
	case Address:
		return b.materializeAddress(matter, t)
	// Irreducible types are materialized as gates that emit the irreducible values
	case int, float64, complex128, string:
		return MaterializeNoun(t)
	// Go-gates are materialized into runtime reflexes
	case func() (Reflex, Value):
		return t()
	case func(*Matter) (Reflex, Value):
		return t(matter)
	case MaterializerFunc:
		return t()
	case MaterializerWithMatterFunc:
		return t(matter)
	case Materializer:
		return t.Materialize()
	case MaterializerWithMatter:
		return t.Materialize(matter)
	case Circuit:
		if recurse {
			return b.MaterializeCircuit(matter, t)
		}
		return MaterializeNoun(t)
	case nil:
		panic("report error")
	default:
		log.Fatalf("Not knowing how to materialize %v/%T", t, t)
	}
	panic(0)
}

func (b *Renderer) MaterializeCircuit(matter *Matter, u Circuit) (Reflex, Value) {
	value := New()
	gates := make(map[Name]Reflex)
	for _, g := range u.Letters() {
		if g == Super {
			log.Fatalf("Circuit design overwrites the %s gate. In:\n%v\n", Super, u)
		}
		m := u.At(g)
		var gv Value
		gates[g], gv = b.Materialize(
			&Matter{
				Address: Address{},
				Design: m,
				View: u.View(g),
				Path: append(matter.Path, g),
				Super: matter,
			},
			m, false,
		)
		value.Gate[g] = gv
	}
	var super Reflex
	super, gates[Super] = make(Reflex), make(Reflex)
	for v, _ := range u.Valves(Super) {
		super[v], gates[Super][v] = NewSynapse()
	}
	// value.Gate["#Matter"] = matter.Circuit()
	for _, g_ := range append(u.Letters(), Super) {
		g := g_
		for v_, t := range u.Valves(g) {
			v := v_
			checkLink(u, gates, g, v, t.Gate, t.Valve)
			value.Link(Vector{g, v}, Vector{t.Gate, t.Valve})
			go Link(gates[g][v], gates[t.Gate][t.Valve])
			// go func() {
			// 	log.Printf("%s:%s -> %s:%s | %v %v", g, v, t.Gate, t.Valve, gates[g][v], gates[t.Gate][t.Valve])
			// 	Link(gates[g][v], gates[t.Gate][t.Valve])
			// }()
		}
	}
	return super, value
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
