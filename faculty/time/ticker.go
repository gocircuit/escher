// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package time

import (
	"log"
	"sync"
	"time"

	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Root.AddTerminal("ticker", Ticker{})
}

// Ticker
type Ticker struct{}

func (Ticker) Materialize() think.Reflex {
	dur1, dur2 := think.NewSynapse()
	sig1, sig2 := think.NewSynapse()
	go func() {
		x := &ticker{}
		x.re = sig1.Focus(think.DontCognize)
		dur1.Focus(x.CognizeDuration)
	}()
	return think.Reflex{"Tick": sig2, "Duration": dur2}
}

type ticker struct {
	re *think.ReCognizer
	sync.Mutex
	abr chan struct{}
}

func (t *ticker) CognizeDuration(v interface{}) {
	var nano int
	switch t := v.(type) {
	case int:
		nano = t
	case float64:
		nano = int(t)
	default:
		log.Printf("non-integer ticker duration")
		return
	}
	t.Lock()
	defer t.Unlock()
	if t.abr != nil {
		close(t.abr)
		t.abr = nil
	}
	if nano <= 0 {
		return
	}
	t.abr = make(chan struct{})
	go tickerLoop(time.Duration(nano), t.abr, t.re)
}

func tickerLoop(dur time.Duration, abr <-chan struct{}, re *think.ReCognizer) {
	if dur <= 0 {
		panic(1)
	}
	start := time.Now()
	tkr := time.NewTicker(dur)
	defer tkr.Stop()
	for {
		select {
		case <-abr:
			return
		case t := <-tkr.C:
			re.ReCognize(int(t.Sub(start)))
		}
	}
}
