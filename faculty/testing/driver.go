// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package testing

import (
	// "log"
	"strings"
	"unicode"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
)

// 
type FilterAll struct {}

func (FilterAll) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	return nil
}

func (FilterAll) CognizeIn(eye *be.Eye, v interface{}) {
	x := v.(Circuit)
	name, view := x.StringAt("Name"), x.CircuitAt("View")
	if !strings.HasPrefix(name, "Test") {
		return
	}
	name = name[len("Test"):]
	if len(name) == 0 || !unicode.IsUpper(rune(name[0])) {
		return
	}
	y := New().Grow("Name", name).Grow("View", view)
	if x.Has("#End") {
		y.Grow("#End", 1)
	}
	eye.Show("Out", y)
}

func (FilterAll) CognizeOut(eye *be.Eye, v interface{}) {}
