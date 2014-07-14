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
type Sentence tree.Tree // Valve:string -> Functional:interface{}

??? sentence has to reflect arrival order


// …
// The first entry is the most recent one.
type ShortCognize func(Sentence)

type Eye struct {
	y []*think.Synapse
	recognizer EyeReCognizer
}

// 
func NewEye(valve ...string) (think.Reflex, *Eye) {
	reflex := make(think.Reflex)
	m := &Eye{
		y: make([]*think.Synapse, len(valve)),
		recognizer: EyeReCognizer{
			recognize: make(map[string]*think.ReCognizer),
			memory: make(Sentence),???
		},
	}
	for i, v := range valve {
		if _, ok := reflex[v]; ok {
			panic("duplicate valve")
		}
		reflex[v], m.y[i] = think.NewSynapse()
		m.recognizer.memory[i] = functional{Valve: v, Value: nil}
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
	for i, s := range m.recognizer.memory {
		s_ := s
		m.recognizer.recognize[s.Valve] = m.y[i].Attach(
			func(w interface{}) {
				m.recognizer.cognizeOn(s_.Valve, w)
			},
		)
	}
	return &m.recognizer
}
