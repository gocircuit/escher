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
	eu "github.com/gocircuit/escher/understand"
	"github.com/gocircuit/escher/think"
)

// Memory
type Memory struct{}

func (Memory) Materialize() think.Reflex {
	focusEndo, focusExo := think.NewSynapse()
	learnEndo, learnExo := think.NewSynapse()
	recallEndo, recallExo := think.NewSynapse()
	useEndo, useExo := think.NewSynapse()
	go func() {
		h := &memory{
			connected: make(chan struct{}),
			focus: make(chan []string),
			learn: make(chan *eu.Circuit),
			recall: make(chan []string),
		}
		h.use = useEndo.Focus(think.DontCognize)
		close(h.connected)
		focusEndo.Focus(h.CognizeFocus)
		learnEndo.Focus(h.CognizeLearn)
		recallEndo.Focus(h.CognizeRecall)
		go h.loop()
	}()
	return think.Reflex{
		"Focus": focusExo, // write-only
		"Learn": learnExo, // write-only
		"Recall": recallExo, // write-only
		"Use": useExo, // read-only
	}
}

type memory struct {
	connected chan struct{}
	focus chan []string
	learn chan *eu.Circuit
	recall chan []string
	use *think.ReCognizer
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
		for _, step := range t.Sort() {
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
		for _, step := range t.Sort() {
			x = append(x, t.String(step))
		}
		h.recall <- x
	case nil:
		h.recall <- nil
	}
}

// attention is a “pointer” to a (parent, name, child) triplet in a root faculty name space.
type attention struct {
	sync.Mutex
	root *space
	parent, child interface{}
	walk []string // walk to child in root
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
	case nil, think.Space, *eu.Circuit:
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
	lesson = a.child.(think.Space).Interpret(lesson)
	return &Materializable{
		root: a.root,
		walk: a.walk,
	}
}

// Materializable captures a circuit space and a pointer to a specific circuit design,
// functionaly solely able to materialize copies of the referenced circuit design.
type Materializable struct {
	root *space
	walk []string
}

func (x *Materializable) Materialize() think.Reflex {
	x.root.Lock()
	defer x.root.Unlock()
	return x.root.Materialize(x.walk...)
}

func (h *memory) loop() {
	<-h.connected
	var root = &space{
		Space: make(think.Space),
	}
	focus, recall := &attention{root: root}, &attention{root: root}
	for {
		select {
		case lesson := <-h.learn:
			if lesson == nil {
				break // lessons cannot be deleted/forgotten, for debugging and self-study purposes
			}
			h.use.ReCognize(focus.Remember(lesson))
		case walk := <-h.focus:
			focus.Point(walk...)
		case walk := <-h.recall:
			recall.Point(walk...)
		}
	}
}

// space captures a reentrant version of a think space as needed by the machinery in the memory reflex.
type space struct {
	sync.Mutex
	think.Space
}

// Roam returns the subspace whose root path is walk as child.
func (x *space) Roam(walk ...string) (parent, child interface{}) {
	x.Lock()
	defer x.Unlock()
	return x.Space.Roam(walk...)
}

// Materialize materializes the circuit design at path walk.
func (x *space) Materialize(walk ...string) think.Reflex {
	x.Lock()
	defer x.Unlock()
	return x.Space.Materialize(walk...)
}
