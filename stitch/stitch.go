// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	// "log"
	"sync"

	. "github.com/gocircuit/escher/circuit"
)

func StitchNoun(given Reflex, memory Circuit) (expected Reflex, residue interface{}) {
	noun := memory.At("Design")
	for _, syn_ := range given {
		go syn.Connect(DontCognize).ReCognize(noun)
	}
	if len(given) > 0 {
		return nil, nil
	}
	return nil, noun
}

func StitchVerb(given Reflex, memory Circuit) (expected Reflex, residue interface{}) {
	// Place backtrace info in memory frame
	memory = memory.Grow("Stitch", "Verb")

	// Read arguments
	index := Index(memory.CircuitAt("Index"))
	syntax := memory.CircuitAt("Design")
	verb, addr := Verb(syntax).Verb(), Verb(syntax).Address()

	// XXX: first, lookup design within the index that encloses the address of this verb
	val := index.Recall(addr...)

	// Perform transform
	tmemory := New().
		Grow("Index", memory.CircuitAt("Index")).
		Grow("Super", memory)

	switch verb {
	case "*":
		return StitchDesign(given, tmemory.Grow("Design", val))
	case "@":
		return StitchNoun(given, tmemory.Grow("Design", val))
	}
	panicf("unknown or missing verb: %v", String(syntax))
}

func StitchDesign(given Reflex, memory Circuit) (expected Reflex, residue interface{}) {
	memory = memory.Grow("Stitch", "Design")
	design := memory.CircuitAt("Design")

	tmemory := New().
		Grow("Index", memory.CircuitAt("Index")).
		Grow("Super", memory)

	switch t := design.(type) {
	case int, float64, complex128, string:
		return StitchNoun(given, tmemory.Grow("Design", design))
	case Circuit:
		if IsVerb(t) {
			return StitchVerb(given, tmemory.Grow("Design", t))
		} else {
			return StitchCircuit(given, tmemory.Grow("Design", t))
		}
	case Stitcher:
		return t(given, tmemory)
	}
	panicf("unknown design type: %T", design)
}

var SpiritAddress = NewVerbAddress("*", "escher", "Spirit")

func StitchCircuit(given Reflex, memory Circuit) (expected Reflex, residue interface{}) {

	memory = memory.Grow("Stitch", "Circuit")
	design := memory.CircuitAt("Design")

	residue = New()
	gates := make(map[Name]Reflex)
	spirit := make(map[Name]interface{}) // channel to pass circuit residue back to spirit gates inside the circuit

	// materialize gates
	for g, _ := range design.Gate {
		if g == Super {
			panicf("Circuit design overwrites the “%s” gate. In design %v\n", Super, design)
		}
		gsyntax := design.At(g)
		var gresidue interface{}

		gmemory := New().
			Grow("Index", memory.CircuitAt("Index")).
			Grow("Super", memory)

		if Same(gsyntax, SpiritAddress) {
			//??
			gates[g], gresidue, spirit[g] = MaterializeNativeInstance(__, &Future{})
		} else {
			??
			gates[g], gresidue = StitchDesign(nil, gmemory.Grow("Design", gsyntax))
		}
		residue.Gate[g] = gresidue
	}

	// create bridge synapses between outer and inner reflexes
	var boundary Reflex
	boundary, gates[Super] = make(Reflex), make(Reflex)
	for v, _ := range design.Valves(Super) {
		boundary[v], gates[Super][v] = NewSynapse()
	}

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

	// connect given valves, and return rest as expected
	for vlv, syn := range given {
		antisyn, ok := boundary[vlv]
		if !ok {
			panic("given valve not known")
		}
		delete(boundary, vlv)
		go Link(syn, antisyn)
	}

	// send residue of this circuit to all escher.Spirit reflexes
	res := CleanUp(residue)
	go func() {
		for _, f := range spirit {
			f.(*Future).Charge(res)
		}
	}()

	return boundary, res
}

func checkLink(u Circuit, gates map[Name]Reflex, sg, sv, tg, tv Name) {
	// log.Printf(" %v:%v <=> %v:%v", sg, sv, tg, tv)
	if _, ok := gates[sg]; !ok {
		panicf("In circuit: %v\nHas no gate %v\n", u, sg)
	}
	if _, ok := gates[tg]; !ok {
		panicf("In circuit: %v\nHas no gate %v\n", u, tg)
	}
	if _, ok := gates[sg][sv]; !ok {
		panicf("In circuit: %v\nGate %v has no valve :%v\n", u, sg, sv)
	}
	if _, ok := gates[tg][tv]; !ok {
		panicf("In circuit: %v\nGate %v has no valve :%v\n", u, tg, tv)
	}
}
