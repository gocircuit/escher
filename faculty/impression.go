// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

import (
	"sort"
	. "github.com/gocircuit/escher/image"
)

// Impression binds time, language and meaning into one.
type Impression struct {
	Image
}

// MakeImpression returns a new empty imptence.
func MakeImpression() Impression {
	return Impression{Make()}
}

func (imp Impression) Show(t int, valve string, value interface{}) Impression {
	f := Make().Grow("Index", t).Grow("Valve", valve).Grow("Value", value)
	imp.Image.Abandon(valve).Grow(valve, f)
	return imp
}

// At returns the functional stored in this impression at the given valve.
func (imp Impression) Valve(valve string) Functional {
	return Functional{imp.Image[valve].(Image)}
}

func (imp Impression) Order() []Functional {
	ff := make([]Functional, 0, imp.Image.Len())
	for _, f := range imp.Image {
		ff = append(ff, Functional{f.(Image)})
	}
	sort.Sort(fading(ff))
	return ff
}

// Functional…
type Functional struct {
	Image
}

func (f Functional) Valve() string {
	return f.Image["Valve"].(string)
}

func (f Functional) Value() interface{} {
	return f.Image["Value"]
}

func (f Functional) Index() int {
	return f.Image["Index"].(int)
}

// fading sorts the functionals in ascending index
type fading []Functional

func (x fading) Len() int {
	return len(x)
}

func (x fading) Less(i, j int) bool {
	return x[i].Index() > x[j].Index()
}

func (x fading) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
