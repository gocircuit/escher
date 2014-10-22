// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	// "fmt"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/kit/plumb"
	. "github.com/gocircuit/escher/kit/reservoir"
)

// ReservoirNoun
type ReservoirNoun struct{
	Reservoir
}

func (r *ReservoirNoun) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	if len(aux) == 1 {
		r.Reservoir = aux[0].(Reservoir)
	} else {
		r.Reservoir = NewReservoir()
	}
	go func() {
		for vlv, _ := range matter.View.Gate {
			eye.Show(vlv, r.Reservoir)
		}
	}()
	return New().
		Grow("Reservoir", be.NewNativeMaterializer(&ReservoirNoun{}, r.Reservoir)).
		Grow("Put", be.NewNativeMaterializer(&ReservoirVerb{}, r.Reservoir)).
		Grow("Get", be.NewNativeMaterializer(&ReservoirVerb{}, r.Reservoir)).
		Grow("Forget", be.NewNativeMaterializer(&ReservoirVerb{}, r.Reservoir))
}

func (r *ReservoirNoun) Cognize(eye *be.Eye, valve Name, value interface{}) {
	eye.Show(valve, r.Reservoir)
}

// ReservoirVerb
type ReservoirVerb struct {
	receiver plumb.Given
}

func (r *ReservoirVerb) Spark(eye *be.Eye, _ *be.Matter, aux ...interface{}) Value {
	r.receiver.Init()
	if len(aux) == 1 {
		r.receiver.Fix(aux[0].(Reservoir))
	}
	return r
}

func (r *ReservoirVerb) OverCognize(eye *be.Eye, valve Name, value interface{}) {
	v := value.(Circuit)
	rsrv := r.receiver.Use().(Reservoir)
	switch valve {
	case "Put":
		addr := plumb.AsAddress(v.At("Address"))
		before := rsrv.Put(addr, v.At("Value"))
		eye.Show(DefaultValve, New().Grow("Address", addr).Grow("Before", before))
	case "Get":
		addr := plumb.AsAddress(v.At("Address"))
		eye.Show(DefaultValve, New().Grow("Address", addr).Grow("Value", rsrv.Get(addr)))
	case "Forget":
		addr := plumb.AsAddress(v.At("Address"))
		forgot := rsrv.Forget(addr)
		eye.Show(DefaultValve, New().Grow("Address", addr).Grow("Forgot", forgot))
	default:
		panic("unkonwn reservoir command")
	}
}

func (r *ReservoirVerb) CognizeUsing(eye *be.Eye, value interface{}) {
	r.receiver.Fix(value.(Reservoir))
}
