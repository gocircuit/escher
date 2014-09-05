// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package plumb provides bits and bobs useful in implementing gates.
package plumb

import (
	"sync"

	// . "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/be"
)

// Eye is an implementation of Leslie Valiant's “Mind's Eye”, described in
//	http://www.probablyapproximatelycorrect.com/
// The mind's eye is a synchronization device which sees changes as ordered
// and thus introduces the illusory perception of time (and, eventually, of the
// higher-level concepts of cause and effect).
type Eye struct {
	see chan *change
	show map[string]*nerve
}

type change struct {
	Valve string
	Value interface{}
}

func NewEye(valve ...string) (be.Reflex, *Eye) {
	return NewEyeCognizer(nil, valve...)
}

type EyeCognizer func(eye *Eye, valve string, value interface{})

func NewEyeCognizer(cog EyeCognizer, valve ...string) (be.Reflex, *Eye) {
	r := make(be.Reflex)
	eye := &Eye{
		see: make(chan *change),
		show: make(map[string]*nerve),
	}
	for i, v_ := range valve {
		v := v_
		x, y := be.NewSynapse()
		r[v] = x
		n := &nerve{
			index: i,
			ch: make(chan *be.ReCognizer),
		}
		eye.show[v] = n
		if cog == nil {
			go func() {
				eye.connect(
					v,
					y.Focus(
						func(w interface{}) {
							eye.cognize(v, w)
						},
					),
				)
			}()
		} else {
			go func() {
				eye.connect(
					v,
					y.Focus(
						func(w interface{}) {
							cog(eye, v, w)
						},
					),
				)
			}()
		}
	}
	return r, eye
}

func (eye *Eye) connect(valve string, r *be.ReCognizer) {
	ch := eye.show[valve].ch 
	ch <- r
	close(ch)
}

type nerve struct {
	index int
	ch chan *be.ReCognizer
	sync.Mutex
	*be.ReCognizer
}

func (eye *Eye) Show(valve string, v interface{}) {
	n := eye.show[valve]
	r, ok := <-n.ch
	n.Lock()
	if !ok {
		r = n.ReCognizer
	} else {
		n.ReCognizer = r
	}
	n.Unlock()
	r.ReCognize(v)
}

func (eye *Eye) cognize(valve string, v interface{}) {
	eye.see <- &change{
		Valve: valve,
		Value: v,
	}
}

func (eye *Eye) See() (valve string, value interface{}) {
	chg := <-eye.see
	return chg.Valve, chg.Value
}

func (eye *Eye) Drain() {
	for {
		eye.See()
	}
}
