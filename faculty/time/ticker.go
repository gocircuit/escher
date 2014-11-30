// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package time

import (
	"time"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/kit/plumb"
)

// Ticker
type Ticker struct {
	ctl chan time.Duration
}

func (t *Ticker) Spark(eye *be.Eye, _ Circuit, _ ...interface{}) Value {
	t.ctl = make(chan time.Duration)
	go func() {
		var start time.Time
		var tkr *time.Ticker
		for {
			select {
			case dur := <-t.ctl:
				if tkr != nil {
					tkr.Stop()
					tkr = nil
				}
				if dur > 0 {
					start, tkr = time.Now(), time.NewTicker(dur)
				}
			case t := <-tkr.C:
				eye.Show(DefaultValve, int(t.Sub(start)))
			}
		}
	}()
	return nil
}

func (t *Ticker) CognizeDuration(eye *be.Eye, value interface{}) {
	t.ctl <- time.Duration(plumb.AsInt(value))
}

func (t *Ticker) Cognize(eye *be.Eye, value interface{}) {}
