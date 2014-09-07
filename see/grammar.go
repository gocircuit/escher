// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	. "github.com/gocircuit/escher/image"
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

func NewName(walk []string) Name {
	return Name(strings.Join(walk, "."))
}

func (x Name) AsWalk() (walk []interface{}) {
	for _, w := range strings.Split(string(x), ".") {
		walk = append(walk, Name(w))
	}
	return
}

func (x Name) String() string {
	return string(x)
}

func Names(img Image) []Name {
	series := make(nameSlice, 0, len(img))
	for key, _ := range x {
		k, ok := key.(Name)
		if !ok {
			continue
		}
		series = append(series, k)
	}
	sort.Sort(series)
	return series
}

// nameSlice
type nameSlice []Name

func (x nameSlice) Len() int {
	return len(x)
}

func (x nameSlice) Less(i, j int) bool {
	return x[i] < x[j]
}

func (x numberSlice) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

// Number is a name type indicating index of element in the slice syntax
type Number int

func (x Number) String() string {
	return fmt.Sprintf("#%d", int(x))
}

func Numbers(img Image) []Number {
	series := make(numberSlice, 0, len(img))
	for key, _ := range x {
		k, ok := key.(Number)
		if !ok {
			continue
		}
		series = append(series, k)
	}
	sort.Sort(series)
	return series
}

// numberSlice
type numberSlice []Number

func (x numberSlice) Len() int {
	return len(x)
}

func (x numberSlice) Less(i, j int) bool {
	return x[i] < x[j]
}

func (x numberSlice) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
