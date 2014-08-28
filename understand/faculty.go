// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package understand

import (
	// "fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/gocircuit/escher/see"
)

// I see forward. I think back. I see that I be. I think that I see. Thinking and seeing are not apart.

type Faculty map[string]interface{} // name -> {Faculty, *Circuit, etc}

func NewFaculty() Faculty {
	return make(Faculty)
}

func (fty Faculty) Forget(name string) (forgotten interface{}) {
	forgotten = fty[name]
	delete(fty, name)
	return
}

func (fty Faculty) Roam(walk ...string) (parent, child interface{}) {
	if len(walk) == 0 {
		return nil, fty
	}
	if parent, child = fty.Walk(walk[0]); parent == nil && child == nil { // If no child, make it
		child = fty.Refine(walk[0])
	}
	fac, ok := child.(Faculty)
	if !ok {
		panic("overwriting")
	}
	return fac.Roam(walk[1:]...)
}

func (fty Faculty) Walk(walk ...string) (parent, child interface{}) {
	if len(walk) == 0 {
		return nil, fty
	}
	v, ok := fty[walk[0]]
	if !ok {
		return nil, nil
	}
	switch t := v.(type) {
	case Faculty:
		if len(walk) == 1 {
			return fty, t
		}
		return t.Walk(walk[1:]...)
	default: // non-faculty children are leaves (e.g. *Circuit, Gate)
		if len(walk) != 1 {
			panic("walk terminated")
		}
		return fty, t
	}
	panic(7)
}

func (fty Faculty) Refine(name string) (child Faculty) {
	if x, ok := fty[name]; ok {
		return x.(Faculty)
	}
	child = NewFaculty()
	fty[name] = child
	return
}

func (fty Faculty) AddTerminal(name string, term interface{}) {
	if _, ok := fty[name]; ok {
		panic(7)
	}
	fty[name] = term
}

func (fty Faculty) UnderstandDirectory(dir string) {
	d, err := os.Open(dir)
	if err != nil {
		log.Fatalln(err)
	}
	defer d.Close()
	fileInfos, err := d.Readdir(0)
	if err != nil {
		log.Fatalln(err)
	}
	for _, fileInfo := range fileInfos {
		filePath := path.Join(dir, fileInfo.Name())
		if fileInfo.IsDir() {
			fty.Refine(fileInfo.Name()).UnderstandDirectory(filePath)
			continue
		}
		if path.Ext(fileInfo.Name()) != ".escher" {
			continue
		}
		fty.UnderstandFile(filePath)
	}
}

func (fty Faculty) UnderstandFile(file string) {
	text, err := ioutil.ReadFile(file)
	if err != nil {
		println(file)
		panic(err)
	}
	src := see.NewSrcString(string(text))
	for {
		s := see.SeeCircuit(src)
		if s == nil {
			break
		}
		// println(s.Print("", "\t"))
		fty.Interpret(Understand(s))
	}
}

func (fty Faculty) Interpret(cir *Circuit) (fresh *Circuit) {
	w, ok := fty[cir.Name]
	if !ok {
		fty[cir.Name] = cir
		return cir
	}
	if wcir, ok := w.(*Circuit); ok {
		wcir.Merge(cir)
		return wcir
	}
	// otherwise overwrite existing design
	fty[cir.Name] = cir
	return cir
}
