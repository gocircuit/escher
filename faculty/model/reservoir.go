// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package model

import (
	"container/list"
	"sync"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
)

type Reservoir struct{
	sync.Mutex
	stop sync.Once
	y, focus Circuit // output and current focus
	path list.List
}

func (r *Reservoir) Spark() {
	r.y = New()
	r.focus = r.y
}

func (r *Reservoir) CognizeY(eye *be.Eye, v interface{}) {}

func (r *Reservoir) CognizeX(eye *be.Eye, v interface{}) {
	r.Lock()
	defer r.Unlock()
	u := v.(Circuit)
	switch u.StringAt("Command") {
	case "Enter":
		r.path.PushBack(r.focus)
		r.focus = r.focus.CircuitAt(u.At("Gate"))

	case "Return":
		r.focus = r.path.Remove(r.path.Back()).(Circuit)

	case "Include":
		if _, over := r.focus.Include(u.At("Gate"), u.At("Value")); over {
			panic("over including")
		}

	case "Exclude":
		if _, forgotten := r.focus.Exclude(u.At("Gate")); !forgotten {
			panic("nothing to exclude")
		}

	case "Link":
		a, b := Vector(u.CircuitAt(0)), Vector(u.CircuitAt(1))
		r.focus.Link(a, b)

	case "Unlink":
		a, b := Vector(u.CircuitAt(0)), Vector(u.CircuitAt(1))
		r.focus.Unlink(a, b)

	case "Yield":
		r.stop.Do(func() { eye.Show("Y", r.y) })
	}
}
