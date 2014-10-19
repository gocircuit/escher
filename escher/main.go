// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package main

import (
	"flag"
	"fmt"
	"log"
	"runtime/debug"
	"os"

	. "github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/circuit"
	. "github.com/gocircuit/escher/kit/memory"
	. "github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/kit/fs"
	"github.com/gocircuit/escher/see"

	// Load faculties
	"github.com/gocircuit/escher/faculty/circuit"
	_os "github.com/gocircuit/escher/faculty/os"
	_ "github.com/gocircuit/escher/faculty/basic"
	_ "github.com/gocircuit/escher/faculty/cmplx"
	_ "github.com/gocircuit/escher/faculty/exp/draw"
	_ "github.com/gocircuit/escher/faculty/escher"
	_ "github.com/gocircuit/escher/faculty/io"
	_ "github.com/gocircuit/escher/faculty/path"
	_ "github.com/gocircuit/escher/faculty/view"
	_ "github.com/gocircuit/escher/faculty/testing"
	_ "github.com/gocircuit/escher/faculty/text"
	_ "github.com/gocircuit/escher/faculty/model"
	_ "github.com/gocircuit/escher/faculty/time"
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
	} else {
		flagMain = "escher.Shell" // escher assembler shell
	}

	// initialize faculties
	_os.Init(*flagSrc, flagArgs)
	circuit.Init(*flagDiscover)
	//
	mem := compile(*flagSrc)
	defer func() {
		if r := recover(); r != nil {
			if flagMain != "" {
				debug.PrintStack()
				log.Printf("Recovered: %v\n", r)
			}
		}
	}()
	Materialize(Circuit(mem), see.ParseAddress(flagMain))
	select {} // wait forever
}

func compile(dir string) Memory {
	if dir != "" {
		Load(Root(), dir)
	}
	return Root()
}
