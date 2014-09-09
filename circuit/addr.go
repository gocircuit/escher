// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	"strings"
)

// Address ...
type Address string

func (a Address) Simple() string {
	if len(strings.Split(string(a), ".")) != 1 {
		panic(1)
	}
	return string(a)
}

func (a Address) String() string {
	return string(a)
}

func (a Address) Strings() []string {
	return strings.Split(string(a), ".")
}

func (a Address) Names() (walk []Name) {
	for _, w := range strings.Split(string(a), ".") {
		walk = append(walk, w)
	}
	return
}

func (a Address) Name() string {
	p := strings.Split(string(a), ".")
	return p[len(p)-1]
}
