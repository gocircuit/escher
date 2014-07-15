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
)

// Eye is an implementation of Leslie Valiant's “Mind's Eye”, described in
//	http://www.probablyapproximatelycorrect.com/
type Eye struct {
	synapse map[tree.Name]*think.Synapse
	attention EyeReCognizer
}

type EyeReCognizer struct {
	cognize ShortCognize
	recognize map[tree.Name]*think.ReCognizer
	sync.Mutex
	age int
	memory Memory
}

// NewEye creates a new short-term memory mechanism.
func NewEye(valve ...tree.Name) (think.Reflex, *Eye) {
	reflex := make(think.Reflex)
	m := &Eye{
		synapse: make(map[tree.Name]*think.Synapse),
		attention: EyeReCognizer{
			recognize: make(map[tree.Name]*think.ReCognizer),
			memory: make(Memory),
		},
	}
	for _, v := range valve {
		if _, ok := reflex[v]; ok {
			panic("two valves, same name")
		}
		reflex[v], m.synapse[v] = think.NewSynapse()
		m.attention.memory.Grow(v, 0, nil)
	}
	return reflex, m
}

func (m *Eye) Focus(cognize ShortCognize) *EyeReCognizer {
	// Locking prevents individual competing Focus invocations 
	// from initiating cogntion before all valves/synapses have been attached.
	m.attention.Lock()
	defer m.attention.Unlock()
	m.attention.cognize = cognize
	for v_, _ := range m.attention.memory {
		v := v_
		m.attention.recognize[v] = m.synapse[v].Focus(
			func(w interface{}) {
				m.attention.cognizeWith(v, w)
			},
		)
	}
	return &m.attention
}
