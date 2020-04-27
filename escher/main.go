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
	"os"

	"github.com/hoijui/escher/a"
	"github.com/hoijui/escher/be"
	cir "github.com/hoijui/escher/circuit"
	fac "github.com/hoijui/escher/faculty"
	"github.com/hoijui/escher/kit/fs"
	kio "github.com/hoijui/escher/kit/io"
	"github.com/hoijui/escher/see"

	// Load faculties
	_ "github.com/hoijui/escher/faculty/basic"
	"github.com/hoijui/escher/faculty/circuit"
	_ "github.com/hoijui/escher/faculty/cmplx"
	_ "github.com/hoijui/escher/faculty/escher"
	_ "github.com/hoijui/escher/faculty/http"
	_ "github.com/hoijui/escher/faculty/index"
	_ "github.com/hoijui/escher/faculty/io"
	_ "github.com/hoijui/escher/faculty/math"
	_ "github.com/hoijui/escher/faculty/model"
	fos "github.com/hoijui/escher/faculty/os"
	"github.com/hoijui/escher/faculty/test"
	_ "github.com/hoijui/escher/faculty/text"
	_ "github.com/hoijui/escher/faculty/time"
	_ "github.com/hoijui/escher/faculty/yield"
)

// usage: escher [-a dir] [-show] address arguments...
var (
	flagSrc      = flag.String("src", "", "source directory")
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
	index := fac.Root()
	if *flagSrc != "" {
		index.Merge(fs.Load(*flagSrc))
	}
	// run main
	if flagMain != "" {
		verb := see.ParseVerb(flagMain)
		if cir.Circuit(verb).IsNil() {
			fmt.Fprintf(os.Stderr, "verb not recognized\n")
			os.Exit(1)
		}
		exec(index, cir.Circuit(verb), false)
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
		src := a.NewSrcString(string(chunk))
		for src.Len() > 0 {
			u := see.SeeChamber(src)
			if u == nil || u.(cir.Circuit).Len() == 0 {
				break
			}
			fmt.Fprintf(os.Stderr, "MATERIALIZING %v\n", u)
			exec(index, u.(cir.Circuit), true)
		}
	}
}

func exec(index be.Index, verb cir.Circuit, showResidue bool) {
	residue := be.MaterializeSystem(cir.Circuit(verb), cir.Circuit(index), cir.New().Grow("Main", cir.New()))
	if showResidue {
		fmt.Fprintf(os.Stderr, "RESIDUE %v\n\n", residue)
	}
}
