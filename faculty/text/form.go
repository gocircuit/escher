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
)

// ForkForm…
type ForkForm struct{}

func (ForkForm) Materialize() (be.Reflex, Value) {
	return be.MaterializeUnion("Form", "Data")
}

// FormBlend …
type FormBlend struct{}

func (FormBlend) Materialize() (be.Reflex, Value) {
	reflex, _ := be.NewEyeCognizer(
		func(eye *be.Eye, valve string, value interface{}) {
			if valve != "FormData" {
				return
			}
			fd := value.(Circuit)
			t, err := template.New("").Parse(fd.StringAt("Form"))
			if err != nil {
				panic(err)
			}
			var w bytes.Buffer
			if err = t.Execute(&w, gateHierarchy(fd.CircuitAt("Data"))); err != nil {
				panic(err)
			}
			eye.Show(DefaultValve, w.String())
		}, 
		"FormData", DefaultValve,
	)
	return reflex, FormBlend{}
}

func gateHierarchy(u Circuit) map[string]interface{} {
	r := make(map[string]interface{})
	for g, v := range u.Gates() {
		switch t := v.(type) {
		case Circuit:
			r[g.(string)] = gateHierarchy(t)
		default:
			r[g.(string)] = v
		}
	}
	return r
}
