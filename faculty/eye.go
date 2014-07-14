// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package faculty

import (
	"sync"

	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/tree"
	"github.com/gocircuit/escher/understand"
)

// functional combines the name of a valve and an associated value.
type functional struct {
	Valve string
	Value interface{}
}

// Memory stores a collection of valve functionals, sorted by recency of update.
// Most recent has lowest integral rank.
type Memory tree.Tree // RecencyRank:int -> Functional:functional

func (m Memory) At(valve string) interface{} {
	for _, f := range m {
		if f.Valve == valve {
			return f.Value
		}
	}
	panic(7)
}

func (m Memory) AtAsTree(valve string) tree.Tree {
	return tree.Make().Grow(valve, m.At(valve))
}

func (m Memory) NumNonNil() (n int) {
	for _, f := range m {
		if f.Value == nil {
			n++
		}
	}
	return
}

type EyeReCognizer struct {
	cognize ShortCognize
	recognize map[string]*think.ReCognizer
	sync.Mutex
	memory Memory
}

func (recognizer *EyeReCognizer) ReCognize(sentence Sentence) {
	ch := make(chan struct{})
	for valve, value := range sentence {
		go func() {
			recognizer.recognize[valve].ReCognize(value)
			ch <- struct{}{}
		}()
	}
	for _ = range sentence {
		<-ch
	}
}

func (recognizer *EyeReCognizer) cognizeOn(valve string, value interface{}) {
	recognizer.Lock()
	i := recognizer.indexOf(valve)
	recognizer.memory[0], recognizer.memory[i] = recognizer.memory[i], recognizer.memory[0]
	recognizer.memory[0].Value = value
	r := make(Sentence, len(recognizer.memory))
	recognizer.Unlock()
	//
	copy(r, recognizer.memory)
	recognizer.cognize(r)
}

// indexOf returns the current index of valve in the most-recent-first order memory slice.
func (recognizer *EyeReCognizer) indexOf(valve string) int {
	for i, meme := range recognizer.memory {
		if meme.Valve == valve {
			return i
		}
	}
	panic(7)
}
