// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"
	"sync"

	"github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/kit/plumb"
)

func init() {
	ns := faculty.Root.Refine("db").Refine("leveldb")
	ns.AddTerminal("File", File{})
}

// File
type File struct{}

func (File) Materialize() think.Reflex {
	reflex, eye := plumb.NewEye("File", "StoreFocus", "StoreValue", "QueryFocus", "QueryResult")
	go func() { // dispatch
		??
	}
	go func() { // Store loop
		??
	}()
	go func() { // Retrieve loop
		??
	}()
	return reflex
}
