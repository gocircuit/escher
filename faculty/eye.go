// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

import (
	"github.com/gocircuit/escher/think"
)

func (nerve *EyeNerve) lookup(name string) *think.ReCognizer {
	nerve.recognize.Lock()
	defer nerve.recognize.Unlock()
	return nerve.recognize.t[name]
}

func (nerve *EyeNerve) bind(name string, q *think.ReCognizer) {
	nerve.recognize.Lock()
	defer nerve.recognize.Unlock()
	nerve.recognize.t[name] = q
}

func (nerve *EyeNerve) ReCognize(imp Impression) {
	<-nerve.connected
	ch := make(chan struct{})
	order := imp.Order()
	for _, f_ := range order {
		f := f_
		go func() {
			nerve.lookup(f.Valve()).ReCognize(f.Value())
			ch <- struct{}{}
		}()
	}
	for _ = range order {
		<-ch
	}
}

func (nerve *EyeNerve) cognizeWith(valve string, value interface{}) {
	<-nerve.connected
	nerve.memory.Lock()
	nerve.memory.Age++
	nerve.memory.Imp.Show(nerve.memory.Age, valve, value)
	reply := nerve.formulate()
	nerve.memory.Unlock()
	nerve.cognize(reply)
}

func (nerve *EyeNerve) formulate() Impression {
	var sorting = nerve.memory.Imp.Order()
	imp := MakeImpression()
	for i, f := range sorting {
		imp.Show(i, f.Valve(), f.Value())
	}
	return imp	
}
