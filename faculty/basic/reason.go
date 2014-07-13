// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package basic

import (
	"fmt"

	"github.com/gocircuit/escher/kit/record"
	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Root.AddTerminal("reason", Reason{})
}

// Reason
type Reason struct{}

func (Reason) Materialize() think.Reflex {
	reflex, m := faculty.NewShortMemory("Belief", "Observation", "Theory") // Create to (yet unattached) memory endpoints.
	SpawnReason(m)
	return reflex
}

type reason struct {
	ready chan struct{}
	recognizer *faculty.ShortMemoryReCognizer
}

func SpawnReason(memory *faculty.ShortMemory) {
	go func() {
		f := &reason{  // Create the object that will handle the cognition of the reason reflex
			ready: make(chan struct{}),
		}
		f.recognizer = memory.Attach(f.ShortCognize)
		close(f.ready) // unblock the servicing of ShortCognize invocations
	}()
}

func (f *reason) ShortCognize(short []Sentence) {
	<-f.ready
	if short[0] == nil || short[1] == nil { // If either of the two most recently updated valves are nil, inaction.
		return
	}
	switch short[2].Valve { // least recently updated valve, the one being computed
	case "Belief":
		f.recognizer.Recognize(
			record.Explain(
				faculty.AtValve("Theory", short), 
				faculty.AtValve("Observation", short),
			),
		)
	case "Observation":
		f.recognizer.Recognize(
			record.Predict(
				faculty.AtValve("Belief", short), 
				faculty.AtValve("Theory", short),
			),
		)
	case "Theory":
		f.recognizer.Recognize(
			record.Generalize(
				faculty.AtValve("Belief", short), 
				faculty.AtValve("Observation", short),
			),
		)
	}
	panic(7)
}
