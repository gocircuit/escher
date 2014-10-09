// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	"sync"

	. "github.com/gocircuit/escher/circuit"
)

// Eye is a runtime facility that delivers messages by invoking gate methods and
// provides methods that the gate can use to send messages out.
//
// Eye is an implementation of Leslie Valiant's “Mind's Eye”, described in
//	http://www.probablyapproximatelycorrect.com/
// The mind's eye is a synchronization device which sees changes as ordered
// and thus introduces the illusory perception of time (and, eventually, of the
// higher-level concepts of cause and effect).
//
type Eye struct {
	see chan *change
	show map[Name]*nerve
}

type change struct {
	Valve Name
	Value interface{}
}

func NewEye(valve ...Name) (Reflex, *Eye) {
	return NewEyeCognizer(nil, valve...)
}

type EyeCognizer func(eye *Eye, valve Name, value interface{})

func NewEyeCognizer(cog EyeCognizer, valve ...Name) (Reflex, *Eye) {
	r := make(Reflex)
	eye := &Eye{
		see: make(chan *change),
		show: make(map[Name]*nerve),
	}
	for i, v_ := range valve {
		v := v_
		x, y := NewSynapse()
		r[v] = x
		n := &nerve{
			index: i,
			ch: make(chan *ReCognizer),
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

func (eye *Eye) connect(valve Name, r *ReCognizer) {
	ch := eye.show[valve].ch 
	ch <- r
	close(ch)
}

type nerve struct {
	index int
	ch chan *ReCognizer
	sync.Mutex
	*ReCognizer
}

func (eye *Eye) Show(valve Name, v interface{}) {
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

func (eye *Eye) cognize(valve Name, v interface{}) {
	eye.see <- &change{
		Valve: valve,
		Value: v,
	}
}

func (eye *Eye) See() (valve Name, value interface{}) {
	chg := <-eye.see
	return chg.Valve, chg.Value
}

func (eye *Eye) Drain() {
	for {
		eye.See()
	}
}
