// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"bytes"
	"fmt"
)

type Design interface {
	String() string
}

func stringifySlice(ss []string) string {
	var w bytes.Buffer
	for i, part := range ss {
		w.WriteString(part)
		if i+1 < len(ss) {
			w.WriteString(".")
		}
	}
	return w.String()
}

// Name
type Name string

func (x Name) String() string {
	return fmt.Sprintf("Name(%s)", x)
}

// Path
type Path []string

func (x Path) Name() Name {
	if len(x) != 1 {
		panic("path not a name")
	}
	return Name(x[0])
}

func (x Path) String() string {
	return fmt.Sprintf("Path(%s)", stringifySlice(x))
}

// RootPath
type RootPath []string

func (x RootPath) Name() Name {
	if len(x) != 1 {
		panic("path not a name")
	}
	return Name(x[0])
}

func (x RootPath) String() string {
	return fmt.Sprintf("RootPath(%s)", stringifySlice(x))
}
