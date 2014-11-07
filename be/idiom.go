// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	. "github.com/gocircuit/escher/circuit"
)

// Idiom is a hierarchy of names with associated meanings.
// Alternatively, it is a key-value store wherein keys are sequences of names.
type Idiom Circuit

func NewIdiom() Idiom {
	return Idiom(New())
}

func (idiom Idiom) Recall(walk ...Name) Value {
	if len(walk) == 0 {
		return idiom
	}
	switch t := Circuit(idiom).At(walk[0]).(type) {
	case Idiom:
		return t.Recall(walk[1:]...)
	default:
		if len(walk) == 1 {
			return t
		}
	}
	return nil
}

func (idiom Idiom) Memorize(value Value, walk ...Name) {
	path, name := walk[:len(walk)-1], walk[len(walk)-1]
	u := idiom
	for _, step := range path {
		r := NewIdiom()
		if Circuit(u).Include(step, r) != nil {
			panic("overwriting idiom")
		}
		u = r
	}
	if Circuit(u).Include(name, value) != nil {
		panic("overwriting value")
	}
}

func (idiom Idiom) Merge(v Value) {
	switch t := v.(type) {
	case Circuit:
		Circuit(idiom).Merge(t)
	case Idiom:
		Circuit(idiom).Merge(Circuit(t))
	default:
		panic("uh")
	}
}
