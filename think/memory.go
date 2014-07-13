// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package think

import (
	"reflect"
)

// Cognize is a called when …
type Cognize func(value interface{})

// Memory is the “wire” connecting two reflexes.
// It remembers the last value transmitted in order to stop propagation of same-value messages.
type Memory struct {
	learn <-chan Cognize
	teach chan<- Cognize
	recognizer ReCognizer
}

func NewMemory() (x, y *Memory) {
	xy, yx := make(chan Cognize, 1), make(chan Cognize, 1)
	x = &Memory{learn: xy, teach: yx}
	y = &Memory{learn: yx, teach: xy}
	return
}

func (m *Memory) Attach(cognize Cognize) *ReCognizer {
	m.teach <- cognize
	m.recognizer.reciprocal = <-m.learn
	return &m.recognizer
}

// Merge attaches two endpoints, of distinct memories, together.
func Merge(m1, m2 *Memory) {
	m2.teach <- <-m1.learn
	m1.teach <- <-m2.learn
}

// The two endpoints of a Memory are ReCognizer objects.
type ReCognizer struct {
	reciprocal Cognize
	recognized interface{}
}

// ReCognize sends value to the reciprocal side of this memory wire.
func (s *ReCognizer) ReCognize(value interface{}) {
	if Same(s.recognized, value) {
		return
	}
	s.recognized = value
	s.reciprocal(value)
}

// Same returns true if its arguments are equal in value.
func Same(v, w interface{}) bool {
	return reflect.DeepEqual(v, w)
}
