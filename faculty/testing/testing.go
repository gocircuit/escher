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
	loaded bool
	value Value
}

func (m *Match) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	return nil
}

func (m *Match) OverCognize(eye *be.Eye, name Name, v interface{}) {
	if name == DefaultValve {
		return
	}
	m.Lock()
	defer m.Unlock()
	if !m.loaded {
		m.value, m.loaded = v, true
		return
	}
	if !Same(m.value, v) {
		log.Fatalf("mismatch between %v and %v", m.value, v)
	}
	eye.Show(DefaultValve, nil)
}

func (m *Match) Cognize(eye *be.Eye, v interface{}) {}
