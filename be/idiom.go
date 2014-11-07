// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	// "log"

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
	if len(walk) == 1 {
		if Circuit(idiom).Include(walk[0], value) != nil {
			panic("overwriting value")
		}
		return
	}
	if !Circuit(idiom).Has(walk[0]) {
		Circuit(idiom).Include(walk[0], NewIdiom())
	}
	Circuit(idiom).At(walk[0]).(Idiom).Memorize(value, walk[1:]...)
}

func (idiom Idiom) Merge(with Idiom) {
	u := Circuit(idiom)
	for n, v := range with.Gate {
		switch t := v.(type) {
		case Idiom:
			if !u.Has(n) {
				u.Include(n, v)
				break
			}
			u.At(n).(Idiom).Merge(t)
		default:
			if u.Include(n, v) != nil {
				panic("overwriting circuit value")
			}
		}
	}
}

func (idiom Idiom) Print(prefix, indent string, recurse int) string {
	return Circuit(idiom).Print(prefix, indent, recurse)
}

func (idiom Idiom) String() string {
	return idiom.Print("", "\t", -1)
}
