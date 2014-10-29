// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package runtime

import (
	"bytes"
	"runtime/pprof"
)

func PrintStack() {
	var w bytes.Buffer
	p := pprof.Lookup("goroutine")
	p.WriteTo(&w, 1)
	println(w.String())
}
