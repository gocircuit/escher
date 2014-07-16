// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package basic

import (
	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/tree"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Root.AddTerminal("reason", Reason{})
}

// Reason
type Reason struct{}

func (Reason) Materialize() think.Reflex {
	reflex, m := faculty.NewEye("Belief", "Observation", "Theory") // Create to (yet unattached) memory endpoints.
	SpawnReason(m)
	return reflex
}

type reason struct {
	ready chan struct{}
	recognizer *faculty.EyeReCognizer
}

func SpawnReason(memory *faculty.Eye) {
	go func() {
		f := &reason{  // Create the object that will handle the cognition of the reason reflex
			ready: make(chan struct{}),
		}
		f.recognizer = memory.Attach(f.ShortCognize)
		close(f.ready) // unblock the servicing of ShortCognize invocations
	}()
}

??
func (f *reason) ShortCognize(short faculty.Sentence) {
	<-f.ready
	if short[0].Value == nil || short[1].Value == nil { // If either of the two most recently updated valves are nil, inaction.
		return
	}
	switch short[2].Valve { // least recently updated valve, the one being computed
	case "Belief":
		f.recognizer.ReCognize(
			tree.Explain(short.AtAsTree("Theory"), short.AtAsTree("Observation")),
		)
	case "Observation":
		f.recognizer.ReCognize(
			tree.Predict(short.AtAsTree("Belief"), short.AtAsTree("Theory")),
		)
	case "Theory":
		f.recognizer.ReCognize(
			tree.Generalize(short.AtAsTree("Belief"), short.AtAsTree("Observation")),
		)
	}
	panic(7)
}
