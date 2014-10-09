// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	// "fmt"

	. "github.com/gocircuit/escher/circuit"
)

// TODO: Rewrite union to not use a permanent goroutine

func MaterializeUnion(field ...Name) (Reflex, Value) {
	reflex, eye := NewEye(append(field, DefaultValve)...) // add the default valve
	go func() {
		conj := New()
		for {
			dvalve, dvalue := eye.See()
			if dvalve == DefaultValve { // conjunction updated
				y := make(chan struct{}) // block and
				for _, f_ := range field { // send updated conjunction to all field valves
					f := f_
					go func() {
						eye.Show(f, dvalue.(Circuit).At(f))
						y <- struct{}{}
					}()
				}
				for _ = range field {
					<-y
				}
			} else { // field updated
				conj.Abandon(dvalve).Grow(dvalve, dvalue)
				if conj.Len() == len(field) {
					eye.Show(DefaultValve, conj)
				}
			}
		}
	}()
	return reflex, 
		func() (Reflex, Value) {
			return MaterializeUnion(field...)
		}
}
