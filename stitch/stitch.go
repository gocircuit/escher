// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package stitch

import (
	// "log"
	"sync"

	. "github.com/gocircuit/escher/circuit"
)

type Stitcher func(Reflex, Circuit) (Reflex, interface{})

func StitchNoun(given Reflex, memory Circuit) (expected Reflex, residue interface{}) {
	noun := memory.At("Noun")
	for _, syn_ := range given {
		go syn.Focus(DontCognize).ReCognize(noun)
	}
	if len(given) > 0 {
		return nil, nil
	}
	return nil, noun
}

func StitchVerb(given Reflex, memory Circuit) (expected Reflex, residue interface{}) {

	// I
	index := Index(memory.CircuitAt("Index"))
	syntax := memory.CircuitAt("Verb")
	verb, addr := Verb(syntax).Verb(), Verb(syntax).Address()

	// II
	// XXX: first, lookup design within the index that encloses the address of this verb
	val := index.Recall(addr...)

	// III
	memory = memory.Copy().Grow("Stitch", "Verb")
	switch verb {
	case "*":
		switch t := val.(type) {
		case int, float64, complex128, string:
		case Circuit:
			if IsVerb(t) {
				return StitchVerb(given, New().Grow("Index", index).Grow("Verb", t).Grow("Super", memory))
			} else {
				return StitchCircuit(given, New().Grow("Design", t).Grow("Super", memory))
			}
		case Stitcher:
			return t(given, New().Grow("Super", memory))
		}
	case "@":
		return StitchNoun(given, New().Grow("Noun", val).Grow("Super", memory))
	}
	panicf("unknown or missing verb: %v", String(syntax))
}

var SpiritAddress = NewVerbAddress("*", "escher", "Spirit")

func StitchCircuit(given Reflex, memory Circuit) (expected Reflex, residue interface{}) {

	design := memory.CircuitAt("Design")

	residue = New()
	gates := make(map[Name]Reflex)
	spirit := make(map[Name]interface{})

	// iterate and materialize gates
	for g, _ := range design.Gate {
		if g == Super {
			panicf("Circuit design overwrites the “%s” gate. In design %v\n", Super, design)
		}
		m := design.At(g)
		var gv Value
		if Same(m, SpiritAddress) {
			gates[g], gv, spirit[g] = MaterializeNativeInstance(
				&Matter{
					Index:  b.index,
					Verb:   nil,
					Design: m,
					View:   design.View(g),
					Path:   append(matter.Path, g),
					Super:  matter,
				},
				&Future{},
			)
		} else {
			gates[g], gv = b.Materialize(
				&Matter{
					Index:  b.index,
					Verb:   nil,
					Design: m,
					View:   design.View(g),
					Path:   append(matter.Path, g),
					Super:  matter,
				},
				m,
				false,
			)
		}
		residue.Gate[g] = gv
	}

	// compute the super reflex to be returned by this circuit's materialization
	var super Reflex
	super, gates[Super] = make(Reflex), make(Reflex)
	for v, _ := range design.Valves(Super) {
		super[v], gates[Super][v] = NewSynapse()
	}

	// residue.Gate[Genus] = matter.Circuit()

	// link up all gates
	for _, g_ := range append(design.Names(), Super) {
		g := g_
		for v_, t := range design.Valves(g) {
			v := v_
			checkLink(u, gates, g, v, t.Gate, t.Valve)
			residue.Link(Vector{g, v}, Vector{t.Gate, t.Valve})
			go Link(gates[g][v], gates[t.Gate][t.Valve])
			// go func() {
			//	log.Printf("%v:%v -> %v:%v | %v %v", g, v, t.Gate, t.Valve, gates[g][v], gates[t.Gate][t.Valve])
			// 	Link(gates[g][v], gates[t.Gate][t.Valve])
			// }()
		}
	}

	// send residue of this circuit to all escher.Spirit reflexes
	res := CleanUp(residue)
	go func() {
		for _, f := range spirit {
			f.(*Future).Charge(res)
		}
	}()

	return super, res

}
