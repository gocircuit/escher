// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	cir "github.com/gocircuit/escher/circuit"
)

// *Spirit gates emit the residue of the enclosing circuit itself
var SpiritVerb = cir.NewVerbAddress("*", "Spirit")

// create all links before materializing gates
func createLinks(design cir.Circuit) map[cir.Name]Reflex {

	// create all links before materializing gates
	gates := make(map[cir.Name]Reflex)
	gates[cir.Super] = make(Reflex)
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

	return gates
}

func calcResidue(matter cir.Circuit, design cir.Circuit, gates map[cir.Name]Reflex) (cir.Circuit, map[cir.Name]interface{}) {

	residue := cir.New()
	spirit := make(map[cir.Name]interface{}) // channel to pass circuit residue back to spirit gates inside the circuit
	for g := range design.Gate {
		if g == cir.Super {
			panicWithMatter(matter, "Circuit design overwrites the empty-string gate, in design %v\n", design)
		}
		gSyntax := design.At(g)
		var gResidue interface{}

		// Compute view of gate within circuit
		view := cir.New()
		for vlv, vec := range design.Flow[g] {
			view.Grow(vlv, design.Gate[vec.Gate])
		}

		if cir.Same(gSyntax, SpiritVerb) {
			gResidue, spirit[g] = MaterializeInstance(gates[g], newSubMatterView(matter, view), &Future{})
		} else {
			if gCir, ok := gSyntax.(cir.Circuit); ok && !cir.IsVerb(gCir) {
				gResidue = materializeNoun(gates[g], newSubMatterView(matter, view).Grow("Noun", gCir))
			} else {
				gResidue = route(gSyntax, gates[g], newSubMatterView(matter, view))
			}
		}
		residue.Gate[g] = gResidue
	}

	return residue, spirit
}

// connect boundary synapses
func connect(given Reflex, matter cir.Circuit, design cir.Circuit, gates map[cir.Name]Reflex) {

	for vlv, s := range given {
		t, ok := gates[cir.Super][vlv]
		if !ok {
			panicWithMatter(matter, "connected valve %v is not connected within circuit design %v", vlv, design)
		}
		delete(gates[cir.Super], vlv)
		go Link(s, t)
		go Link(t, s)
	}
}

// send residue of this circuit to all escher.Spirit reflexes
func distributeResidue(residue cir.Circuit, spirit map[cir.Name]interface{}) cir.Circuit {

	res := CleanUp(residue)
	go func() {
		for _, f := range spirit {
			f.(*Future).Charge(res)
		}
	}()

	return res
}

// Required matter: Index, View, Circuit
func materializeCircuit(given Reflex, matter cir.Circuit) interface{} {

	design := matter.CircuitAt("Circuit")

	gates := createLinks(design) // materialize gates

	residue, spirit := calcResidue(matter, design, gates)

	connect(given, matter, design, gates)

	res := distributeResidue(residue, spirit)

	if len(gates[cir.Super]) > 0 {
		panicWithMatter(matter, "circuit valves left unconnected")
	}

	return res
}

func newSubMatterView(matter cir.Circuit, view cir.Circuit) cir.Circuit {
	r := newSubMatter(matter)
	r.Include("View", view)
	return r
}

// CleanUp removes nil-valued gates and their incident edges.
// CleanUp never returns nil.
func CleanUp(u cir.Circuit) cir.Circuit {
	for n, g := range u.Gate {
		if g != nil {
			continue
		}
		delete(u.Gate, n)
		for vlv, vec := range u.Flow[n] {
			u.Unlink(cir.Vector{Gate: n, Valve: vlv}, vec)
		}
	}
	return u
}
