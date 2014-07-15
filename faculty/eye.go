// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package faculty

import (
	"sort"
	"sync"

	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/tree"
	"github.com/gocircuit/escher/understand"
)

func (attendant *EyeReCognizer) ReCognize(sentence Sentence) {
	ch := make(chan struct{})
	for _, sf := range sentence {
		go func() {
			attendant.recognize[sf.Valve()].ReCognize(sf.Value())
			ch <- struct{}{}
		}()
	}
	for _ = range sentence {
		<-ch
	}
}

func (attendant *EyeReCognizer) cognize(valve string, value interface{}) {
	attendant.Lock()
	attendant.Age++
	attendant.memory[valve].Age = attendant.Age
	attendant.memory[valve].Value = value
	x := attendant.formulate()
	attendant.Unlock()
	attendant.cognize(x)
}

func (attendant *EyeReCognize) formulate() Sentence {
	var sf sortFunctional
	??
}

type sortFunctional []tree.Tree

func (sf sortFunctional) Len() int {
	return len(sf)
}

func (sf sortFunctional) Less(i, j int) bool {
	return sf[i].Int("Age") > sf[j].Int("Age")
}

func (sf sortFunctional) Swap(i, j int) {
	sf[i], sf[j] = sf[j], sf[i]
}
