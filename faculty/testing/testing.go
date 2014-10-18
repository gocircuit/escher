// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package testing

import (
	"log"
	"sync"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Register("testing.Match", be.NewNativeMaterializer(&Match{}))
}

type Match struct {
	sync.Mutex
	sign map[Name]int
	history []Circuit
}

func (m *Match) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	m.sign = make(map[Name]int)
	var i int
	for vlv, _ := range matter.View.Gate {
		if vlv == DefaultValve {
			continue
		}
		m.sign[vlv] = i
		m.history = append(m.history, New())
		i++
	}
	if len(m.sign) != 2 {
		panic("match gates need exactly two opposing non-default valves")
	}
	return nil
}

func (m *Match) OverCognize(eye *be.Eye, name Name, v interface{}) {
	if name == DefaultValve {
		return
	}
	m.Lock()
	defer m.Unlock()
	//
	i := m.sign[name]
	h := m.history[i]
	h.Grow(h.Len(), v)
	//
	g := m.history[1-i]
	if g.Len() < h.Len() {
		return
	}
	if !Same(g.At(h.Len()-1), v) {
		log.Fatalf("mismatch between %v and %v\n%v\n%v\n", g.At(h.Len()-1), v, h, g)
	}
	if v == Term { // EOF indicator
		eye.Show(DefaultValve, Term)
	}
	// eye.Show(DefaultValve, h.Len()) // send number of matches so far
}

func (m *Match) Cognize(eye *be.Eye, v interface{}) {}
