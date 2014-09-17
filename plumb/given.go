// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package plumb

type Given struct {
	cycle chan interface{}
}

func (a *Given) Init() {
	a.cycle = make(chan interface{})
}

func (a *Given) Fix(v interface{}) {
	a.cycle <- v
}

func (a *Given) Use() interface{} {
	v := <-a.cycle
__Drain:
	for { // drain the cycle until the latest value is gotten
		select {
		case v = <-a.cycle:
		default:
			break __Drain
		}
	}
	go func() {
		a.cycle <- v // return the value to the cycle
	}()
	return v
}
