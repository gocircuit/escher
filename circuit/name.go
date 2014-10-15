// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	// "log"
	"sort"
)

// Name represents a gate name within a circuit.
type Name interface{}

// SortNames sorts the argument slice of names, prioritizing strings before ints and everything else after.
func SortNames(names []Name) {
	sort.Sort(sortNames(names))
}

type sortNames []Name

func (s sortNames) Len() int {
	return len(s)
}

// nil, string, int, ...
func (s sortNames) Less(i, j int) bool {
	switch ti := s[i].(type) {
	case int:
		switch tj := s[j].(type) {
		case int:
			return ti < tj
		case string:
			return false
		}
		return false
	case string:
		switch tj := s[j].(type) {
		case int:
			return true
		case string:
			return ti < tj
		}
		return false
	case nil:
		return true
	}
	return false
}

func (s sortNames) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
