// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	cir "github.com/gocircuit/escher/circuit"
)

// Index is a hierarchy of names with associated meanings.
// Alternatively, it is a key-value store wherein keys are sequences of names.
type Index cir.Circuit

func NewIndex() Index {
	return Index(cir.New())
}

func IsIndex(v cir.Value) bool {
	_, ok := v.(cir.Circuit)
	return ok
}

func AsIndex(v cir.Value) Index {
	return Index(v.(cir.Circuit))
}

func (x Index) Recall(walk ...cir.Name) cir.Value {
	if len(walk) == 0 {
		return cir.Circuit(x)
	}
	v := cir.Circuit(x).At(walk[0])
	if u, ok := v.(cir.Circuit); ok && IsIndex(u) {
		return AsIndex(u).Recall(walk[1:]...)
	}
	if len(walk) == 1 {
		return v
	}
	return nil
}

func (x Index) Memorize(value cir.Value, walk ...cir.Name) {
	cx, step := cir.Circuit(x), walk[0]
	//
	y, ok := value.(Index)
	if ok {
		value = cir.Circuit(y)
	}
	//
	if len(walk) == 1 {
		// if w, ok := cx.OptionAt(step); ok {
		// 	panic(fmt.Sprintf("index memorize overwriting gate (%s->%v) with (%v)", step, w, value))
		// }
		cx.Include(step, value)
		return
	}
	if !cx.Has(step) { // next step is an index
		cx.Include(step, cir.Circuit(NewIndex()))
	}
	Index(cx.CircuitAt(step)).Memorize(value, walk[1:]...)
}

func (x Index) Merge(with Index) {
	cir.Circuit(x).Merge(cir.Circuit(with))
}
