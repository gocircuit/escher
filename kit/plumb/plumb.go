// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package plumb provides bits and bobs useful in implementing gates.
package plumb

import (
	"bytes"
	"io"
	"math"
	"strconv"
)

// AsInt accepts an int or float64 value and converts it to an int value.
func AsInt(v interface{}) int {
	switch t := v.(type) {
	case int:
		return t
	case float64:
		if math.Floor(t) == t {
			return int(t)
		}
		panic("precision")
	case complex128:
		if imag(t) != 0 {
			panic("imaginary integers")
		}
		f := real(t)
		if math.Floor(f) == f {
			return int(f)
		}
		panic("real precision")
	case string:
		i, err := strconv.Atoi(t)
		if err != nil {
			panic("illegible integer")
		}
		return i
	}
	panic(4)
}

func AsString(v interface{}) string {
	switch t := v.(type) {
	case []byte:
		return string(t)
	case string:
		return t
	case bytes.Buffer:
		return t.String()
	case io.Reader:
		var w bytes.Buffer
		io.Copy(&w, t)
		return w.String()
	}
	panic(4)
}
