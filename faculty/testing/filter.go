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
	// check for #End markers
	if x.Has("#End") {
		eye.Show("Out", New().Grow("#End", x.At("#End")))
		return
	}

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
	y := New().
		Grow("Address", x.AddressAt("Address")).
		Grow("Name", name).
		Grow("View", view)
	eye.Show("Out", y)
}

func (FilterAll) CognizeOut(eye *be.Eye, v interface{}) {}
