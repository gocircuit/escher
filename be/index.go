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

// Index is a hierarchy of names with associated meanings.
// Alternatively, it is a key-value store wherein keys are sequences of names.
type Index Circuit

func NewIndex() Index {
	return Index(New())
}

func (x Index) Recall(walk ...Name) Value {
	if len(walk) == 0 {
		return x
	}
	switch t := Circuit(x).At(walk[0]).(type) {
	case Index:
		return t.Recall(walk[1:]...)
	default:
		if len(walk) == 1 {
			return t
		}
	}
	return nil
}

func (x Index) Memorize(value Value, walk ...Name) {
	if len(walk) == 1 {
		if Circuit(x).Include(walk[0], value) != nil {
			panic("overwriting value")
		}
		return
	}
	if !Circuit(x).Has(walk[0]) {
		Circuit(x).Include(walk[0], NewIndex())
	}
	Circuit(x).At(walk[0]).(Index).Memorize(value, walk[1:]...)
}

func (x Index) Merge(with Index) {
	u := Circuit(x)
	for n, v := range with.Gate {
		switch t := v.(type) {
		case Index:
			if !u.Has(n) {
				u.Include(n, v)
				break
			}
			u.At(n).(Index).Merge(t)
		default:
			if u.Include(n, v) != nil {
				panic("overwriting circuit value")
			}
		}
	}
}

func (x Index) Print(prefix, indent string, recurse int) string {
	return Circuit(x).Print(prefix, indent, recurse)
}

func (x Index) String() string {
	return x.Print("", "\t", -1)
}
