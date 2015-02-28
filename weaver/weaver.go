package weaver

import (
	"sync"
)

// Syntactic objects
type Address []string
type Name string

// Weaver is a concurrent, write-once namespace of values.
type Weaver struct {
	sync.Mutex
	val Value
	sub map[string]*Weaver
}

// Semantic object
type Value interface{}

func NewWeaver() *Weaver {
	return &Weaver{
		sub: make(map[string]*Weaver),
	}
}

func (w *Weaver) Reflex() Reflex {
	w.Lock()
	defer w.Unlock()
	return w.val.(Reflex)
}

func (w *Weaver) Fix(val Value) bool {
	w.Lock()
	defer w.Unlock()
	if w.val != nil {
		return false
	}
	w.val = val
	return true
}

func (w *Weaver) Reach(p Address) *Weaver {
	if len(p) == 0 {
		return w
	}
	return w.refine(p[0]).Reach(p[1:])
}

func (w *Weaver) refine(s string) *Weaver {
	w.Lock()
	defer w.Unlock()
	u, ok := w.sub[s]
	if !ok {
		u = NewWeaver()
		w.sub[s] = u
	}
	return u
}
