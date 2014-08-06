// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package think

import (
	"github.com/gocircuit/escher/star"
)

// Cognize routines are called when a change in value is to be delivered to a reflex.
type Cognize func(value see.Image)

// Synapse is the “wire” connecting two reflexes.
// It remembers the last value transmitted in order to stop propagation of same-value messages.
type Synapse struct {
	learn <-chan Cognize
	teach chan<- Cognize
	recognizer ReCognizer
}

func NewSynapse() (x, y *Synapse) {
	xy, yx := make(chan Cognize, 1), make(chan Cognize, 1)
	x = &Synapse{learn: xy, teach: yx}
	y = &Synapse{learn: yx, teach: xy}
	return
}

func (m *Synapse) Focus(cognize Cognize) *ReCognizer {
	m.teach <- cognize
	m.recognizer.reciprocal = <-m.learn
	return &m.recognizer
}

// Merge attaches two endpoints, of distinct memories, together.
func Merge(m1, m2 *Synapse) {
	m2.teach <- <-m1.learn
	m1.teach <- <-m2.learn
}

// The two endpoints of a Synapse are ReCognizer objects.
type ReCognizer struct {
	reciprocal Cognize
	recognized *star.Star
}

// ReCognize sends value to the reciprocal side of this synapse.
func (s *ReCognizer) ReCognize(value *star.Star) {
	if star.Same(s.recognized, value) {
		return
	}
	s.recognized = value.Copy()
	s.reciprocal(value)
}
