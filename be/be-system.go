// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	"fmt"
	"os"

	cir "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/kit/runtime"
)

func MaterializeSystem(system interface{}, index, barrier cir.Circuit) (residue interface{}) {
	defer func() {
		if r := recover(); r != nil {
			switch t := r.(type) {
			case Panic:
				fmt.Fprintf(os.Stderr, t.Msg)
				fmt.Fprintf(os.Stderr, PrintableMatter(t.Matter))
				runtime.PrintStack()
				os.Exit(1)
			default:
				panic(r)
			}
		}
	}()
	if barrier.IsNil() {
		barrier = cir.New()
	}
	parent := cir.New().
		Grow("Index", index).
		Grow("View", cir.New()).
		Grow("System", system).
		Grow("Barrier", barrier)

	return route(system, nil, newSubMatter(parent))
}

// Required matter: Index, View
func route(design interface{}, given Reflex, matter cir.Circuit) (residue interface{}) {
	switch t := design.(type) {
	case int, float64, complex128, string:
		return materializeNoun(given, matter.Grow("Noun", t))
	case cir.Circuit:
		if cir.IsVerb(t) {
			return materializeVerb(given, matter.Grow("Verb", t))
		} else {
			return materializeCircuit(given, matter.Grow("Circuit", t))
		}
	case Materializer:
		defer func() {
			if r := recover(); r != nil {
				panicWithMatter(matter, "materialization glitch (%v)", r)
			}
		}()
		matter.Grow("Material", t)
		return t(given, matter)
	default:
		return materializeNoun(given, matter.Grow("Noun", t))
	}
	panicWithMatter(matter, "unknown design type: %T", design)
	return
}
