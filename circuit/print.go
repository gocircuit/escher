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
	Print(prefix, indent string, recurse int) string
}

func (u Circuit) Print(prefix, indent string, recurse int) string {
	if u.IsNil() {
		return "<nil>"
	}
	if len(u.Gate) + len(u.Flow) == 0 {
		return "{}"
	}
	if recurse == 0 {
		return "{…}"
	}
	recurse--
	var w bytes.Buffer
	w.WriteString("{")

	// super
	if valves := u.ValveNames(Super); len(valves) > 0 {
		fmt.Fprintf(&w, " // ")
		SortNames(valves)
		for _, vn := range valves {
			fmt.Fprintf(&w, ":%v ", vn)
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
		PrintValue(&w, prefix+indent, indent, n, p, recurse)
	}
	// numbers
	for _, n := range u.SortedNumbers() {
		p := u.Gate[n]
		w.WriteString(prefix + indent)
		PrintValue(&w, prefix+indent, indent, n, p, recurse)
	}
	//
	o := make(Orient)
	for sg, valves := range u.Flow {
		for sv, t := range valves {
			if o.Has(t.Gate, t.Valve) {
				continue
			}
			o.Include(sg, sv)
			//
			fmt.Fprintf(&w, "%s%s%s = %s\n", 
				prefix, indent,  
				u.resugar(sg, sv),
				u.resugar(t.Gate, t.Valve),
			)
		}
	}
	w.WriteString(prefix + "}")
	return w.String()
}

func (u Circuit) resugar(gate, valve Name) string {
	g, ok := gate.(string)
	if !ok || len(g) == 0 || g[0] != '#' {
		return fmt.Sprintf("%v:%v", gate, valve)
	}
	return fmt.Sprintf("%s", printValueInline(u.Gate[gate]))
}

func printValueInline(v Value) string {
	switch t := v.(type) {
	case Vector:
		return fmt.Sprintf("%#v", t)
	case Circuit:
		return Linearize(t.String())
	case Address:
		return fmt.Sprintf("%s", t)
	case string:
		return fmt.Sprintf("%q", t)
	case int, float64, complex128:
		return fmt.Sprintf("%v", t)
	default:
		return fmt.Sprintf("(%T)", t)
	}
	panic(1)
}

type Stringer interface {
	String() string
}

func PrintValue(w io.Writer, prefix, indent string, n Name, p Value, recurse int) {
	switch t := p.(type) {
	case Printer:
		fmt.Fprintf(w, "%v %v\n", n, t.Print(prefix, indent, recurse))
	case Address:
		fmt.Fprintf(w, "%v %s\n", n, t)
	case string:
		fmt.Fprintf(w, "%v %q\n", n, t)
	case int, float64, complex128:
		fmt.Fprintf(w, "%v %v\n", n, t)
	case Stringer:
		fmt.Fprintf(w, "%v %v\n", n, t)
	default:
		fmt.Fprintf(w, "%v other/%T\n", n, t)
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
