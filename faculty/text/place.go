// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package text provides gates for manipulating text.
package text

import (
	"bytes"
	"text/template"
	"sync"

	"github.com/gocircuit/escher/think"
)

// Place …
type Place struct{}

func (Place) Materialize() think.Reflex {
	dataEndo, dataExo := think.NewSynapse()
	formEndo, formExo := think.NewSynapse()
	_Endo, _Exo := think.NewSynapse()
	go func() {
		h := &place{
			formed: make(chan struct{}),
		}
		h.ReCognizer = _Endo.Focus(think.DontCognize)
		formEndo.Focus(h.CognizeForm)
		dataEndo.Focus(h.CognizeData)
	}()
	return think.Reflex{
		"_": _Exo, 
		"Form": formExo, 
		"Data": dataExo, 
	}
}

type place struct {
	sync.Mutex
	t *template.Template
	formed chan struct{}
	*think.ReCognizer
}

func (h *place) CognizeForm(v interface{}) {
	h.Lock()
	defer h.Unlock()
	var err error
	h.t, err = template.New("").Parse(v.(string))
	if err != nil {
		panic(err)
	}
	close(h.formed)
}

func (h *place) CognizeData(v interface{}) {
	<-h.formed
	h.Lock()
	defer h.Unlock()
	var w bytes.Buffer
	if err := h.t.Execute(&w, v); err != nil {
		panic(err)
	}
	h.ReCognizer.ReCognize(w.String())
}
