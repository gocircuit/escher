// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package understand

import (
	// "fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/gocircuit/escher/see"
)

// I see forward. I think back. I see that I think. I think that I see. Thinking and seeing are not apart.

type Faculty map[string]interface{} // name -> {Faculty, *Circuit, etc}

func NewFaculty() Faculty {
	return make(Faculty)
}

func (fty Faculty) Walk(walk ...string) (parent, endpoint interface{}) {
	if len(walk) == 0 {
		return nil, fty
	}
	v, ok := fty[walk[0]]
	if !ok {
		panic("walk broken")
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
	if _, ok := fty[name]; ok {
		panic(7)
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
		panic(err)
	}
	defer d.Close()
	fileInfos, err := d.Readdir(0)
	if err != nil {
		panic(err)
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
		fty.interpretCircuit(Understand(s))
	}
}

func (fty Faculty) interpretCircuit(cir *Circuit) {
	w, ok := fty[cir.Name]
	if !ok {
		fty[cir.Name] = cir
		return
	}
	if wcir, ok := w.(*Circuit); ok {
		wcir.Merge(cir)
		return
	}
	// otherwise overwrite existing design
	fty[cir.Name] = cir
}
