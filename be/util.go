// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	"bytes"
	"fmt"
	"io"

	. "github.com/gocircuit/escher/circuit"
	// "github.com/gocircuit/escher/kit/runtime"
)

type Panic struct {
	Matter Circuit
	Msg    string
}

func panicWithMatter(matter Circuit, format string, arg ...interface{}) {
	var w bytes.Buffer
	fmt.Fprintf(&w, format, arg...)
	fmt.Fprintf(&w, "\n")
	panic(Panic{Matter: matter, Msg: w.String()})
}

func Panicf(f string, a ...interface{}) {
	var w bytes.Buffer
	fmt.Fprintf(&w, f, a...)
	fmt.Fprintf(&w, "\n")
	panic(w.String())
}

func PrintableMatter(u Circuit) string {
	var w bytes.Buffer
	PrintMatter(&w, u)
	return w.String()
}

func PrintMatter(w io.Writer, matter Circuit) {
	for {
		view, _ := matter.CircuitOptionAt("View")
		switch {
		case matter.Has("Circuit"):
			cir := matter.CircuitAt("Circuit")
			fmt.Fprintf(w, "CIRCUIT(%v) %v\n", PrintView(view), cir)

		case matter.Has("Verb"):
			verb := Verb(matter.CircuitAt("Verb"))
			addr := Verb(matter.CircuitAt("Resolved"))
			fmt.Fprintf(w, "DIRECTIVE(%v) %v/%v\n", PrintView(view), verb, addr)

		case matter.Has("System"):
			system := matter.At("System")
			fmt.Fprintf(w, "MATERIALIZE(%v) %v\n", PrintView(view), String(system))

		case matter.Has("Noun"):
			noun := matter.At("Noun")
			fmt.Fprintf(w, "NOUN(%v) %v\n", PrintView(view), noun)

		case matter.Has("Material"):
			fmt.Fprintf(w, "BASIS(%v)\n", PrintView(view))

		case matter.Has("Main"):
			fmt.Fprintf(w, "MAIN()\n")

		default:
			fmt.Fprintf(w, "UNKNOWN (%v) {%v}\n", PrintView(view), PrintView(matter))
		}
		sup, ok := matter.CircuitOptionAt("Super")
		if ok {
			matter = sup
			continue
		}
		bar, ok := matter.CircuitOptionAt("Barrier")
		if ok {
			matter = bar
			continue
		}
		break
	}
}

func PrintView(u Circuit) string {
	var w bytes.Buffer
	for i, n := range u.SortedNames() {
		if i > 0 {
			w.WriteString(" ")
		}
		fmt.Fprintf(&w, ":%v", n)
	}
	return w.String()
}
