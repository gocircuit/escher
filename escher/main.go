// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	. "github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/circuit"
	. "github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/kit/fs"
	kio "github.com/gocircuit/escher/kit/io"
	"github.com/gocircuit/escher/see"

	// Load faculties
	"github.com/gocircuit/escher/faculty/circuit"
	fos "github.com/gocircuit/escher/faculty/os"
	_ "github.com/gocircuit/escher/faculty/basic"
	_ "github.com/gocircuit/escher/faculty/cmplx"
	_ "github.com/gocircuit/escher/faculty/spin"
	_ "github.com/gocircuit/escher/faculty/escher"
	_ "github.com/gocircuit/escher/faculty/io"
	_ "github.com/gocircuit/escher/faculty/path"
	_ "github.com/gocircuit/escher/faculty/view"
	"github.com/gocircuit/escher/faculty/test"
	_ "github.com/gocircuit/escher/faculty/http"
	_ "github.com/gocircuit/escher/faculty/index"
	_ "github.com/gocircuit/escher/faculty/text"
	_ "github.com/gocircuit/escher/faculty/model"
	_ "github.com/gocircuit/escher/faculty/time"
	_ "github.com/gocircuit/escher/faculty/yield"
)

// usage: escher [-a dir] [-show] address arguments...
var (
	flagSrc        = flag.String("src", "", "source directory")
	flagDiscover = flag.String("d", "", "multicast UDP discovery address for gocircuit.org faculty")
)

func main() {
	// parse flags
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %v [-src Dir] [-d NetAddress] MainCircuit Arguments...\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	var flagMain string
	var flagArgs = flag.Args()
	if len(flagArgs) > 0 {
		flagMain, flagArgs = flagArgs[0], flagArgs[1:]
	}
	// parse env
	if *flagSrc == "" {
		*flagSrc = os.Getenv("ESCHER")
	}

	// initialize faculties
	fos.Init(flagArgs)
	test.Init(*flagSrc)
	circuit.Init(*flagDiscover)
	//
	index := Root()
	if *flagSrc != "" {
		index.Merge(Load(*flagSrc))
	}
	// run main
	if flagMain != "" {
		exec(index, see.ParseAddress(flagMain), false)
	}
	// standard loop
	r := kio.NewChunkReader(os.Stdin)
	for {
		chunk, err := r.Read()
		if err != nil {
			fmt.Fprintf(os.Stderr, "end of session (%v)\n", err)
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				break
			}
		}
		src := see.NewSrcString(string(chunk))
		for src.Len() > 0 {
			u := see.SeeChamber(src)
			if u == nil || u.(Circuit).Len() == 0 {
				break
			}
			fmt.Fprintf(os.Stderr, "MATERIALIZING %v\n", u)
			exec(index, u, true)
		}
	}
}

func exec(index Index, v Value, showResidue bool) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("execution glitch (%v)", r)
		}
	}()
	residue := Materialize(index, v)
	if showResidue {
		fmt.Fprintf(os.Stderr, "RESIDUE %v\n\n", residue)
	}
}
