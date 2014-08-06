// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

import (
	"github.com/gocircuit/escher/star"
)

func (attendant *EyeReCognizer) ReCognize(imp Impression) {
	ch := make(chan struct{})
	order := imp.Order()
	for _, f := range order {
		go func() {
			attendant.recognize[f.Valve()].ReCognize(f.Value())
			ch <- struct{}{}
		}()
	}
	for _ = range order {
		<-ch
	}
}

func (attendant *EyeReCognizer) cognizeWith(valve string, value *star.Star) {
	attendant.Lock()
	attendant.age++
	attendant.memory.Show(attendant.age, valve, value)
	reply := attendant.formulate()
	attendant.Unlock()
	attendant.cognize(reply)
}

func (attendant *EyeReCognizer) formulate() Impression {
	var sorting = attendant.memory.Order()
	imp := MakeImpression()
	for i, f := range sorting {
		imp.See(i, f.Valve(), f.Value())
	}
	return imp	
}
