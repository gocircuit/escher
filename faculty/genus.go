// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

import (
	. "github.com/gocircuit/escher/union"
)

// Genus_ is a name type used to mark a genus structure within a union.
type Genus_ struct{}
type Walk_ struct{}

//
type Genus Union

func NewGenus() Genus {
	return Genus(New())
}

func (g Genus) SetWalk(walk []Name) {
	Union(g).Add(Walk_{}, walk)
}

func (g Genus) GetWalk() []Name {
	w, ok := Union(g).At(Walk_{})
	if !ok {
		return nil
	}
	return w.([]Name)
}

func (g Genus) AddAcid(acid, dir string) {
	Union(g).At(acid, dir)
}
