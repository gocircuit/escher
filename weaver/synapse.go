package weaver

import (
	"sync"
)

//
type Synapse struct {
	sync.Mutex
	value  Value
	reflex Reflex
	valve  Name
}

func NewSynapse() *Synapse {
	return &Synapse{}
}

func (y *Synapse) Fix(value Value) {
	y.Lock()
	defer y.Unlock()
	if y.value != nil {
		return
	}
	y.value = value
	if y.reflex != nil {
		y.reflex.Fix(y.valve, y.value)
	}
}

func (y *Synapse) Link(reflex Reflex, valve Name) {
	y.Lock()
	defer y.Unlock()
	if y.reflex != nil {
		return
	}
	y.reflex, y.valve = reflex, valve
	if y.value != nil {
		y.reflex.Fix(y.valve, y.value)
	}
}
