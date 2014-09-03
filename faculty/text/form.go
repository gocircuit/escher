// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package text provides gates for manipulating text.
package text

import (
	"bytes"
	"sync"
	"text/template"

	"github.com/gocircuit/escher/be"
)

// Form …
type Form struct{}

func (Form) Materialize(*be.Matter) be.Reflex {
	dataEndo, dataExo := be.NewSynapse()
	formEndo, formExo := be.NewSynapse()
	_Endo, _Exo := be.NewSynapse()
	go func() {
		h := &form{
			formed: make(chan struct{}),
		}
		h.ReCognizer = _Endo.Focus(be.DontCognize)
		formEndo.Focus(h.CognizeForm)
		dataEndo.Focus(h.CognizeData)
	}()
	return be.Reflex{
		"_":    _Exo,
		"Form": formExo,
		"Data": dataExo,
	}
}

type form struct {
	sync.Mutex
	t      *template.Template
	formed chan struct{}
	*be.ReCognizer
}

func (h *form) CognizeForm(v interface{}) {
	h.Lock()
	defer h.Unlock()
	var err error
	h.t, err = template.New("").Parse(v.(string))
	if err != nil {
		panic(err)
	}
	close(h.formed)
}

func (h *form) CognizeData(v interface{}) {
	<-h.formed
	h.Lock()
	defer h.Unlock()
	var w bytes.Buffer
	if err := h.t.Execute(&w, v); err != nil {
		panic(err)
	}
	h.ReCognizer.ReCognize(w.String())
}
