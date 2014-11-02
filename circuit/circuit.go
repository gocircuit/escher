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

// Circuit ...
type Circuit struct {
	Gate map[Name]Value
	Flow map[Name]map[Name]Vector // gate -> valve -> opposing gate and valve
}

func New() Circuit {
	return Circuit{
		Gate: make(map[Name]Value),
		Flow: make(map[Name]map[Name]Vector),
	}
}

func (u Circuit) IsNil() bool {
	return u.Gate == nil || u.Flow == nil
}

func (u Circuit) IsEmpty() bool {
	return len(u.Gate) == 0 && len(u.Flow) == 0
}

func (u Circuit) SortedLetters() []string {
	x := u.Letters()
	sort.Strings(x)
	return x
}

func (u Circuit) Letters() []string {
	var l []string
	for key, _ := range u.Gate {
		if s, ok := key.(string); ok {
			l = append(l, s)
		}
	}
	return l
}

func (u Circuit) SortedNumbers() []int {
	x := u.Numbers()
	sort.Ints(x)
	return x
}

func (u Circuit) Numbers() []int {
	var l []int
	for key, _ := range u.Gate {
		if i, ok := key.(int); ok {
			l = append(l, i)
		}
	}
	return l
}

func (u Circuit) Names() []Name {
	var r []Name
	for n, _ := range u.Gate {
		r = append(r, n)
	}
	return r
}

func (u Circuit) SortedNames() []Name {
	n := u.Names()
	SortNames(n)
	return n
}

func (u Circuit) Gates() map[Name]Value {
	return u.Gate
}

func (u Circuit) String() string {
	return u.Print("", "\t", -1)
}
