// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	"bytes"
	"fmt"
	"os"

	"github.com/hoijui/escher/be"
	cir "github.com/hoijui/escher/circuit"
)

type Help struct {
	index cir.Circuit
}

func (h *Help) Spark(_ *be.Eye, matter cir.Circuit, _ ...interface{}) cir.Value {
	h.index = matter.CircuitAt("Index")
	return nil
}

func (h *Help) Cognize(eye *be.Eye, v interface{}) {
	h.value(v)
}

func (h *Help) value(v interface{}) {
	switch u := v.(type) {
	case cir.Circuit:
		if cir.IsVerb(u) {
			fmt.Fprintf(os.Stderr, "\nThis is a verb constant equal to %v\n\n", cir.Verb(u))
		} else {
			h.circuit(u)
		}
	case int:
		fmt.Fprintf(os.Stderr, "\nThis is an integer constant equal to %v\n\n", u)
	case float64:
		fmt.Fprintf(os.Stderr, "\nThis is a float constant equal to %v\n\n", u)
	case complex128:
		fmt.Fprintf(os.Stderr, "\nThis is a complex constant equal to %v\n\n", u)
	case string:
		fmt.Fprintf(os.Stderr, "\nThis is a string constant equal to %q\n\n", u)
	default:
		fmt.Fprintf(os.Stderr, "\nThis is a value of uncommon type %T equal to %v\n\n", u, u)
	}
}

func (h *Help) circuit(u cir.Circuit) {
	var w bytes.Buffer
	fmt.Fprintf(&w, "\nWe are looking at a circuit design \n%v\n\n", u)

	valves := u.ValveNames(cir.Super)
	if len(valves) == 0 {
		fmt.Fprintf(&w, "The circuit has no super valves.\n\n")
	} else {
		fmt.Fprintf(&w, "The circuit has %d super valve(s) ", len(valves))
		cir.SortNames(valves)
		for _, vn := range valves {
			fmt.Fprintf(&w, ":%v ", vn)
		}
		w.WriteString("\n\n")
	}

	os.Stderr.Write(w.Bytes())
}
