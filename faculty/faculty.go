// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

import (
	// "fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	. "github.com/gocircuit/escher/union"
	"github.com/gocircuit/escher/see"
)

// I see forward. I think back. I see that I think. I think that I see. Thinking and seeing are not apart.

// Faculty is a node in a hierarchy of nodes that can hold subnodes as well as circuit designs (themselves union structures).
type Faculty Union

func NewFaculty() Faculty {
	fty := Faculty(New())
	Union(fty).Add(Genus_{}, NewGenus())
	return fty
}

func (fty Faculty) Genus() Genus {
	g, _ := Union(fty).At(Genus_{})
	return g.(Genus)
}

func (fty Faculty) Forget(name Name) (forgotten Meaning) {
	return Union(fty).Forget(name)
}

// Roam traverses the hierarchy, creating faculty nodes if necessary, returning the final two nodes.
func (fty Faculty) Roam(walk ...Name) (parent, child Meaning) {
	if len(walk) == 0 {
		return nil, fty
	}
	if parent, child = fty.Walk(walk[0]); parent == nil && child == nil { // If no child, make it
		child = fty.Refine(walk[0])
	}
	fac, ok := child.(Faculty)
	if !ok {
		panic("walking thru a non-faculty")
	}
	return fac.Roam(walk[1:]...)
}

// Walk ...
func (fty Faculty) Walk(walk ...Name) (parent, child Meaning) {
	if len(walk) == 0 {
		return nil, fty
	}

	v, ok := Union(fty).At(walk[0])
	if !ok {
		return nil, nil
	}
	switch t := v.(type) {
	case Faculty:
		if len(walk) == 1 {
			return fty, t
		}
		return t.Walk(walk[1:]...)
	default: // non-faculty children are leaves (e.g. Union, Circuit, Gate)
		if len(walk) != 1 {
			panic("walk terminated")
		}
		return fty, t
	}
	panic(7)
}

func (fty Faculty) Refine(name Name) Faculty {
	if x, ok := Union(fty).At(name); ok {
		return x.(Faculty)
	}
	y := NewFaculty()
	y.Genus().SetWalk(append(fty.Genus().GetWalk(), name))
	Union(fty).Add(name, y)
	return y
}

func (fty Faculty) AddTerminal(name Name, term Meaning) {
	if _, over := Union(fty).Add(name, term); over {
		panic(7)
	}
}

// UnderstandDirectory ...
func (fty Faculty) UnderstandDirectory(acid, dir string) {
	d, err := os.Open(dir)
	if err != nil {
		log.Fatalln(err)
	}
	defer d.Close()
	//
	fty.Genus().AddAcid(acid, dir)
	fileInfos, err := d.Readdir(0)
	if err != nil {
		log.Fatalln(err)
	}
	for _, fileInfo := range fileInfos {
		filePath := path.Join(dir, fileInfo.Name())
		if fileInfo.IsDir() {
			fty.Refine(fileInfo.Name()).UnderstandDirectory(acid, filePath)
			continue
		}
		if path.Ext(fileInfo.Name()) != ".escher" {
			continue
		}
		fty.UnderstandFile(dir, filePath)
	}
}

// UnderstandFile ...
func (fty Faculty) UnderstandFile(dir, filePath string) {
	text, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Problem reading source file %s (%v)", filePath, err)
	}
	src := see.NewSrcString(string(text))
	for {
		n_, u_ := see.SeePeer(src)
		if n_ == nil {
			break
		}
		n := n_.(see.Address).Simple() // n is a string
		u := u_.(Union)
		sanitize(n, u)

		t := Understand(s)
		t.sourceDir, t.sourceFile = dir, filePath
		fty.Interpret(t)
	}
}

func sanitize(name Name, u Union) {
	for nm, y := range u.Symbols() {
		if y == nil && nm != name {
			log.Fatalf("implicit non-super peer")
		}
	}
}

func (fty Faculty) Interpret(cir *Circuit) (fresh *Circuit) {
	w, ok := fty[cir.Name()]
	if !ok {
		fty[cir.Name()] = cir
		return cir
	}
	if wcir, ok := w.(*Circuit); ok {
		wcir.Merge(cir)
		return wcir
	}
	// otherwise overwrite existing design
	fty[cir.Name()] = cir
	return cir
}
