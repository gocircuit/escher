// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

import (
	// "fmt"
	"sync"

	"github.com/gocircuit/escher/think"
)

// Eye is an implementation of Leslie Valiant's “Mind's Eye”, described in
//	http://www.probablyapproximatelycorrect.com/
type Eye struct {
	retina map[string]*think.Synapse
	nerve  EyeNerve
}

// ShortCognize is the cognition interface provided by the Mind's Eye (short-term memory) mechanism.
// The short-term memory is what allows people to process a linguistic sentence with all its structure.
type ShortCognize func(Impression)

type EyeNerve struct {
	cognize   ShortCognize
	connected chan struct{}
	recognize think.MapReCognizer
	memory
}

type memory struct {
	sync.Mutex
	Age int
	Imp Impression
}

// NewEye creates a new short-term memory mechanism, called an eye.
func NewEye(valve ...string) (think.Reflex, *Eye) {
	reflex := make(think.Reflex)
	m := &Eye{
		retina: make(map[string]*think.Synapse),
		nerve: EyeNerve{
			connected: make(chan struct{}),
			memory: memory{
				Imp: MakeImpression(),
			},
		},
	}
	m.nerve.memory.Imp = MakeImpression()
	for _, v := range valve {
		if _, ok := reflex[v]; ok {
			panic("two valves, same name")
		}
		reflex[v], m.retina[v] = think.NewSynapse()
		m.nerve.memory.Imp.Show(0, v, nil)
	}
	return reflex, m
}

// Focus binds this short memory reflex to the response function cognize.
func (m *Eye) Focus(cognize ShortCognize) *EyeNerve {
	m.nerve.memory.Lock()         // Locking prevents individual competing Focus invocations
	defer m.nerve.memory.Unlock() // from initiating cognition before all valves/synapses have been attached.
	m.nerve.cognize = cognize
	ch := make(chan struct{})
	for v_, _ := range m.nerve.memory.Imp.Image {
		v := v_
		go func() {
			m.nerve.recognize.Bind(
				v,
				m.retina[v].Focus(
					func(w interface{}) {
						m.nerve.cognizeWith(v, w)
					},
				),
			)
			ch <- struct{}{}
		}()
	}
	for range m.nerve.memory.Imp.Image {
		<-ch
	}
	close(m.nerve.connected)
	return &m.nerve
}
