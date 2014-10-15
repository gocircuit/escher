// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package draw

import (
	// "log"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

// Split…
type Split struct {
	view Circuit
}

func (s *Split) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	s.view = matter.View
	return nil
}

func (s *Split) Cognize(eye *be.Eye, val interface{}) {
	for arm, _ := range s.view.Gate {
		if arm == DefaultValve {
			continue
		}
		go eye.Show(arm, val.(Circuit).Copy())
	}
}

func (s *Split) OverCognize(*be.Eye, Name, interface{}) {}
