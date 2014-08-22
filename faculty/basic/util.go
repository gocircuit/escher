// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	"bytes"
	"io"
	"math"
)

// AsInt accepts an int or float64 value and converts it to an int value.
func AsInt(v interface{}) (int, bool) {
	if v == nil {
		return 0, true
	}
	if i, ok := v.(int); ok {
		return i, true
	}
	if f, ok := v.(float64); ok && math.Floor(f) == f {
		return int(f), true
	}
	return 0, false
}

func AsString(v interface{}) (string, bool) {
	switch t := v.(type) {
	case string:
		return t, true
	case bytes.Buffer:
		return t.String(), true
	case io.Reader:
		var w bytes.Buffer
		io.Copy(&w, t)
		return w.String(), true
	case nil:
		return "", false
	}
	panic(2)
}
