// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	"testing"
)

func TestSynapse(t *testing.T) {
	x, y := NewSynapse()
	p, q := NewSynapse()
	go Link(y, p)
	go Link(p, y)
	ch := make(chan int)
	go func() {
		x.Connect(DontCognize).ReCognize(1)
		ch <- 1
	}()
	go func() {
		q.Connect(DontCognize).ReCognize(1)
		ch <- 1
	}()
	<-ch
	<-ch
}
