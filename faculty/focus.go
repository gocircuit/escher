// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package faculty

import (
	"sync"

	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/tree"
	"github.com/gocircuit/escher/understand"
)

// Sentence is a collection of functional values, indexed by valve name.
// The tree-scheme of sentence is:
//	Sentence: Rank—>Functional
//				Valve—>string
//				Value—>interface{}
//				Age—>int
type Sentence tree.Tree

// ShortCognize is the cognition interface provided by the Mind's Eye (short-term memory) mechanism.
// The short-term memory is what allows people to process a linguistic sentence with all its structure.
type ShortCognize func(Sentence)

// Eye is an implementation of Leslie Valiant's “Mind's Eye”, described in
//	http://www.probablyapproximatelycorrect.com/
type Eye struct {
	y []*think.Synapse
	recognizer EyeReCognizer
}

// Memory is an internal representation of the 
//	Memory:	Valve—>
//				Valve—>string
//				Value—>interface{}
//				Age—>int
//				Index—>int
type Memory tree.Tree

type EyeReCognizer struct {
	cognize ShortCognize
	recognize map[string]*think.ReCognizer
	sync.Mutex
	age int
	memory Memory
}

// NewEye creates a new short-term memory mechanism.
func NewEye(valve ...string) (think.Reflex, *Eye) {
	reflex := make(think.Reflex)
	m := &Eye{
		y: make([]*think.Synapse, len(valve)),
		recognizer: EyeReCognizer{
			recognize: make(map[string]*think.ReCognizer),
			memory: make(Memory),
		},
	}
	for i, v := range valve {
		if _, ok := reflex[v]; ok {
			panic("duplicate valve")
		}
		reflex[v], m.y[i] = think.NewSynapse()
		m.recognizer.memory.Grow(v, tree.Plant("Valve", v).Grow("Value", nil).Grow("Age", 0).Grow("Index", i))
	}
	return reflex, m
}

func (m *Eye) Attach(cognize ShortCognize) *EyeReCognizer {
	// Locking prevents individual completing Attach invocations 
	// from initiating cogntion, before all valves have been attached.
	m.recognizer.Lock()
	defer m.recognizer.Unlock()
	//
	m.recognizer.cognize = cognize
	for v_, f_ := range m.recognizer.memory {
		v := v_
		f := f_.(tree.Tree)
		m.recognizer.recognize[v] = m.y[f.Int("Index")].Attach(
			func(w interface{}) {
				m.recognizer.cognizeOn(v, w)
			},
		)
	}
	return &m.recognizer
}
