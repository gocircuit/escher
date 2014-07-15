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

func (recognizer *EyeReCognizer) ReCognize(sentence Sentence) {
	ch := make(chan struct{})
	for _, funcl := range sentence {
		go func() {
			recognizer.recognize[funcl.String("Valve")].ReCognize(funcl.At("Value"))
			ch <- struct{}{}
		}()
	}
	for _ = range sentence {
		<-ch
	}
}

func (recognizer *EyeReCognizer) cognizeOn(valve string, value interface{}) {
	recognizer.Lock()
	recognizer.Age++
	??
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
