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

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Register("text.Merge", be.NewMaterializer(Merge{}))
	faculty.Register("text.Form", be.NewMaterializer(Form{}))
}

// Merge concatenates the string values of string-named gates into a single string output,
// where concatenation takes place in the lexicographic order of the gate names.
type Merge struct{}

func (Merge) Spark(*be.Eye, Circuit, ...interface{}) Value {
	return nil
}

func (Merge) CognizeIn(eye *be.Eye, v interface{}) {
	var w bytes.Buffer
	x := v.(Circuit)
	for _, name := range x.SortedLetters() {
		w.WriteString(flatten(x.StringAt(name)))
	}
	eye.Show("Out", w.String())
}

func (Merge) CognizeOut(eye *be.Eye, v interface{}) {}

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
