// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package memory

import (
	// "fmt"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/kit/plumb"
	. "github.com/gocircuit/escher/kit/reservoir"
)

func init() {
	faculty.Register("memory.Memory", be.NewNativeMaterializer(&MemoryNoun{}))
	faculty.Register("memory.Get", be.NewNativeMaterializer(&MemoryVerb{}))
	faculty.Register("memory.Put", be.NewNativeMaterializer(&MemoryVerb{}))
	faculty.Register("memory.Forget", be.NewNativeMaterializer(&MemoryVerb{}))
}

// MemoryNoun
type MemoryNoun struct{
	Memory
}

func (r *MemoryNoun) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	if len(aux) == 1 {
		r.Memory = aux[0].(Memory)
	} else {
		r.Memory = NewMemory()
	}
	go func() {
		for vlv, _ := range matter.View.Gate {
			eye.Show(vlv, r.Memory)
		}
	}()
	return New().
		Grow("Memory", be.NewNativeMaterializer(&MemoryNoun{}, r.Memory)).
		Grow("Put", be.NewNativeMaterializer(&MemoryVerb{}, r.Memory)).
		Grow("Get", be.NewNativeMaterializer(&MemoryVerb{}, r.Memory)).
		Grow("Forget", be.NewNativeMaterializer(&MemoryVerb{}, r.Memory))
}

func (r *MemoryNoun) Cognize(eye *be.Eye, valve Name, value interface{}) {
	eye.Show(valve, r.Memory)
}

// MemoryVerb
type MemoryVerb struct {
	receiver plumb.Given
}

func (r *MemoryVerb) Spark(eye *be.Eye, _ *be.Matter, aux ...interface{}) Value {
	r.receiver.Init()
	if len(aux) == 1 {
		r.receiver.Fix(aux[0].(Memory))
	}
	return r
}

func (r *MemoryVerb) OverCognize(eye *be.Eye, valve Name, value interface{}) {
	v := value.(Circuit)
	rsrv := r.receiver.Use().(Memory)
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

func (r *MemoryVerb) CognizeUsing(eye *be.Eye, value interface{}) {
	r.receiver.Fix(value.(Memory))
}
