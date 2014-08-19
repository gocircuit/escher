// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package time

import (
	"sync"
	"time"

	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/faculty/basic"
)

func init() {
	ns := faculty.Root.Refine("time")
	ns.AddTerminal("Ticker", Ticker{})
	ns.AddTerminal("Delay", Delay{})
}

// Delay…
type Delay struct{}

func (Delay) Materialize() think.Reflex {
	xEndo, xExo := think.NewSynapse()
	yEndo, yExo := think.NewSynapse()
	durEndo, durExo := think.NewSynapse()
	go func() {
		h := &delay{
			born: make(chan struct{}),
		}
		go func() {
			h.xRe = xEndo.Focus(h.CognizeX)
		}()
		go func() {
			h.yRe = yEndo.Focus(h.CognizeY)
		}()
		go func() {
			durEndo.Focus(h.CognizeDuration)
		}()
	}()
	return think.Reflex{
		"X": xExo, 
		"Y": yExo, 
		"Duration": durExo,
	}
}

type delay struct {
	xRe *think.ReCognizer
	yRe *think.ReCognizer
	sync.Once
	born chan struct{}
	sync.Mutex
	dur time.Duration
}

func (h *delay) CognizeDuration(v interface{}) {
	i, ok := basic.AsInt(v)
	if !ok {
		panic("non-numeric delay duration")
	}
	h.dur = time.Duration(i)
	h.Once.Do(func() { close(h.born) })
}

func (h *delay) CognizeX(v interface{}) {
	<-h.born
	h.Lock()
	dur := h.dur
	h.Unlock()
	go func() {
		time.Sleep(dur)
		println("x->y")
		h.yRe.ReCognize(v)
	}()
}

func (h *delay) CognizeY(v interface{}) {
	<-h.born
	h.Lock()
	dur := h.dur
	h.Unlock()
	go func() {
		time.Sleep(dur)
		println("y->x")
		h.xRe.ReCognize(v)
	}()
}
