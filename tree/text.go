// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package tree

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (rec Tree) Marshal() []byte {
	buf, err := json.Marshal(rec)
	if err != nil {
		panic(err)
	}
	return buf
}

func (rec Tree) String() string {
	var w bytes.Buffer
	for l, s := range rec {
		fmt.Fprintf(&w, "%s(%d): %v\n", l, len(s), s[len(s)-1])
	}
	return w.String()
}
