// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package image

import (
	// "bytes"
	// "fmt"
	. "reflect"
	"strconv"
)

func Imagine(x interface{}) interface{} {
	return imagine(ValueOf(x)).Interface()
}

func imagine(v Value) Value {
	switch v.Kind() {
	// case Map:
	case Ptr:
		w := v.Elem()
		switch w.Kind() {
		case Ptr:
			return imagine(w)
		case Struct:
			return imagine(w)
		default:
			return v
		}
	case Slice:
		img := Make()
		for i := 0; i < v.Len(); i++ {
			img.Grow(strconv.Itoa(i), imagine(v.Index(i)).Interface())
		}
		return ValueOf(img)
	case Struct:
		img := Make()
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			img.Grow(t.Field(i).Name, imagine(v.Field(i)).Interface())
		}
		return ValueOf(img)
	}
	return v
}
