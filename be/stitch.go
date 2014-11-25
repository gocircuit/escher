// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	"fmt"

	. "github.com/gocircuit/escher/circuit"
)

// TODO: Add view
// TODO: matter -> history

func StitchNoun(given Reflex, memory Circuit) (expected Reflex, residue interface{}) {
	noun := memory.At("Design")
	for _, syn_ := range given {
		syn := syn_
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
	panic(fmt.Sprintf("unknown or missing verb: %v", String(syntax)))
}

func StitchDesign(given Reflex, memory Circuit) (expected Reflex, residue interface{}) {
	memory = memory.Grow("Stitch", "Design")
	design := memory.At("Design")

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
	panic(fmt.Sprintf("unknown design type: %T", design))
}

var SpiritAddress = NewVerbAddress("*", "escher", "Spirit")

func StitchCircuit(given Reflex, memory Circuit) (Reflex, interface{}) {

	memory = memory.Grow("Stitch", "Circuit")
	design := memory.CircuitAt("Design")

	// make links
	gates := make(map[Name]Reflex)
	gates[Super] = make(Reflex)
	for name, view := range design.Flow {
		if gates[name] == nil {
			gates[name] = make(Reflex)
		}
		for vlv, vec := range view {
			if gates[name][vlv] != nil {
				continue
			}
			if gates[vec.Gate] == nil {
				gates[vec.Gate] = make(Reflex)
			}
			gates[name][vlv], gates[vec.Gate][vec.Valve] = NewSynapse()
		}
	}

	// materialize gates
	residue := New()
	spirit := make(map[Name]interface{}) // channel to pass circuit residue back to spirit gates inside the circuit
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
			gates[g], gresidue, spirit[g] = MaterializeInstance(gates[g], gmemory, &Future{})
		} else {
			var leftover Reflex
			leftover, gresidue = StitchDesign(gates[g], gmemory.Grow("Design", gsyntax))
			if len(leftover) > 0 {
				panic(2)
			}
		}
		residue.Gate[g] = gresidue
	}

	// connect given synapses
	for vlv, s := range given {
		t, ok := gates[Super][vlv]
		if !ok {
			continue
		}
		delete(gates[Super], vlv)
		go Link(s, t)
	}

	// send residue of this circuit to all escher.Spirit reflexes
	res := CleanUp(residue)
	go func() {
		for _, f := range spirit {
			f.(*Future).Charge(res)
		}
	}()

	if len(gates[Super]) > 0 {
		panic("circuit valves left unconnected")
	}

	return gates[Super], res
}
