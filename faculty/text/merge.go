// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package text provides gates for manipulating text.
package text

import (
	"bytes"
	"io"
	// "log"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/be"
)

func init() {
	faculty.Register("text.ForkMerge", ForkMerge{})
	faculty.Register("text.MergeBlend", MergeBlend{})
	faculty.Register("text.ForkForm", ForkForm{})
	faculty.Register("text.FormBlend", FormBlend{})
}

// ForkMerge…
type ForkMerge struct{}

func (ForkMerge) Materialize() (be.Reflex, Value) {
	return be.MaterializeUnion("X", "Y", "Z")
}

// MergeBlend …
type MergeBlend struct{}

func (MergeBlend) Materialize() (be.Reflex, Value) {
	reflex, _ := be.NewEyeCognizer(
		func(eye *be.Eye, valve string, value interface{}) {
			if valve != "XYZ" {
				return
			}
			xyz := value.(Circuit)
			var w bytes.Buffer
			w.WriteString(flatten(xyz.StringAt("X")))
			w.WriteString(flatten(xyz.StringAt("Y")))
			w.WriteString(flatten(xyz.StringAt("Z")))
			eye.Show(DefaultValve, w.String())
		}, 
		"XYZ", DefaultValve,
	)
	return reflex, MergeBlend{}
}

func flatten(v interface{}) string {
	switch t := v.(type) {
	case string:
		return t
	case []byte:
		return string(t)
	case byte:
		return string(t)
	case rune:
		return string(t)
	case io.Reader:
		var w bytes.Buffer 
		io.Copy(&w, t)
		return w.String()
	case nil:
		return ""
	}
	panic("unsupported")
}
