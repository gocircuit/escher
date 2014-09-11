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

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/see"
)

func (fty Faculty) LoadDirectoryOrFile(acid, filedir string) {
	fi, err := os.Stat(filedir)
	if err != nil {
		log.Fatalf("cannot read source file %s (%v)", filedir, err)
	}
	if fi.IsDir() {
		fty.LoadDirectory(acid, filedir)
	} else {
		fty.LoadFile(acid, filedir)
	}
}

// LoadDirectory ...
func (fty Faculty) LoadDirectory(acid, dir string) {
	d, err := os.Open(dir)
	if err != nil {
		log.Fatalln(err)
	}
	defer d.Close()
	//
	fty.Genus().Acid[acid] = dir
	fileInfos, err := d.Readdir(0)
	if err != nil {
		log.Fatalln(err)
	}
	for _, fileInfo := range fileInfos {
		filePath := path.Join(dir, fileInfo.Name())
		if fileInfo.IsDir() {
			fty.Refine(fileInfo.Name()).LoadDirectory(acid, filePath)
			continue
		}
		if path.Ext(fileInfo.Name()) != ".escher" {
			continue
		}
		fty.LoadFile(dir, filePath)
	}
}

// LoadFile ...
func (fty Faculty) LoadFile(dir, filePath string) {
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
		n := n_.(string) // n is a string
		u := u_.(Circuit)
		u.Include(Genus_{}, 
			&CircuitGenus{
				Dir: dir,
				File: filePath,
			},
		)
		if _, ok := fty.Include(n, u); ok {
			log.Fatalf("file circuit overwrites %s", n)
		}
	}
}
