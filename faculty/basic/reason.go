// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	"fmt"
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
	reflex, eye := faculty.NewEye("Belief", "Observation", "Theory") // Create to (yet unattached) memory endpoints.
	go func() {
		f := &reason{  // Create the object that will handle the cognition of the reason reflex
			ready: make(chan struct{}),
		}
		f.x = eye.Focus(f.ShortCognize)
		close(f.ready)
	}()
	return reflex
}

type reason struct {
	ready chan struct{}
	x *faculty.EyeReCognizer
}

func (f *reason) ShortCognize(saw faculty.Sentence) {
	println(fmt.Sprintf("saw=%v", saw))
	<-f.ready
	if saw.At(0).Value() == nil || saw.At(1).Value() == nil { // If either of the two most recently updated valves are nil, inaction.
		return
	}
	switch saw.At(2).Valve() { // least recently updated valve, the one being computed
	case "Belief":
		f.x.ReCognize(
			faculty.MakeSentence().Grow(
				0,
				"Belief",
				tree.Explain(saw.AtName("Theory").TreeValue(), saw.AtName("Observation").TreeValue()),
			),
		)
	case "Observation":
		f.x.ReCognize(
			faculty.MakeSentence().Grow(
				0,
				"Observation",
				tree.Predict(saw.AtName("Belief").TreeValue(), saw.AtName("Theory").TreeValue()),
			),
		)
	case "Theory":
		f.x.ReCognize(
			faculty.MakeSentence().Grow(
				0,
				"Theory",
				tree.Generalize(saw.AtName("Belief").TreeValue(), saw.AtName("Observation").TreeValue()),
			),
		)
	}
	panic(7)
}
