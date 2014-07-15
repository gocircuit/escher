// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package faculty

import (
	"github.com/gocircuit/escher/tree"
)

// Time—>SentenceFunctional
// Time zero is most recent, time one is earlier, and so on.
type Sentence tree.Tree

func (s Sentence) Grow(valve string, value tree.Tree) Sentence {
	return Sentence(tree.Tree(s).Grow(valve, value))
}

func (s Sentence) At(valve string) SentenceFunctional {
	return tree.Tree(s).At(valve).(SentenceFunctional)
}

// "Valve"—>string, "Value"—>interface{}, "Time"—>int
type SentenceFunctional tree.Tree

func (sf SentenceFunctional) Valve() string {
	return tree.Tree(sf).String("Valve")
}

func (sf SentenceFunctional) Value() interface{} {
	return tree.Tree(sf).At("Value")
}

func (sf SentenceFunctional) Time() int {
	return tree.Tree(sf).Int("Time")
}

// Valve—>MemoryFunctional
type Memory tree.Tree

func (m Memory) Grow(valve string, value tree.Tree) Memory {
	return Memory(tree.Tree(m).Grow(valve, value))
}

func (m Memory) At(valve string) MemoryFunctional {
	return tree.Tree(m).At(valve).(MemoryFunctional)
}

// "Valve"—>string, "Value"—>interface{}, "Age"—>int
type MemoryFunctional tree.Tree

func (mf MemoryFunctional) Valve() string {
	return tree.Tree(sf).String("Valve")
}

func (mf MemoryFunctional) Value() interface{} {
	return tree.Tree(sf).At("Value")
}

func (mf MemoryFunctional) Age() int {
	return tree.Tree(sf).Int("Age")
}

// ShortCognize is the cognition interface provided by the Mind's Eye (short-term memory) mechanism.
// The short-term memory is what allows people to process a linguistic sentence with all its structure.
type ShortCognize func(Sentence)
