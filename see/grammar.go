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

// Path
type Path struct {
	walk []string
}

func NewPath(walk []string) *Path {
	return &Path{walk}
}

func (x *Path) String() string {
	return fmt.Sprintf("Path(%s)", stringifySlice(x.walk))
}

// RootPath
type RootPath struct {
	walk []string
}

func NewRootPath(walk []string) *RootPath {
	return &RootPath{walk}
}

func (x RootPath) String() string {
	return fmt.Sprintf("RootPath(%s)", stringifySlice(x.walk))
}
