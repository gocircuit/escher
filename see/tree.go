// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package see

import (
	// "fmt"
	"github.com/gocircuit/escher/tree"
)

func SeeTree(src *Src) (rec TreeDesign, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			rec, ok = nil, false
		}
	}()
	// println(fmt.Sprintf("R<=[%s]", src.String()))
	rec = TreeDesign(tree.Make())
	t := src.Copy()
	t.Match("{")
	Space(t)
	for {
		q := t.Copy()
		Space(q)
		name, scope, ok := SeeField(q)
		if !ok {
			break
		}
		Space(q)
		q.TryMatch(",")
		Space(q)
		for _, w := range scope {
			(tree.Tree)(rec).Grow(name, w)
		}
		t.Become(q)
	}
	Space(t)
	t.TryMatch(",")
	Space(t)
	t.Match("}")
	src.Become(t)
	return rec, true
}
