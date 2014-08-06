// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

import (
	"sort"
	"github.com/gocircuit/escher/star"
)

// Impression binds time, language and meaning into one.
type Impression struct {
	*star.Star
}

// MakeImpression returns a new empty imptence.
func MakeImpression() Impression {
	return Impression{ star.Make() }
}

func (imp Impression) Show(t int, valve string, value ??) Impression {
	f := star.Make().
		Grow("Index", t).
		Grow("Valve", valve).
		Grow("Value", value)
	imp.Unwrap().Split(valve).Merge(valve, f)
	return imp
}

// Unwrap returns the star underlying this imptence.
func (imp Impression) Unwrap() *star.Star {
	return (*star.Star)(imp)
}

// At returns the functional stored in this impression at the given valve.
func (imp Impression) Valve(valve string) *Functional {
	return (*Functional)(imp.Unwrap().Down(valve))
}

func (imp Impression) Order() []*Functional {
	ff := make([]*Functional, 0, imp.Unwrap().Len())
	for n, f := range imp.Unwrap().Choice() {
		if n == star.Parent {
			continue
		}
		ff = append(ff, (*Functional)(f))
	}
	sort.Sort(fading(ff))
	return ff
}

// Functional…
type Functional struct {
	*star.Star
}

func (f Functional) Unwrap() *star.Star {
	return f.Star
}

func (f Functional) Valve() string {
	return f.Unwrap().Down("Valve").String()
}

func (f Functional) Value() interface{} {
	return f.Unwrap().Down("Value").Interface()
}

func (f Functional) Index() int {
	return f.Unwrap().Down("Index").Int()
}

// fading sorts the functionals in ascending index
type fading []*Functional

func (x fading) Len() int {
	return len(x)
}

func (x fading) Less(i, j int) bool {
	return x[i].Index() > x[j].Index()
}

func (x fading) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
