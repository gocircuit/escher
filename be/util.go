// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	"fmt"
	"os"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/kit/runtime"
)

func Panicf(f string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, f, a...)
	fmt.Fprintf(os.Stderr, "\n")
	runtime.PrintStack()
	os.Exit(1)
}

// Printable matter returns a deep copy of u wherein the values of all gates named Index,
// which are circuits, are truncated to depth one.
func PrintableMatter(u Circuit) Circuit {
	r := New()
	for n, v := range u.Gate {
		if n != "Index" {
			switch t := v.(type) {
			case Circuit:
				r.Include(n, PrintableMatter(t))
			default:
				r.Include(n, t)
			}
			continue
		}
		// x := v.(Circuit)
		// s := New()
		// for xn, xv := range x.Gate {
		// 	switch xv.(type) {
		// 	case Circuit:
		// 		s.Include(xn, "…")
		// 	default:
		// 		s.Include(xn, xv)
		// 	}
		// }
		// r.Include(n, s)
		r.Include(n, "…")
	}
	return r
}
