// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package i provides reflection primitives.
package i

import (
	"sync"

	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/be"
	eu "github.com/gocircuit/escher/understand"
	"github.com/gocircuit/escher/kit/plumb"
)

// Memory
type Memory struct{}

func (Memory) Materialize(*be.Matter) be.Reflex {
	h := &memory{
		focus:     make(chan []string),
		learn:     make(chan *eu.Circuit),
		recall:    make(chan []string),
	}
	reflex, eye := plumb.NewEyeCognizer(h.Cognize, "Focus", "Learn", "Recall", "Use")
	go h.loop(eye)
	return reflex
}

type memory struct {
	focus     chan []string
	learn     chan *eu.Circuit
	recall    chan []string
	use       *be.ReCognizer
}

func (h *memory) Cognize(_ *plumb.Eye, dvalve string, dvalue interface{}) {
	switch dvalve {
	case "Focus":
		h.CognizeFocus(dvalue)
	case "Learn":
		h.CognizeLearn(dvalue)
	case "Recall":
		h.CognizeRecall(dvalue)
	}
}

func (h *memory) CognizeLearn(v interface{}) {
	switch t := v.(type) {
	case *eu.Circuit:
		h.learn <- t
	case nil:
		h.learn <- nil
	}
}

func (h *memory) CognizeFocus(v interface{}) {
	switch t := v.(type) {
	case []string:
		h.focus <- t
	case Image:
		var x []string
		for _, step := range t.Numbers() {
			x = append(x, t.String(step))
		}
		h.focus <- x
	case nil:
		h.focus <- nil
	}
}

func (h *memory) CognizeRecall(v interface{}) {
	switch t := v.(type) {
	case []string:
		h.recall <- t
	case Image:
		var x []string
		for _, step := range t.Numbers() {
			x = append(x, t.String(step))
		}
		h.recall <- x
	case nil:
		h.recall <- nil
	}
}

func (h *memory) loop(eye *plumb.Eye) {
	var root = &space{
		Space: make(be.Space),
	}
	focus, recall := &attention{root: root}, &attention{root: root}
	var lesson *eu.Circuit
	for {
		select {
		case lesson = <-h.learn:
		case walk := <-h.focus:
			focus.Point(walk...)
			if lesson == nil {
				break
			}
			eye.Show("Use", focus.Remember(lesson))
		case walk := <-h.recall:
			recall.Point(walk...)
			eye.Show("Use", recall.Recall())
		}
	}
}

// Materializable captures a circuit space and a pointer to a specific circuit design,
// functionaly solely able to materialize copies of the referenced circuit design.
type Materializable struct {
	root *space
	walk []string
}

func (x *Materializable) Materialize(*be.Matter) be.Reflex {
	x.root.Lock()
	defer x.root.Unlock()
	return x.root.Materialize(x.walk...)
}

// space captures a reentrant version of a think space as needed by the machinery in the memory reflex.
type space struct {
	sync.Mutex
	be.Space
}

// Roam returns the subspace whose root path is walk as child.
func (x *space) Roam(walk ...string) (parent, child interface{}) {
	x.Lock()
	defer x.Unlock()
	return x.Space.Roam(walk...)
}

// Materialize materializes the circuit design at path walk.
func (x *space) Materialize(walk ...string) be.Reflex {
	x.Lock()
	defer x.Unlock()
	return x.Space.Materialize(walk...)
}

// attention is a “pointer” to a (parent, name, child) triplet in a root faculty name space.
type attention struct {
	sync.Mutex
	root          *space
	parent, child interface{}
	walk          []string // walk to child in root
}

func (a *attention) name() string {
	if len(a.walk) > 0 {
		return a.walk[len(a.walk)-1]
	}
	return ""
}

func (a *attention) Point(walk ...string) {
	a.Lock()
	defer a.Unlock()
	if a.parent, a.child = a.root.Roam(walk...); a.parent != nil { // if parent is non-nil, child is not root
		a.walk = walk
	} else { // otherwise, child is root
		a.walk = nil
	}
	switch a.child.(type) { // if child is non-nil, it is a subspace or a circuit design
	case nil, be.Space, *eu.Circuit:
		return
	}
	panic(1)
}

func (a *attention) Remember(lesson *eu.Circuit) *Materializable {
	a.Lock()
	defer a.Unlock()
	if a.child == nil {
		return nil // can't remember if not focused
	}
	a.root.Lock()
	defer a.root.Unlock()
	lesson = a.child.(be.Space).Interpret(lesson)
	return &Materializable{
		root: a.root,
		walk: a.walk,
	}
}

func (a *attention) Recall() *Materializable {
	a.Lock()
	defer a.Unlock()
	switch a.child.(type) {
	case *eu.Circuit:
	default:
		panic("recall path not a circuit design")
	}
	//
	a.root.Lock()
	defer a.root.Unlock()
	return &Materializable{
		root: a.root,
		walk: a.walk,
	}
}
