// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package time

import (
	"sync"
	"time"

	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/kit/plumb"
	"github.com/gocircuit/escher/think"
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
			connected: make(chan struct{}),
			born:      make(chan struct{}),
		}
		go func() {
			h.x.Bind(xEndo.Focus(h.CognizeX))
			h.y.Bind(yEndo.Focus(h.CognizeY))
			close(h.connected)
		}()
		go func() {
			durEndo.Focus(h.CognizeDuration)
		}()
	}()
	return think.Reflex{
		"X":        xExo,
		"Y":        yExo,
		"Duration": durExo,
	}
}

type delay struct {
	x, y      think.PtrReCognizer
	connected chan struct{}
	sync.Once
	born chan struct{}
	sync.Mutex
	dur time.Duration
}

func (h *delay) CognizeDuration(v interface{}) {
	i, ok := plumb.OptionallyInt(v)
	if !ok {
		panic("non-numeric delay duration")
	}
	h.dur = time.Duration(i)
	h.Once.Do(func() { close(h.born) })
}

func (h *delay) CognizeX(v interface{}) {
	<-h.connected
	<-h.born
	h.Lock()
	dur := h.dur
	h.Unlock()
	go func() {
		time.Sleep(dur)
		h.y.ReCognize(v)
	}()
}

func (h *delay) CognizeY(v interface{}) {
	<-h.connected
	<-h.born
	h.Lock()
	dur := h.dur
	h.Unlock()
	go func() {
		time.Sleep(dur)
		h.x.ReCognize(v)
	}()
}
