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

// OptionallyInt accepts an int or float64 value and converts it to an int value.
func OptionallyInt(v interface{}) (int, bool) {
	switch t := v.(type) {
	case nil:
		return 0, true
	case int:
		return t, true
	case float64:
		if math.Floor(t) == t {
			return int(t), true
		}
		panic("precision")
	case complex128:
		if imag(t) != 0 {
			panic("imaginary integers")
		}
		f := real(t)
		if math.Floor(f) == f {
			return int(f), true
		}
		panic("real precision")
	case string:
		i, err := strconv.Atoi(t)
		if err != nil {
			panic("illegible integer")
		}
		return i, true
	}
	return 0, false
}

func OptionallyString(v interface{}) (string, bool) {
	switch t := v.(type) {
	case string:
		return t, true
	case bytes.Buffer:
		return t.String(), true
	case io.Reader:
		var w bytes.Buffer
		io.Copy(&w, t)
		return w.String(), true
	case nil:
		return "", false
	}
	panic(2)
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
	??
}
