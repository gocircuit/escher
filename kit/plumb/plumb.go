// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package plumb provides bits and bobs useful in implementing gates.
package plumb

import (
	"bytes"
	"io"
	"math"
	"strconv"
	"sync"

	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/think"
)

// AsInt accepts an int or float64 value and converts it to an int value.
func AsInt(v interface{}) int {
	switch t := v.(type) {
	case int:
		return t
	case float64:
		if math.Floor(t) == t {
			return int(t)
		}
		panic("precision")
	case complex128:
		if imag(t) != 0 {
			panic("imaginary integers")
		}
		f := real(t)
		if math.Floor(f) == f {
			return int(f)
		}
		panic("real precision")
	case string:
		i, err := strconv.Atoi(t)
		if err != nil {
			panic("illegible integer")
		}
		return i
	}
	panic(4)
}

func AsString(v interface{}) string {
	switch t := v.(type) {
	case string:
		return t
	case bytes.Buffer:
		return t.String()
	case io.Reader:
		var w bytes.Buffer
		io.Copy(&w, t)
		return w.String()
	}
	panic(4)
}

// Condition …
type Condition struct {
	sync.Mutex
	ch chan interface{}
	value interface{}
	ok bool
}

func NewCondition() *Condition {
	return &Condition{
		ch: make(chan interface{}, 1),
	}
}

func (x *Condition) Determine(v interface{}) {
	x.Lock()
	defer x.Unlock()
	x.ch <- v
}

func (x *Condition) String() string {
	x.Lock()
	defer x.Unlock()
	var v interface{}
	v, x.ok = <- x.ch
	x.value = v.(string)
	return v.(string)
}

func (x *Condition) Image() Image {
	x.Lock()
	defer x.Unlock()
	var v interface{}
	v, x.ok = <- x.ch
	x.value = v.(Image)
	return v.(Image)
}

// Speak
type Speak struct {
	connect chan *think.ReCognizer
}

func NewSpeak() *Speak {
	return &Speak{
		connect: make(chan *think.ReCognizer, 1),
	}
}

func (x *Speak) Connect(r *think.ReCognizer) {
	x.connect <- r
	close(x.connect)
}

func (x *Speak) Connected() *think.ReCognizer {
	return <-x.connect
}

// Hear
type Hear struct {
	flow chan interface{}
}

func NewHear() *Hear {
	return &Hear{
		flow: make(chan interface{}, 1),
	}
}

func (x *Hear) Cognize(v interface{}) {
	x.flow <- v
}

func (x *Hear) Chan() <-chan interface{} {
	return x.flow
}

// Eye is an implementation of Leslie Valiant's “Mind's Eye”, described in
//	http://www.probablyapproximatelycorrect.com/
type Eye struct {
	see chan *change
	show map[string]*nerve
}

type change struct {
	Valve string
	Value interface{}
}

func NewEye(valve ...string) (think.Reflex, *Eye) {
	r := make(think.Reflex)
	eye := &Eye{
		see: make(chan *change),
		show: make(map[string]*nerve),
	}
	for i, v_ := range valve {
		v := v_
		x, y := think.NewSynapse()
		r[v] = x
		n := &nerve{
			index: i,
			ch: make(chan *think.ReCognizer),
		}
		eye.show[v] = n
		go eye.connect(v, y.Focus(eye.cognizeValve(v)))
	}
	return r, eye
}

func (eye *Eye) connect(valve string, r *think.ReCognizer) {
	ch := eye.show[valve].ch 
	ch <- r
	close(ch)
}

type nerve struct {
	index int
	ch chan *think.ReCognizer
	sync.Mutex
	*think.ReCognizer
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

func (eye *Eye) cognizeValve(valve string) think.Cognize {
	return func(v interface{}) {
		eye.cognize(valve, v)
	}
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
