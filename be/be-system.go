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

func MaterializeSystem(system interface{}, index, barrier Circuit) (residue interface{}) {
	defer func() {
		if r := recover(); r != nil {
			Panicf("system materialization glitch (%v), at barrier %v", r, PrintableMatter(barrier))
		}
	}()

	if barrier.IsNil() {
		barrier = New()
	}
	parent := New().
		Grow("Index", index).
		Grow("View", New()).
		Grow("System", system).
		Grow("Barrier", barrier)

	return route(system, nil, newSubMatter(parent))
}

// Required matter: Index, View
func route(design interface{}, given Reflex, matter Circuit) (residue interface{}) {
	switch t := design.(type) {
	case int, float64, complex128, string:
		return materializeNoun(given, matter.Grow("Noun", t))
	case Circuit:
		if IsVerb(t) {
			return materializeVerb(given, matter.Grow("Verb", t))
		} else {
			return materializeCircuit(given, matter.Grow("Circuit", t))
		}
	case Materializer:
		matter.Grow("Material", t)
		return t(given, matter)
	default:
		return materializeNoun(given, matter.Grow("Noun", t))
	}
	panic(fmt.Sprintf("unknown design type: %T", design))
}
