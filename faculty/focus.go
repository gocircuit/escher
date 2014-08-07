// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

import (
	"sync"

	"github.com/gocircuit/escher/think"
)

// Eye is an implementation of Leslie Valiant's “Mind's Eye”, described in
//	http://www.probablyapproximatelycorrect.com/
type Eye struct {
	synapse map[string]*think.Synapse
	attention EyeReCognizer
}

// ShortCognize is the cognition interface provided by the Mind's Eye (short-term memory) mechanism.
// The short-term memory is what allows people to process a linguistic sentence with all its structure.
type ShortCognize func(Impression)

type EyeReCognizer struct {
	cognize ShortCognize
	recognize map[string]*think.ReCognizer
	sync.Mutex
	age int
	memory Impression
}

// NewEye creates a new short-term memory mechanism, called an eye.
func NewEye(valve ...string) (think.Reflex, *Eye) {
	reflex := make(think.Reflex)
	m := &Eye{
		synapse: make(map[string]*think.Synapse),
		attention: EyeReCognizer{
			recognize: make(map[string]*think.ReCognizer),
			memory: MakeImpression(),
		},
	}
	for _, v := range valve {
		if _, ok := reflex[v]; ok {
			panic("two valves, same name")
		}
		reflex[v], m.synapse[v] = think.NewSynapse()
	}
	return reflex, m
}

// Focus binds this short memory reflex to the response function cognize.
func (m *Eye) Focus(cognize ShortCognize) *EyeReCognizer {
	m.attention.Lock()  // Locking prevents individual competing Focus invocations 
	defer m.attention.Unlock()  // from initiating cognition before all valves/synapses have been attached.
	m.attention.cognize = cognize
	for v_, _ := range m.attention.memory.Image {
		v := v_
		m.attention.recognize[v] = m.synapse[v].Focus(
			func(w interface{}) {
				m.attention.cognizeWith(v, w)
			},
		)
	}
	return &m.attention
}
