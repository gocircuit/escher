// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package twitter installs a faculty for access to the Twitter API.
package twitter

import (
	"github.com/gocircuit/escher/faculty/basic"
	"github.com/gocircuit/escher/think"
)

// ForkAnswer ...
type ForkAnswer struct{}

func (ForkAnswer) Materialize() think.Reflex {
	return basic.MaterializeConjunction("_", "Name", "Sentence")
}

// ForkConsumer ...
type ForkConsumer struct{}

func (ForkConsumer) Materialize() think.Reflex {
	return basic.MaterializeConjunction("_", "Key", "Secret")
}

// ForkAccess ...
type ForkAccess struct{}

func (ForkAccess) Materialize() think.Reflex {
	return basic.MaterializeConjunction("_", "Token", "Secret")
}

// ForkUserTimelineQuery ...
type ForkUserTimelineQuery struct{}

func (ForkUserTimelineQuery) Materialize() think.Reflex {
	return basic.MaterializeConjunction("_", "UserId", "ScreenName", "AfterId", "NotAfterId", "Count")
}
