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

type Format struct {
	Prefix  string
	Indent  string
	Recurse int
}

func (u Circuit) String() string {
	var w bytes.Buffer
	u.Print(&w, Format{"", "\t", -1})
	return w.String()
}

func (u Circuit) Print(w io.Writer, f Format) {
	if u.IsNil() {
		io.WriteString(w, "<nil>")
		return
	}
	if len(u.Gate)+len(u.Flow) == 0 {
		io.WriteString(w, "{}")
		return
	}
	if f.Recurse == 0 {
		io.WriteString(w, "{…}")
		return
	}
	io.WriteString(w, "{\n")

	// letters
	for _, n := range u.SortedLetters() {
		io.WriteString(w, f.Prefix+f.Indent)
		g := Format{
			Prefix:  f.Prefix + f.Indent,
			Indent:  f.Indent,
			Recurse: f.Recurse - 1,
		}
		printGate(w, g, n, u.Gate[n])
	}
	// numbers
	for _, n := range u.SortedNumbers() {
		io.WriteString(w, f.Prefix+f.Indent)
		g := Format{
			Prefix:  f.Prefix + f.Indent,
			Indent:  f.Indent,
			Recurse: f.Recurse - 1,
		}
		printGate(w, g, n, u.Gate[n])
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
			io.WriteString(w, f.Prefix+f.Indent)
			u.resugar(w, f, sg, sv)
			io.WriteString(w, " = ")
			u.resugar(w, f, t.Gate, t.Valve)
		}
	}
	io.WriteString(w, f.Prefix+"}")
}

func (u Circuit) resugar(w io.Writer, f Format, gate, valve Name) {
	g, ok := gate.(string)
	if !ok || len(g) == 0 || g[0] != '#' {
		Print(w, f, gate)
		io.WriteString(w, ":")
		Print(w, f, valve)
		return
	}
	Print(w, f, u.Gate[gate])
}

func String(v Value) string {
	switch t := v.(type) {
	case Circuit:
		if IsVerb(t) {
			return Verb(t).String()
		} else {
			return t.String()
		}
	case int, float64, complex128, bool:
		return fmt.Sprintf("%v", t)
	case string:
		return fmt.Sprintf("%q", t)
	}
	return fmt.Sprintf("%T/%v", v, v)
}

func Print(w io.Writer, f Format, v Value) {
	switch t := v.(type) {
	case Circuit:
		if IsVerb(t) {
			Verb(t).Print(w, f)
		} else {
			t.Print(w, f)
		}
	default:
		io.WriteString(w, String(v))
	}
}

func printGate(w io.Writer, f Format, n Name, p Value) {
	Print(w, f, n)
	io.WriteString(w, " ")
	Print(w, f, p)
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

func NameString(n Name) string {
	switch n.(type) {
	case string, int:
		return fmt.Sprintf("%v", n)
	}
	panic("non alpha-numeric name has no string representation")
}
