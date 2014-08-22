// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"
	"sync"

	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/think"
)

func init() {
	faculty.Root.AddTerminal("Junction", Junction{})
}

// Junction
type Junction struct{}

func (Junction) Materialize() think.Reflex {
	a0Endo, a0Exo := think.NewSynapse()
	a1Endo, a1Exo := think.NewSynapse()
	a2Endo, a2Exo := think.NewSynapse()
	go func() {
		h := &junction{
			connected: make(chan struct{}),
			born:      make(chan struct{}),
		}
		h.Lock()
		defer h.Unlock()
		h.reply[0] = a0Endo.Focus(func(v interface{}) { h.Cognize(0, v) })
		h.reply[1] = a1Endo.Focus(func(v interface{}) { h.Cognize(1, v) })
		h.reply[2] = a2Endo.Focus(func(v interface{}) { h.Cognize(2, v) })
		close(h.connected)
	}()
	return think.Reflex{
		"X": a0Exo,
		"Y": a1Exo,
		"Z": a2Exo,
	}
}

type junction struct {
	connected chan struct{}
	sync.Once
	born chan struct{}
	sync.Mutex
	reply [3]*think.ReCognizer
}

func (h *junction) copy() []*think.ReCognizer {
	h.Lock()
	defer h.Unlock()
	r := make([]*think.ReCognizer, 3)
	copy(r, h.reply[:])
	return r
}

func (h *junction) Cognize(way int, v interface{}) {
	<-h.connected
	ch := make(chan struct{})
	hh := h.copy()
	for i, re_ := range hh {
		// println(fmt.Sprintf("Junction *** %T vs %T", i, way))
		if i == way {
			continue
		}
		re := re_
		go func() {
			re.ReCognize(v)
			ch <- struct{}{}
		}()
	}
	for i, _ := range hh {
		// println(fmt.Sprintf("Junction <…> %#T vs %#T ", i, way))
		if i == way {
			continue
		}
		<-ch
	}
	h.Once.Do(func() { close(h.born) })
}
