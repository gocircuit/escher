// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package io

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/gocircuit/escher/be"
	cir "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Register(be.NewMaterializer(&WriteFile{}), "io", "WriteFile")
}

type WriteFile struct {
	name  string
	named chan struct{}
}

func (h *WriteFile) Spark(*be.Eye, cir.Circuit, ...interface{}) cir.Value {
	h.named = make(chan struct{})
	return &WriteFile{}
}

func (h *WriteFile) CognizeName(eye *be.Eye, v interface{}) {
	h.name = v.(string)
	close(h.named)
}

func (h *WriteFile) CognizeContent(eye *be.Eye, v interface{}) {
	<-h.named
	switch t := v.(type) {
	case string:
		ioutil.WriteFile(h.name, []byte(t), 0644)
	case []byte:
		ioutil.WriteFile(h.name, t, 0644)
	case io.Reader:
		var w bytes.Buffer
		io.Copy(&w, t)
		ioutil.WriteFile(h.name, w.Bytes(), 0644)
	default:
		panic("eh?")
	}
	eye.Show("Ready", 1)
}

func (h *WriteFile) CognizeReady(eye *be.Eye, v interface{}) {}
