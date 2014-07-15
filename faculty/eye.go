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

func (re *EyeReCognizer) ReCognize(sentence Sentence) {
	ch := make(chan struct{})
	for _, funcl := range sentence {
		go func() {
			re.recognize[funcl.String("Valve")].ReCognize(funcl.At("Value"))
			ch <- struct{}{}
		}()
	}
	for _ = range sentence {
		<-ch
	}
}

func (re *EyeReCognizer) cognizeOn(valve string, value interface{}) {
	re.Lock()
	re.Age++
	re.memory[valve].Age = re.Age
	re.memory[valve].Value = value
	x := re.formulate()
	re.Unlock()
	re.cognize(x)
}

func (re *EyeReCognize) formulate() Sentence {
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
