// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package io provides gates for manipulating Go's I/O types.
package io

import (
	"io"
	"io/ioutil"

	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	ns := faculty.Root.Refine("io")
	ns.AddTerminal("Clunk", Clunk{})
}

// Clunk…
type Clunk struct{}

func (Clunk) Materialize() think.Reflex {
	ioEndo, ioExo := think.NewSynapse()
	go func() {
		ioEndo.Focus(clunk)
	}()
	return think.Reflex{
		"IO": ioExo, 
	}
}

func clunk(v interface{}) {
	switch t := v.(type) {
	case io.Closer:
		t.Close()
	case io.Reader:
		io.Copy(ioutil.Discard, t)
	default:
		panic("io.clunk sees unrecognized type")
	}
}
