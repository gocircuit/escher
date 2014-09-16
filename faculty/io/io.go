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
	// "log"

	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/be"
)

func init() {
	ns := faculty.Root.Refine("io") // XXX: Root is not protected from races.
	ns.Grow("Clunk", Clunk{})
}

// Clunk…
type Clunk struct{}

func (Clunk) Materialize() be.Reflex {
	_Endo, _Exo := be.NewSynapse()
	go func() {
		_Endo.Focus(clunk)
	}()
	return be.Reflex{"_": _Exo}
}

func clunk(v interface{}) {
	go func() {
		switch t := v.(type) {
		case io.ReadCloser:
			io.Copy(ioutil.Discard, t)
			t.Close()
		case io.Reader:
			io.Copy(ioutil.Discard, t)
		case io.Closer:
			t.Close()
		default:
			panic("io.clunk sees unrecognized type")
		}
	}()
}
