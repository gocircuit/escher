// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"
	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Root.AddTerminal("Grow", Grow{})
}

// Grow
// XXX: Redo from basics
type Grow struct{}

func (Grow) Materialize() think.Reflex {
	reflex, eye := faculty.NewEye("Img", "Key", "Value", "_")
	go func() {
		x := &grow{
			connected: make(chan struct{}),
		}
		x.EyeNerve = eye.Focus(x.ShortCognize)
		close(x.connected)
	}()
	return reflex
}

type grow struct {
	connected chan struct{}
	*faculty.EyeNerve
}

func (s *grow) ShortCognize(imp faculty.Impression) {
	<-s.connected
	img, ik := imp.Valve("Img").Value().(Image)
	key, kk := imp.Valve("Key").Value().(string)
	value := imp.Valve("Value").Value()
	if !ik || !kk {
		return
	}
	y := img.Abandon(key).Grow(key, value)
	z := faculty.MakeImpression().Show(0, "_", y)
	go s.ReCognize(z)
}
