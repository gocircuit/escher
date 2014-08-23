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
	"strconv"
)

// AsInt accepts an int or float64 value and converts it to an int value.
func AsInt(v interface{}) (int, bool) {
	switch t := v.(type) {
	case nil:
		return 0, true
	case int:
		return t, true
	case float64:
		if math.Floor(t) == t {
			return int(t), true
		}
		panic("precision")
	case complex128:
		if imag(t) != 0 {
			panic("imaginary integers")
		}
		f := real(t)
		if math.Floor(f) == f {
			return int(f), true
		}
		panic("real precision")
	case string:
		i, err := strconv.Atoi(t)
		if err != nil {
			panic("illegible integer")
		}
		return i, true
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
