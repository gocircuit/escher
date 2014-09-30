// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	"bytes"
	"fmt"
	"io"
)

type Printer interface {
	Print(prefix, indent string) string
}

func (u Circuit) Print(prefix, indent string) string {
	if u.IsNil() {
		return "<nil>"
	}
	var w bytes.Buffer
	w.WriteString("{")

	// super
	if valves := u.Valves(Super); len(valves) > 0 {
		fmt.Fprintf(&w, " // ")
		if len(valves) > 0 {
			var i int
			for vn, _ := range valves {
				fmt.Fprintf(&w, "%v", vn)
				i++
				if i < len(valves) {
					w.WriteString(", ")
				}
			}
		}
	}
	w.WriteString("\n")

	// letters
	for _, n := range u.SortedLetters() {
		if len(n) > 0 && n[0] == '#' { // skip sugar gates
			continue
		}
		p := u.Gate[n]
		w.WriteString(prefix + indent)
		PrintValue(&w, prefix+indent, indent, n, p)
	}
	// numbers
	for _, n := range u.SortedNumbers() {
		p := u.Gate[n]
		w.WriteString(prefix + indent)
		PrintValue(&w, prefix+indent, indent, n, p)
	}
	//
	o := make(Orient)
	for sg, valves := range u.Flow {
		for sv, t := range valves {
			tg, tv := t.Reduce()
			if o.Has(tg, tv) {
				continue
			}
			o.Include(sg, sv)
			//
			fmt.Fprintf(&w, "%s%s%s = %s\n", 
				prefix, indent,  
				u.resugar(sg, sv),
				u.resugar(tg, tv),
			)
		}
	}
	w.WriteString(prefix + "}")
	return w.String()
}

func (u Circuit) resugar(gate, valve Name) string {
	g, ok := gate.(string)
	if !ok || len(g) == 0 || g[0] != '#' {
		return fmt.Sprintf("%s:%s", gate, valve)
	}
	return fmt.Sprintf("%s", u.Gate[gate])
}

func PrintValue(w io.Writer, prefix, indent string, n Name, p Value) {
	switch t := p.(type) {
	case Printer:
		fmt.Fprintf(w, "%v %v\n", n, t.Print(prefix, indent))
	case Address:
		fmt.Fprintf(w, "%v %s\n", n, t)
	case string:
		fmt.Fprintf(w, "%v %q\n", n, t)
	case int, float64, complex128:
		fmt.Fprintf(w, "%v %v\n", n, t)
	default:
		fmt.Fprintf(w, "%v (builtin/%T)\n", n, t)
	}
}

func Linearize(s string) string {
	x := []byte(s)
	for i, b := range x {
		if b == '\n' {
			x[i] = ','
		}
		if b == '\t' {
			x[i] = ' '
		}
	}
	return string(x)
}
