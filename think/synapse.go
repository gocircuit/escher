// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package think

import (
	// "fmt"
	"sync"

	. "github.com/gocircuit/escher/image"
)

// Cognize routines are called when a change in value is to be delivered to a reflex.
type Cognize func(value interface{})

// Synapse is the “wire” connecting two reflexes.
// It remembers the last value transmitted in order to stop propagation of same-value messages.
type Synapse struct {
	learn <-chan Cognize
	teach chan<- Cognize
	sync.Mutex
	q *ReCognizer
}

func NewSynapse() (x, y *Synapse) {
	xy, yx := make(chan Cognize, 1), make(chan Cognize, 1)
	x = &Synapse{
		learn: xy,
		teach: yx,
	}
	y = &Synapse{
		learn: yx,
		teach: xy,
	}
	return
}

func (m *Synapse) Focus(cognize Cognize) *ReCognizer {
	m.teach <- cognize
	q := <-m.learn
	m.Lock()
	defer m.Unlock()
	m.q = &ReCognizer{q: q}
	return m.q
}

// Merge attaches two endpoints, of distinct memories, together.
func Merge(m1, m2 *Synapse) {
	m2.teach <- <-m1.learn
	m1.teach <- <-m2.learn
}

// The two endpoints of a Synapse are ReCognizer objects.
type ReCognizer struct {
	q Cognize
	sync.Mutex
	memory interface{}
}

// ReCognize sends value to the reciprocal side of this synapse.
func (s *ReCognizer) ReCognize(value interface{}) {
	s.Lock()
	defer s.Unlock()
	r, okr := s.memory.(Image)
	v, okv := value.(Image)
	if okr && okv {
		if Same(r, v) {
			return
		}
		s.memory = v.Copy()
	} else {
		if s.memory == value {
			return
		}
		s.memory = value
	}
	s.q(value)
}

// PtrReCognizer
type PtrReCognizer struct {
	sync.Mutex
	q *ReCognizer
}

func (p *PtrReCognizer) Bind(q *ReCognizer) {
	p.Lock()
	defer p.Unlock()
	p.q = q
}

func (p *PtrReCognizer) ReCognize(v interface{}) {
	p.Lock()
	q := p.q
	p.Unlock()
	q.ReCognize(v)
}

// MapReCognizer
type MapReCognizer struct {
	sync.Mutex
	t map[string]*ReCognizer
}

func (p *MapReCognizer) Bind(name string, re *ReCognizer) {
	p.Lock()
	defer p.Unlock()
	if p.t == nil {
		p.t = make(map[string]*ReCognizer)
	}
	if _, present := p.t[name]; present {
		panic(1)
	}
	p.t[name] = re
}

func (p *MapReCognizer) ReCognize(name string, v interface{}) {
	p.Lock()
	q := p.t[name]
	p.Unlock()
	q.ReCognize(v)
}
