// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	// "fmt"

	. "github.com/gocircuit/escher/circuit"
)

// Required matter: Index, View, Verb
func materializeVerb(given Reflex, matter Circuit) (residue interface{}) {
	defer func() {
		if r := recover(); r != nil {
			Panicf("verb materialization glitch (%v), at matter %v", r, PrintableMatter(matter))
		}
	}()

	index, syntax := Index(matter.CircuitAt("Index")), matter.CircuitAt("Verb")
	verb, addr := Verb(syntax).Verb(), Verb(syntax).Address()

	rel := relativize(matter)
	var val interface{}
	if len(rel) > 0 {
		val = index.Recall(append(rel, addr...)...) // lookup relative to enclosing circuit's parent circuit
	}
	if val == nil {
		val = index.Recall(addr...) // otherwise lookup globally
	}
	if val == nil {
		Panicf("dangling address %v, at %v", Verb(syntax), PrintableMatter(matter))
	}

	switch verb {
	case "*":
		return route(val, given, newSubMatter(matter))
	case "@":
		return materializeNoun(given, newSubMatter(matter).Grow("Noun", val))
	}
	Panicf("unknown or missing verb %v, at %v", String(syntax), PrintableMatter(matter))
	panic(2)
}

func newSubMatter(matter Circuit) Circuit {
	return New().
		Grow("Index", matter.CircuitAt("Index")).
		Grow("View", matter.CircuitAt("View")).
		Grow("Super", matter)
}

func relativize(matter Circuit) []Name {
	sup, ok := matter.CircuitOptionAt("Super")
	if !ok {
		return nil
	}
	if !sup.Has("Circuit") {
		return nil
	}
	supsup, ok := sup.CircuitOptionAt("Super")
	if !ok {
		return nil
	}
	supverb, ok := supsup.CircuitOptionAt("Verb")
	if !ok {
		return nil
	}
	reladdr := Verb(supverb).Address()
	if len(reladdr) < 2 {
		return nil
	}
	return reladdr[:len(reladdr)-1] // chop off the circuit name at the end
}
