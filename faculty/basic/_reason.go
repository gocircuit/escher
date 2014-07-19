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
	reflex, eye := faculty.NewEye("Belief", "Observation", "Theory") // Create to (yet unattached) memory endpoints.
	go func() {
		f := &reason{}  // Create the object that will handle the cognition of the reason reflex
		f.x = eye.Focus(f.ShortCognize)
	}()
	return reflex
}

type reason struct {
	x *faculty.EyeReCognizer
}

func (f *reason) ShortCognize(saw faculty.Sentence) {
	if saw.At(0).Value() == nil || saw.At(1).Value() == nil { // If either of the two most recently updated valves are nil, inaction.
		return
	}
	switch saw.At(2).Valve() { // least recently updated valve, the one being computed
	case "Belief":
		f.x.ReCognize(tree.Explain(saw.AtAsTree("Theory"), saw.AtAsTree("Observation")))
	case "Observation":
		f.x.ReCognize(tree.Predict(saw.AtAsTree("Belief"), saw.AtAsTree("Theory")))
	case "Theory":
		f.x.ReCognize(tree.Generalize(saw.AtAsTree("Belief"), saw.AtAsTree("Observation")))
	}
	panic(7)
}
