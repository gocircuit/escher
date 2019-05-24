// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	"fmt"
)

// Vector identifies a specific valve of a gate (== instance of a circuit)
type Vector struct {
	Gate Name
	Valve Name
}

// String prints a string representation of the vector
func (v Vector) String() string {
	return fmt.Sprintf("%v:%v", v.Gate, v.Valve)
}
