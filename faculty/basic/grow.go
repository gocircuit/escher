// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"

	"github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/be"
)

func init() {
	faculty.Root.AddTerminal("Grow", Grow{})
}

// Grow
type Grow struct{}

func (Grow) Materialize(*be.Matter) be.Reflex {
	imgEndo, imgExo := be.NewSynapse()
	keyEndo, keyExo := be.NewSynapse()
	valueEndo, valueExo := be.NewSynapse()
	_Endo, _Exo := be.NewSynapse()
	go func() {
		h := &grow{
			connected: make(chan struct{}),
			key:       make(chan interface{}),
			img:       make(chan interface{}),
			value:     make(chan interface{}),
		}
		h.z = _Endo.Focus(be.DontCognize)
		close(h.connected)
		keyEndo.Focus(h.CognizeKey)
		imgEndo.Focus(h.CognizeImg)
		valueEndo.Focus(h.CognizeValue)
		go h.loop()
	}()
	return be.Reflex{
		"Img":   imgExo,
		"Key":   keyExo,
		"Value": valueExo,
		"_":     _Exo,
	}
}

type grow struct {
	connected       chan struct{}
	key chan interface{}
	img chan interface{}
	value chan interface{}
	z               *be.ReCognizer
}

func (h *grow) CognizeKey(v interface{}) {
	h.key <- v
	close(h.key)
}

func (h *grow) CognizeImg(v interface{}) {
	h.img <- v
	close(h.img)
}

func (h *grow) CognizeValue(v interface{}) {
	h.value <- v
	close(h.value)
}

func (h *grow) loop() {
	<-h.connected
	ch := make(chan struct{})
	var key, img, value interface{}
	go func() {
		key = <-h.key
		ch <- struct{}{}
	}()
	go func() {
		img = <-h.img
		ch <- struct{}{}
	}()
	go func() {
		value = <-h.value
		ch <- struct{}{}
	}()
	<-ch
	<-ch
	<-ch
	switch key.(type) {
	case string, int:
		h.z.ReCognize(img.(Image).Grow(key, value))
	default:
		panic("non-textual non-integral key")
	}
}
