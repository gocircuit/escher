// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package time

import (
	"time"

	"github.com/gocircuit/escher/kit/plumb"
	"github.com/gocircuit/escher/think"
)

// Ticker
type Ticker struct{}

func (Ticker) Materialize() think.Reflex {
	reflex, eye := plumb.NewEye("Tick", "Duration")
	go func() {
		for {
			valve, value := eye.See()
			switch valve {
			case "Duration":
				dur := time.Duration(plumb.AsInt(value))
				if dur <= 0 {
					panic(4)
				}
				abr := make(chan struct{})
				go func() {
					start := time.Now()
					tkr := time.NewTicker(dur)
					defer tkr.Stop()
					for {
						select {
						case <-abr:
							return
						case t := <-tkr.C:
							eye.Show("Tick", int(t.Sub(start)))
						}
					}
				}()
			}
		}
	}()
	return reflex
}	
