// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"github.com/gocircuit/escher/a"
)

func SpaceNoNewline(src *a.Src) {
	if len(a.Whitespace(src)) > 0 {
		return
	}
	panic("whitespace")
}

func Space(src *a.Src) (newLine bool) {
	for endOfLine(src) {
		newLine = true
	}
	if src.Len() == 0 || src.RuneAt(0) == '}' {
		newLine = true
	}
	return
}

func endOfLine(src *a.Src) bool {
	a.Whitespace(src)
	return len(src.Consume(a.IsCommaOrSemicolonOrNewline)) > 0
}
