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

	cir "github.com/gocircuit/escher/circuit"
)

type Panic struct {
	Matter cir.Circuit
	Msg    string
}

func panicWithMatter(matter cir.Circuit, format string, arg ...interface{}) {
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

func PrintableMatter(u cir.Circuit) string {
	var w bytes.Buffer
	PrintMatter(&w, u)
	return w.String()
}

func PrintMatter(w io.Writer, matter cir.Circuit) {
	for {
		view, _ := matter.CircuitOptionAt("View")
		switch {
		case matter.Has("Circuit"):
			cir := matter.CircuitAt("Circuit")
			fmt.Fprintf(w, "CIRCUIT(%v)%v %v\n", PrintView(view), SummarizeIndex(matter), cir)

		case matter.Has("Verb"):
			verb := cir.Verb(matter.CircuitAt("Verb"))
			addr := cir.Verb(matter.CircuitAt("Resolved"))
			fmt.Fprintf(w, "DIRECTIVE(%v)%v %v/%v\n", PrintView(view), SummarizeIndex(matter), verb, addr)

		case matter.Has("System"):
			system := matter.At("System")
			fmt.Fprintf(w, "MATERIALIZE(%v)%v %v\n", PrintView(view), SummarizeIndex(matter), cir.String(system))

		case matter.Has("Noun"):
			noun := matter.At("Noun")
			fmt.Fprintf(w, "NOUN(%v)%v %v\n", PrintView(view), SummarizeIndex(matter), noun)

		case matter.Has("Material"):
			fmt.Fprintf(w, "BASIS(%v)%v\n", PrintView(view), SummarizeIndex(matter))

		case matter.Has("Main"):
			fmt.Fprintf(w, "MAIN()\n")

		default:
			fmt.Fprintf(w, "UNKNOWN(%v)%v\n", PrintView(view), SummarizeIndex(matter))
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

func PrintView(u cir.Circuit) string {
	var w bytes.Buffer
	for i, n := range u.SortedNames() {
		if i > 0 {
			w.WriteString(" ")
		}
		fmt.Fprintf(&w, ":%v", n)
	}
	return w.String()
}

func SummarizeIndex(matter cir.Circuit) string {
	x := matter.CircuitAt("Index")
	var w bytes.Buffer
	w.WriteString(" Index{ ")
	for i, n := range x.SortedNames() {
		if i > 2 {
			break
		}
		w.WriteString(fmt.Sprintf("%v", n))
		w.WriteString("… ")
	}
	w.WriteString("}")
	return w.String()
}
