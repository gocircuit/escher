// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package test

import (
	"log"

	"github.com/gocircuit/escher/be"
	cir "github.com/gocircuit/escher/circuit"
)

// TODO: Make sure matching works when opposing streams come at different speeds
// Rename gates to Got and Expected.

//
type Match struct {
	name []cir.Name
	flow []chan interface{}
}

func (m *Match) Spark(eye *be.Eye, matter cir.Circuit, aux ...interface{}) cir.Value {
	for vlv := range matter.CircuitAt("View").Gate {
		if vlv == cir.DefaultValve {
			continue
		}
		m.name = append(m.name, vlv)
		m.flow = append(m.flow, make(chan interface{}, 1))
	}
	if len(m.name) != 2 {
		log.Fatalf("match gates need exactly two opposing non-default valves; have %v", m.name)
	}
	return nil
}

func (m *Match) OverCognize(eye *be.Eye, name cir.Name, v interface{}) {
	// compute valve index
	var i int
	for j, n := range m.name {
		if cir.Same(n, name) {
			i = j
			break
		}
	}
	// match
	select {
	case u := <-m.flow[1-i]: // if the opposing channel is ready
		if !cir.Same(u, v) {
			log.Fatalf("mismatch %v vs %v: %v vs %v\n", m.name[1-i], name, u, v)
		}
		eye.Show(cir.DefaultValve, v) // emit the matched object
	default: // otherwise, offer our value
		m.flow[i] <- v
	}
}

func (m *Match) Cognize(eye *be.Eye, v interface{}) {}
