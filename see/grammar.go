// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"bytes"
	"fmt"
	"strings"
)

type Design interface {
	String() string
}

func stringifySlice(ss []string) string {
	var w bytes.Buffer
	for i, part := range ss {
		w.WriteString(part)
		if i+1 < len(ss) {
			w.WriteString(".")
		}
	}
	return w.String()
}

// Name
type Name string

func NewName(walk []string) Name {
	return Name(strings.Join(walk, "."))
}

func (x Name) String() string {
	return fmt.Sprintf("Name(%s)", string(x))
}
