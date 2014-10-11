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
	. "github.com/gocircuit/escher/circuit"
)

func init() {
	faculty.Register("io.Clunk", be.NewGateMaterializer(Clunk{}))
}

// Clunk…
type Clunk struct{}

func (Clunk) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	return Clunk{}
}

func (Clunk) Cognize(_ *be.Eye, v interface{}) {
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
