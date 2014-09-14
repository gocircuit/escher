// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package model

import (
	"sync"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/faculty/basic"
	"github.com/gocircuit/escher/be"
)

type Reservoir struct{
	sync.Mutex
	stop sync.Once
	y, f Circuit // output and current focus
}

func (r *Reservoir) Is() {
	r.y = New()
	r.f = r.y
}

func (r *Reservoir) CognizeY(eye *be.Eye, v interface{}) {}

/*
	{
		Command string
		Name string
		Meaning *
	}
*/
func (r *Reservoir) CognizeX(eye *be.Eye, v interface{}) {
	r.Lock()
	defer r.Unlock()
	u := v.(Circuit)
	switch u.StringAt("Command") {
	case "Open":
		r.f = r.f.CircuitAt(u.At("Name"))

	case "Close":
		r.f = r.f.CircuitAt(backtrack{})

	case "Include":
		v := u.At("Meaning")
		switch t := v.(type) {
		case Circuit:
			if _, over := t.Include(backtrack{}, r.f); over {
				panic("circuit already part of another reservoir")
			}
		}
		if _, over := r.f.Include(u.At("Name"), v); over {
			panic("over including")
		}

	case "Exclude":
		if _, forgotten := r.f.Exclude(u.At("Name")); !forgotten {
			panic("nothing to exclude")
		}

	case "Link":
		a, b := u.CircuitAt(0), u.CircuitAt(1)
		r.f.Form(
			Real{
				Image: [2]Name{a.At("Image"), b.At("Image")},
				Valve: [2]Name{a.At("Valve"), b.At("Valve")},
			},
		)

	case "Unlink":
		??

	case "Stop":
		r.stop.Do(func() { r.emit(eye) })
	}
}

func (r *Reservoir) emit(eye *be.Eye) {
	// remove all backtracks
	??
	eye.Show("Y", r.y)
}

type backtrack struct{}
