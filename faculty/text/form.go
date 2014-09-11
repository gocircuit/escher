// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package text provides gates for manipulating text.
package text

import (
	"bytes"
	"text/template"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/faculty/basic"
	"github.com/gocircuit/escher/kit/plumb"
)

// ForkForm…
type ForkForm struct{}

func (ForkForm) Materialize() be.Reflex {
	return basic.MaterializeUnion("_", "Form", "Data")
}

// FormBlend …
type FormBlend struct{}

func (FormBlend) Materialize() be.Reflex {
	reflex, _ := plumb.NewEyeCognizer(
		func(eye *plumb.Eye, valve string, value interface{}) {
			if valve != "FormData" {
				return
			}
			fd := value.(Circuit)
			t, err := template.New("").Parse(fd.StringAt("Form"))
			if err != nil {
				panic(err)
			}
			var w bytes.Buffer
			if err = t.Execute(&w, fd.At("Data")); err != nil {
				panic(err)
			}
			eye.Show("_", w.String())
		}, 
		"FormData", "_",
	)
	return reflex
}
