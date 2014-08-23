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
	"github.com/gocircuit/escher/think"
)

func init() {
	faculty.Root.AddTerminal("Grow", Grow{})
}

// Grow
type Grow struct{}

func (Grow) Materialize() think.Reflex {
	imgEndo, imgExo := think.NewSynapse()
	keyEndo, keyExo := think.NewSynapse()
	valueEndo, valueExo := think.NewSynapse()
	_Endo, _Exo := think.NewSynapse()
	go func() {
		h := &grow{
			connected: make(chan struct{}),
			key:       make(chan interface{}),
			img:       make(chan interface{}),
			value:     make(chan interface{}),
		}
		h.z = _Endo.Focus(think.DontCognize)
		close(h.connected)
		keyEndo.Focus(h.CognizeKey)
		imgEndo.Focus(h.CognizeImg)
		valueEndo.Focus(h.CognizeValue)
		go h.loop()
	}()
	return think.Reflex{
		"Img":   imgExo,
		"Key":   keyExo,
		"Value": valueExo,
		"_":     _Exo,
	}
}

type grow struct {
	connected       chan struct{}
	key, img, value chan interface{}
	z               *think.ReCognizer
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
	h.z.ReCognize(img.(Image).Grow(key.(string), value))
}
