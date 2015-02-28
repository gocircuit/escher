package weaver

import (
	"sync"
)

type Reflex interface {
	Fix(Name, Value)
	Link(Name, Reflex, Name)
}

// A reflex has a set of source valves and a set of sink valves.
type reflex struct {
	sink map[Name]*Synapse
	sync.Mutex
	expecting map[Name]struct{} // set of valves, whose values are still not received
	rule      Rule
}

func NewReflex(rule Rule) Reflex {
	x := &reflex{
		sink:      make(map[Name]*Synapse),
		expecting: make(map[Name]struct{}),
		rule:      rule,
	}
	for _, valve := range rule.Sinks() {
		x.sink[valve] = NewSynapse()
	}
	for _, valve := range rule.Sources() {
		x.expecting[valve] = struct{}{}
	}
	return x
}

func (x *reflex) Link(sink Name, reflex Reflex, valve Name) {
	x.sink[sink].Link(reflex, valve)
}

func (x *reflex) Fix(valve Name, value Value) {
	effect, remain := x.fix(valve, value)
	if !effect {
		return
	}
	if remain > 0 {
		return
	}
	// This code will execute only once
	x.rule.Spark() // Rule executes in the goroutine invoking Reflex.Fix()
	for valve, synapse := range x.sink {
		synapse.Fix(x.rule.Read(valve))
	}
}

func (x *reflex) fix(valve Name, value Value) (effect bool, remain int) {
	x.Lock()
	defer x.Unlock()
	if _, ok := x.expecting[valve]; !ok {
		return false, len(x.expecting)
	}
	delete(x.expecting, valve)
	x.rule.Write(valve, value)
	return true, len(x.expecting)
}
