// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"

	"github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
)

func init() {
	faculty.Register("Switch", be.NewNativeMaterializer(&Switch{}))
}

type Switch struct {
	view Circuit
}

func (s *Switch) Spark(_ *be.Eye, matter *be.Matter, _ ...interface{}) Value {
	s.view = matter.View
	return nil
}

func (s *Switch) Cognize(eye *be.Eye, value interface{}) {
	switch value.(type) {
	case Circuit:
		if s.view.Has("Circuit") {
			eye.Show("Circuit", value)
		}
	case int:
		if s.view.Has("Int") {
			eye.Show("Int", value)
		}
	case float64:
		if s.view.Has("Float") {
			eye.Show("Float", value)
		}
	case complex128:
		if s.view.Has("Complex") {
			eye.Show("Complex", value)
		}
	case string:
		if s.view.Has("String") {
			eye.Show("String", value)
		}
	case Address:
		if s.view.Has("Address") {
			eye.Show("Address", value)
		}
	default:
		if s.view.Has("Other") {
			eye.Show("Other", value)
		}
	}
}

func (s *Switch) OverCognize(eye *be.Eye, name Name, value interface{}) {}
