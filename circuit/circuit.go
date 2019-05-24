// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	"fmt"
	"sort"
)

// Circuit ...
type Circuit struct {
	Gate map[Name]Value
	Flow map[Name]map[Name]Vector // gate -> valve -> opposing gate and valve
}

const Super = ""

// New creates a new circuit without gates nor flows
func New() Circuit {
	return Circuit{
		Gate: make(map[Name]Value),
		Flow: make(map[Name]map[Name]Vector),
	}
}

// IsNil checks whether the argument circuit is uninitialized
func (u Circuit) IsNil() bool {
	return u.Gate == nil || u.Flow == nil
}

// IsEmpty checks whether the argument circuit has no gates nor flows
func (u Circuit) IsEmpty() bool {
	return len(u.Gate) == 0 && len(u.Flow) == 0
}

// SortedLetters returns a sorted list of all gate IDs that are strings
func (u Circuit) SortedLetters() []string {
	x := u.Letters()
	sort.Strings(x)
	return x
}

// Letters returns a list of all gate IDs that are strings
func (u Circuit) Letters() []string {
	var l []string
	for key := range u.Gate {
		if s, ok := key.(string); ok {
			l = append(l, s)
		}
	}
	return l
}

// SortedNumbers returns a sorted list of all gate IDs that are ints
func (u Circuit) SortedNumbers() []int {
	x := u.Numbers()
	sort.Ints(x)
	return x
}

// Numbers returns a list of all gate IDs that are ints
func (u Circuit) Numbers() []int {
	var l []int
	for key := range u.Gate {
		if i, ok := key.(int); ok {
			l = append(l, i)
		}
	}
	return l
}

// Names returns a list of all gate IDs (whether they are string or int)
func (u Circuit) Names() []Name {
	var r []Name
	for n := range u.Gate {
		r = append(r, n)
	}
	return r
}

// SortedNames returns a sorted list of all gate IDs (whether they are string or int)
func (u Circuit) SortedNames() []Name {
	n := u.Names()
	SortNames(n)
	return n
}

// Gates returns a map of gate IDs to their values
func (u Circuit) Gates() map[Name]Value {
	return u.Gate
}

// Unify creates a most simple string to narrowly identify a circuit
// by its (supplied) name and number of gates
func (u Circuit) Unify(name string) string {
	return fmt.Sprintf("%s#%d", name, u.Len())
}
