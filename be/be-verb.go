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

// Required matter: Index, View, Verb
func materializeVerb(given Reflex, matter Circuit) (residue interface{}) {
	index, syntax := Index(matter.CircuitAt("Index")), matter.CircuitAt("Verb")
	verb, addr := Verb(syntax).Verb(), Verb(syntax).Address()

	val := index.Recall(addr...)

	switch verb {
	case "*":
		return route(val, given, newSubMatter(matter))
	case "@":
		return materializeNoun(given, newSubMatter(matter).Grow("Noun", val))
	}
	panic(fmt.Sprintf("unknown or missing verb: %v", String(syntax)))
}

func newSubMatter(matter Circuit) Circuit {
	return New().
		Grow("Index", matter.CircuitAt("Index")).
		Grow("View", matter.CircuitAt("View")).
		Grow("Super", matter)
}

func newSubMatterView(matter Circuit, view Circuit) Circuit {
	r := newSubMatter(matter)
	r.Include("View", view)
	return r
}
