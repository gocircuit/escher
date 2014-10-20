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
	"github.com/gocircuit/escher/see"
)

func Materialize(memory Circuit, design Value) (residual Value) {
	var reflex Reflex
	reflex, residual = materialize(memory, design)
	if len(reflex) > 0 {
		panic("circuit not closed")
	}
	return
}

func materialize(memory Circuit, design Value) (reflex Reflex, residual Value) {
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
	memory Circuit
}

func NewRenderer(memory Circuit) *Renderer {
	return &Renderer{memory}
}

func (b *Renderer) MaterializeAddress(addr Address) (Reflex, Value) {
	matter := &Matter{
		View: Circuit{},
		Path: []Name{},
		Super: nil,
	}
	return b.materializeAddress(matter, addr)
}

func filter(a Address) (addr Address, monkey bool) {
	if len(a.Path) == 0 {
		return a, false
	}
	f, ok := a.Path[0].(string)
	if !ok {
		return a, false
	}
	if len(f) == 0 {
		return a, false
	}
	if f[0] != '@' {
		return a, false
	}
	a.Path[0] = see.ParseName(f[1:])
	return a, true
}

func (b *Renderer) materializeAddress(matter *Matter, addr Address) (Reflex, Value) {
	// parse @-sign out from front of address
	addr, monkey := filter(addr)

	// looking up locally first: starting from enclosing circuit's parent (a directory circuit)
	var val Value
	if matter != nil && matter.Super != nil {
		enclosing := matter.Super.Address
		if len(enclosing.Path) > 0 {
			abs := Address{enclosing.Path[:len(enclosing.Path)-1]}
			abs = abs.Append(addr)
			val = Memory(b.memory).Lookup(abs)
			if val != nil {
				addr = abs
			}
		}
	}
	// lookup from root, if local lookup not resolved
	if val == nil {
		val = Memory(b.memory).Lookup(addr)
	}
	if val == nil {
		panicf("Address %v is dangling", addr)
	}
	matter.Address = addr
	matter.Design = val
	return b.Materialize(matter, val, !monkey)
}

func (b *Renderer) Materialize(matter *Matter, x Value, recurse bool) (Reflex, Value) {
	switch t := x.(type) {
	// Addresses are materialized recursively
	case Address:
		return b.materializeAddress(matter, t)
	// Irreducible types are materialized as gates that emit the irreducible values
	case int, float64, complex128, string:
		return MaterializeNoun(matter, t)
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
		return MaterializeNoun(matter, t)
	default:
		return MaterializeNoun(matter, t)
	}
	panic(0)
}

func (b *Renderer) MaterializeCircuit(matter *Matter, u Circuit) (Reflex, Value) {
	residual := New()
	gates := make(map[Name]Reflex)
	for g, _ := range u.Gate {
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
		residual.Gate[g] = gv
	}
	var super Reflex
	super, gates[Super] = make(Reflex), make(Reflex)
	for v, _ := range u.Valves(Super) {
		super[v], gates[Super][v] = NewSynapse()
	}
	// residual.Gate[Genus] = matter.Circuit()
	for _, g_ := range append(u.Names(), Super) {
		g := g_
		for v_, t := range u.Valves(g) {
			v := v_
			checkLink(u, gates, g, v, t.Gate, t.Valve)
			residual.Link(Vector{g, v}, Vector{t.Gate, t.Valve})
			go Link(gates[g][v], gates[t.Gate][t.Valve])
			// go func() {
			//	log.Printf("%v:%v -> %v:%v | %v %v", g, v, t.Gate, t.Valve, gates[g][v], gates[t.Gate][t.Valve])
			// 	Link(gates[g][v], gates[t.Gate][t.Valve])
			// }()
		}
	}
	return super, CleanUp(residual)
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
