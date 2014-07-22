// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package understand

import (
	"fmt"
)

func panicf(format string, arg ...interface{}) {
	panic(fmt.Sprintf(format, arg...))
}

func printf(format string, arg ...interface{}) {
	println(fmt.Sprintf(format, arg...))
}

func printf2(format string, arg ...interface{}) {
	fmt.Printf(format, arg...)
}
