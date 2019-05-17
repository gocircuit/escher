// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package test

import (
	// "log"
	"strings"
	"unicode"

	"github.com/gocircuit/escher/be"
	cir "github.com/gocircuit/escher/circuit"
)

//
type Filter struct{ be.Sparkless }

func (Filter) CognizeIn(eye *be.Eye, v interface{}) {
	x := v.(cir.Circuit)
	//
	name_, view := x.NameAt("Name"), x.CircuitAt("View")
	name, ok := name_.(string)
	if !ok {
		return
	}
	if !strings.HasPrefix(name, "Test") {
		return
	}
	sfx := name[len("Test"):]
	if len(sfx) == 0 || !unicode.IsUpper(rune(sfx[0])) {
		return
	}
	y := cir.New().
		Grow("Address", x.CircuitAt("Address")).
		Grow("Name", name).
		Grow("View", view)
	eye.Show("Out", y)
}

func (Filter) CognizeOut(eye *be.Eye, v interface{}) {}
