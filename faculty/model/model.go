// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package model provides a basis of gates for circuit traversal and transformation.
package model

import (
	"github.com/hoijui/escher/be"
	"github.com/hoijui/escher/faculty"
)

func init() {
	faculty.Register(be.NewMaterializer(IO{}), "model", "IO")
}
