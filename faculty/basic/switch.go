// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Register(be.NewMaterializer(&Switch{}), "Switch")
}

type Switch struct {
	view Circuit
}

func (s *Switch) Spark(_ *be.Eye, matter Circuit, _ ...interface{}) Value {
	s.view = matter.CircuitAt("View")
	return nil
}

func (s *Switch) Cognize(eye *be.Eye, value interface{}) {
	switch t := value.(type) {
	case Circuit:
		if IsVerb(t) {
			if s.view.Has("Verb") {
				eye.Show("Verb", value)
				return
			}
		}
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
	default:
		if s.view.Has("Other") {
			eye.Show("Other", value)
		}
	}
}

func (s *Switch) OverCognize(eye *be.Eye, name Name, value interface{}) {}
