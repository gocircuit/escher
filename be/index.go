// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	// "fmt"
	// "log"

	. "github.com/gocircuit/escher/circuit"
)

// Index is a hierarchy of names with associated meanings.
// Alternatively, it is a key-value store wherein keys are sequences of names.
type Index Circuit

func NewIndex() Index {
	return Index(New().Grow("?", "Index"))
}

func IsIndex(v Value) bool {
	u, ok := v.(Circuit)
	if !ok {
		return false
	}
	s, ok := u.StringOptionAt("?")
	return ok && s == "Index"
}

func AsIndex(v Value) Index {
	if v.(Circuit).StringAt("?") != "Index" {
		panic("not an index")
	}
	return Index(v.(Circuit))
}

func (x Index) Recall(walk ...Name) Value {
	if len(walk) == 0 {
		return Circuit(x)
	}
	v := Circuit(x).At(walk[0])
	if u, ok := v.(Circuit); ok && IsIndex(u) {
		return AsIndex(u).Recall(walk[1:]...)
	}
	if len(walk) == 1 {
		return v
	}
	return nil
}

func (x Index) Memorize(value Value, walk ...Name) {
	cx, step := Circuit(x), walk[0]
	//
	y, ok := value.(Index)
	if ok {
		value = Circuit(y)
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
		cx.Include(step, Circuit(NewIndex()))
	}
	Index(cx.CircuitAt(step)).Memorize(value, walk[1:]...)
}

func (x Index) Merge(with Index) {
	Circuit(x).Merge(Circuit(with))
}

func (x Index) Print(prefix, indent string, recurse int) string {
	return Circuit(x).Print(prefix, indent, recurse)
}

func (x Index) String() string {
	return x.Print("", "\t", -1)
}
