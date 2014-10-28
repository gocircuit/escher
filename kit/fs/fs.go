// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package fs provides routines for reading Escher circuits from source directories and files.
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

func Load(filedir string) Circuit {
	fi, err := os.Stat(filedir)
	if err != nil {
		log.Fatalf("cannot read source file %s (%v)", filedir, err)
	}
	if fi.IsDir() {
		return loadDirectory(filedir)
	} 
	return loadFile("", filedir)
}

// loadDirectory ...
func loadDirectory(dir string) Circuit {
	d, err := os.Open(dir)
	if err != nil {
		log.Fatalln(err)
	}
	defer d.Close()
	//
	x := New()
	x.Include(Source{}, New().Grow("Dir", dir))
	//
	fileInfos, err := d.Readdir(0)
	if err != nil {
		log.Fatalln(err)
	}
	for _, fileInfo := range fileInfos {
		filePath := path.Join(dir, fileInfo.Name())
		if fileInfo.IsDir() { // directory
			x.Include(fileInfo.Name(), loadDirectory(filePath))
			continue
		}
		if path.Ext(fileInfo.Name()) != ".escher" { // file
			continue
		}
		x.Merge(loadFile(dir, filePath))
	}
	return x
}

// loadFile ...
func loadFile(dir, file string) Circuit {
	text, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Problem reading source file %s (%v)", file, err)
	}
	x := New()
	src := see.NewSrcString(string(text))
	for {
		n_, u_ := see.SeePeer(src)
		if n_ == nil {
			break
		}
		n := n_.(string) // n is a string
		if u, ok := u_.(Circuit); ok {
			u.Include(Source{}, New().Grow("Dir", dir).Grow("File", file))
		}
		x.Include(n, u_)
	}
	return x
}

// Source is a name type for a genus structure.
type Source struct{}

func (Source) String() string {
	return "(SourceInfo)"
}
