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
	"sync"
)

type MutexImage struct {
	sync.Mutex
	Image
}

func Imagine(x interface{}) interface{} {
	return imagine(false, ValueOf(x)).Interface()
}

func ImagineWithMaps(x interface{}) interface{} {
	return imagine(true, ValueOf(x)).Interface()
}

func imagine(withMaps bool, v Value) Value {
	switch v.Kind() {
	case Map:
		if !withMaps {
			break
		}
		img := Make()
		for _, k := range v.MapKeys() {
			img.Grow(k.String(), imagine(withMaps, v.MapIndex(k)).Interface())
		}
		return ValueOf(img)
	case Ptr:
		w := v.Elem()
		switch w.Kind() {
		case Ptr:
			return imagine(withMaps, w)
		case Struct:
			return imagine(withMaps, w)
		default:
			return v
		}
	case Slice:
		img := Make()
		for i := 0; i < v.Len(); i++ {
			img.Grow(strconv.Itoa(i), imagine(withMaps, v.Index(i)).Interface())
		}
		return ValueOf(img)
	case Struct:
		img := Make()
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			img.Grow(t.Field(i).Name, imagine(withMaps, v.Field(i)).Interface())
		}
		return ValueOf(img)
	}
	return v
}
