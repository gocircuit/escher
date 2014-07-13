// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package faculty

import (
	"sync"

	"github.com/petar/maymounkov.io/escher/think"
	"github.com/petar/maymounkov.io/escher/understand"
)

// Root is a global variable where packages can add gates as side-effect of being imported.
var Root = understand.NewFaculty()

// Sentence combines the name of a valve and an associated value.
type Sentence struct {
	Valve string
	Value interface{}
}

func AtValve(valve string, short []Sentence) interface{} {
	for _, s := range short {
		if s.Valve == valve {
			return s.Value
		}
	}	
}

func NumNonNil(short []Sentence) (n int) {
	for _, s := range short {
		if s.Value == nil {
			n++
		}
	}
}

// …
// The first entry is the most recent one.
type ShortCognize func([]Sentence)

type ShortMemory struct {
	y []*think.Memory
	recognizer ShortMemoryReCognizer
}

// 
func NewShortMemory(valve ...string) (think.Reflex, *ShortMemory) {
	reflex := make(think.Reflex)
	m := &ShortMemory{
		y: make([]*think.Memory, len(valve)),
		recognizer: ShortMemoryReCognizer{
			recognize: make(map[string]*think.ReCognizer),
			memory: make([]Sentence, len(valve)),
		},
	}
	for i, v := range valve {
		if _, ok := reflex[v]; ok {
			panic("duplicate valve")
		}
		reflex[v], m.y[i] = think.NewMemory()
		m.recognizer.memory[i] = Sentence{Valve: v, Value: nil}
	}
	return reflex, m
}

func (m *ShortMemory) Attach(cognize ShortCognize) *ShortMemoryReCognizer {
	// Locking prevents individual completing Attach invocations 
	// from initiating cogntion, before all valves have been attached.
	m.recognizer.Lock()
	defer m.recognizer.Unlock()
	//
	m.recognizer.cognize = cognize
	for i, s := range m.recognizer.memory {
		s_ := s
		m.recognizer.recognize[s.Valve] = m.y[i].Attach(
			func(w interface{}) {
				m.recognizer.cognizeOn(s_.Valve, w)
			},
		)
	}
	return &m.recognizer
}

type ShortMemoryReCognizer struct {
	cognize ShortCognize
	recognize map[string]*think.ReCognizer
	sync.Mutex
	memory []Sentence // most-recent to least recent
}

func (recognizer *ShortMemoryReCognizer) ReCognize(sentence []Sentence) {
	for _, s := range sentence {
		go func() {
			recognizer.recognize[s.Valve].ReCognize(s.Value)
		}()
	}
}

func (recognizer *ShortMemoryReCognizer) cognizeOn(valve string, value interface{}) {
	recognizer.Lock()
	i := recognizer.indexOf(valve)
	recognizer.memory[0], recognizer.memory[i] = recognizer.memory[i], recognizer.memory[0]
	recognizer.memory[0].Value = value
	r := make([]Sentence, len(recognizer.memory))
	recognizer.Unlock()
	//
	copy(r, recognizer.memory)
	recognizer.cognize(r)
}

// indexOf returns the current index of valve in the most-recent-first order memory slice.
func (recognizer *ShortMemoryReCognizer) indexOf(valve string) int {
	for i, meme := range recognizer.memory {
		if meme.Valve == valve {
			return i
		}
	}
	panic(7)
}
