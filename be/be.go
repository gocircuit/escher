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

// XXX: (materializeVerb) first, lookup design within the index that encloses the address of this verb

func materializeNoun(given Reflex, matter Circuit) (residue interface{}) {
	noun := matter.At("Noun")
	for _, syn_ := range given {
		syn := syn_
		go syn.Connect(DontCognize).ReCognize(noun)
	}
	if len(given) > 0 {
		return nil
	}
	return noun
}

func materializeVerb(given Reflex, matter Circuit) (residue interface{}) {
	// Place backtrace info in matter frame
	matter = matter.Grow("Materialize", "Verb")

	// Read arguments
	index := Index(matter.CircuitAt("Index"))
	syntax := matter.CircuitAt("Verb")
	verb, addr := Verb(syntax).Verb(), Verb(syntax).Address()

	val := index.Recall(addr...)

	// Perform transform
	tmatter := New().
		Grow("Index", matter.CircuitAt("Index")).
		Grow("View", matter.CircuitAt("View")).
		Grow("Super", matter)

	switch verb {
	case "*":
		return MaterializeDesign(given, tmatter.Grow("Design", val))
	case "@":
		return materializeNoun(given, tmatter.Grow("Noun", val))
	}
	panic(fmt.Sprintf("unknown or missing verb: %v", String(syntax)))
}

func MaterializeDesign(given Reflex, matter Circuit) (residue interface{}) {
	matter = matter.Grow("Materialize", "Design")
	design := matter.At("Design")

	// MaterializeDesign can be called from Go, and thus there might be no contextual view.
	// In this case, we create an empty view for the downstream materializer.
	view, ok := matter.CircuitOptionAt("View")
	if !ok {
		view = New()
	}
	tmatter := New().
		Grow("Index", matter.CircuitAt("Index")).
		Grow("View", view).
		Grow("Super", matter)

	switch t := design.(type) {
	case int, float64, complex128, string:
		return materializeNoun(given, tmatter.Grow("Design", design))
	case Circuit:
		if IsVerb(t) {
			return materializeVerb(given, tmatter.Grow("Verb", t))
		} else {
			return materializeCircuit(given, tmatter.Grow("Circuit", t))
		}
	case Materializer:
		return t(given, tmatter)
	}
	panic(fmt.Sprintf("unknown design type: %T", design))
}

var SpiritAddress = NewVerbAddress("*", "escher", "Spirit")

func materializeCircuit(given Reflex, matter Circuit) interface{} {

	matter = matter.Grow("Materialize", "Circuit")
	design := matter.CircuitAt("Circuit")

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
			Panicf("Circuit design overwrites the “%s” gate. In design %v\n", Super, design)
		}
		gsyntax := design.At(g)
		var gresidue interface{}

		// Compute view of gate within circuit
		view := New()
		for vlv, vec := range design.Flow[g] {
			view.Grow(vlv, design.Gate[vec.Gate])
		}

		gmatter := New().
			Grow("Index", matter.CircuitAt("Index")).
			Grow("View", view).
			Grow("Super", matter)

		if Same(gsyntax, SpiritAddress) {
			gresidue, spirit[g] = MaterializeInstance(gates[g], gmatter, &Future{})
		} else {
			gresidue = MaterializeDesign(gates[g], gmatter.Grow("Design", gsyntax))
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

	return res
}
