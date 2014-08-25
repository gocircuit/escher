// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package understand

import (
	"bytes"
	"fmt"
	"sort"
)

type printer interface {
	Print(string, string) string
}

func (fty Faculty) Print(prefix, indent string) string {
	var w bytes.Buffer
	var keys []string
	for k, _ := range fty {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := fty[k]
		w.WriteString(prefix)
		w.WriteString(indent)
		switch v.(type) {
		case Faculty:
		default:
			w.WriteString("*")
		}
		w.WriteString(k)
		switch t := v.(type) {
		case Faculty:
			w.WriteString(":\n")
			w.WriteString(t.Print(prefix+indent+indent, indent))
		case *Circuit:
			w.WriteString("\n")
			w.WriteString(t.Print(prefix+indent+indent, indent))
		default: // reflex or circuit
			w.WriteString(fmt.Sprintf(" [%T]\n", v))
		}
	}
	return w.String()
}

func (x *Circuit) printValves() string {
	valve := x.Peer[""].Valve
	if len(valve) == 0 {
		return ""
	}
	var w bytes.Buffer
	w.WriteString("(")
	var i int
	for v, _ := range valve {
		w.WriteString(v)
		if i + 2 < len(valve) {
			w.WriteString(", ")
		}
		i++
	}
	w.WriteString(")")
	return w.String()
}

func (x *Circuit) Print(prefix, indent string) string {
	var w bytes.Buffer
	fmt.Fprintf(&w, "%s%s%s {\n", prefix, x.Name, x.printValves())
	for _, p := range x.Peer {
		if p.Name == "" {
			continue
		}
		fmt.Fprintf(&w, "%s%s%s %v\n", prefix, indent, printable(p.Name), p.Design)
		for _, v := range p.Valve {
			fmt.Fprintf(&w, "%s%s%s%s.%s = %s.%s\n",
				prefix, indent, indent,
				printable(p.Name), printable(v.Name),
				printable(v.Matching.Of.Name), printable(v.Matching.Name),
			)
		}
	}
	fmt.Fprintf(&w, "%s}\n", prefix)
	return w.String()
}

func printable(s string) string {
	if s != "" {
		return s
	}
	return "@"
}
