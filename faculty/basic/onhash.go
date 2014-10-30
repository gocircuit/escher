// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "log"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
)

type OnHash struct {
	Tag string // e.g. "#End"
}

func (h *OnHash) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	h.Tag = aux[0].(string)
	return nil
}

func (h *OnHash) CognizeView(eye *be.Eye, v interface{}) {
	if v.(Circuit).Has(h.Tag) {
		eye.Show(DefaultValve, v)
	}
}

func (h *OnHash) Cognize(eye *be.Eye, v interface{}) {}
