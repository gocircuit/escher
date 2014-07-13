// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package record

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (rec Record) Marshal() []byte {
	buf, err := json.Marshal(rec)
	if err != nil {
		panic(err)
	}
	return buf
}

func (rec Record) String() string {
	var w bytes.Buffer
	for l, s := range rec {
		fmt.Fprintf(&w, "%s(%d): %v\n", l, len(s), s[len(s)-1])
	}
	return w.String()
}
