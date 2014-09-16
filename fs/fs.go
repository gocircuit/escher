// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package fs provides routines for reading in Escher circuits from directory structures.
package fs

import (
	// "fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/see"
)

func Load(into Circuit, acid, filedir string) {
	fi, err := os.Stat(filedir)
	if err != nil {
		log.Fatalf("cannot read source file %s (%v)", filedir, err)
	}
	if fi.IsDir() {
		loadDirectory(into, acid, filedir)
	} else {
		loadFile(into, acid, filedir)
	}
}

// loadDirectory ...
func loadDirectory(into Circuit, acid, dir string) Circuit {
	d, err := os.Open(dir)
	if err != nil {
		log.Fatalln(err)
	}
	defer d.Close()
	//
	addSource(into, acid, dir)
	fileInfos, err := d.Readdir(0)
	if err != nil {
		log.Fatalln(err)
	}
	for _, fileInfo := range fileInfos {
		filePath := path.Join(dir, fileInfo.Name())
		if fileInfo.IsDir() {
			into.Grow(fileInfo.Name(), loadDirectory(New(), acid, filePath))
			continue
		}
		if path.Ext(fileInfo.Name()) != ".escher" {
			continue
		}
		loadFile(into, dir, filePath)
	}
	return into
}

func addSource(into Circuit, acid, dir string) {
	src, ok := into.CircuitOptionAt(Source{})
	if !ok {
		src = New()
		into.Include(Source{}, src)
	}
	src.Include(src.Len(), New().Grow("Acid", acid).Grow("Dir", dir))
}

// loadFile ...
func loadFile(into Circuit, dir, file string) Circuit {
	text, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Problem reading source file %s (%v)", file, err)
	}
	src := see.NewSrcString(string(text))
	for {
		n_, u_ := see.SeePeer(src)
		if n_ == nil {
			break
		}
		n := n_.(string) // n is a string
		u := u_.(Circuit)
		u.Include(Source{}, New().Grow("Dir", dir).Grow("File", file))
		into.Include(n, u)
	}
	return into
}

// Source is a name type for a genus structure.
type Source struct{}

func (Source) String() string {
	return "*Source"
}
