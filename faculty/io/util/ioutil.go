// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package util

import (
	"bytes"
	"io"
	"io/ioutil"
	// "log"

	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	ns := faculty.Root.Refine("io").Refine("util")
	ns.AddTerminal("WriteFile", WriteFile{})
}

// WriteFile …
type WriteFile struct{}

func (WriteFile) Materialize() think.Reflex {
	nameEndo, nameExo := think.NewSynapse()
	contentEndo, contentExo := think.NewSynapse()
	go func() {
		h := writeFile{
			named: make(chan struct{}),
		}
		nameEndo.Focus(h.CognizeName)
		contentEndo.Focus(h.CognizeContent)
	}()
	return think.Reflex{
		"FileName": nameExo, 
		"Content": contentExo, 
	}
}

type writeFile struct {
	name string
	named chan struct{}
}

func (h *writeFile) CognizeName(v interface{}) {
	h.name = v.(string)
	close(h.named)
}

func (h *writeFile) CognizeContent(v interface{}) {
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
}
