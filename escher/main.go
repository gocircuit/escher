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
	"github.com/gocircuit/escher/kit/shell"
	"github.com/gocircuit/escher/see"

	// Load faculties
	"github.com/gocircuit/escher/faculty/circuit"
	_os "github.com/gocircuit/escher/faculty/os"
	_ "github.com/gocircuit/escher/faculty/basic"
	_ "github.com/gocircuit/escher/faculty/escher"
	_ "github.com/gocircuit/escher/faculty/io"
	_ "github.com/gocircuit/escher/faculty/path"
	_ "github.com/gocircuit/escher/faculty/text"
	_ "github.com/gocircuit/escher/faculty/model"
	_ "github.com/gocircuit/escher/faculty/time"
)

// usage: escher [-a dir] [-show] address arguments...
var (
	flagShow     = flag.Bool("show", false, "print only")
	flagSrc        = flag.String("src", "", "source directory")
	flagDiscover = flag.String("d", "", "multicast UDP discovery address for gocircuit.org faculty")
)

func main() {
	// parse flags
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %v [-src Dir] [-show] [-d NetAddress] MainCircuit Arguments...\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	var flagMain string
	var flagArgs = flag.Args()
	if len(flagArgs) > 0 {
		flagMain, flagArgs = flagArgs[0], flagArgs[1:]
	}

	// initialize faculties
	_os.Init(*flagSrc, flagArgs)
	circuit.Init(*flagDiscover)
	//
	switch {

	case *flagShow:
		cd := compile(*flagSrc).Lookup(see.ParseAddress(flagMain))
		switch t := cd.(type) {
		case Circuit:
			fmt.Println(t.Print("", "\t", -1))
		// case Faculty:
		default:
			fmt.Printf("%T/%v\n", t, t)
		}

	default:
		mem := compile(*flagSrc)
		b := NewRenderer(Memory(mem))
		defer func() {
			if r := recover(); r != nil {
				debug.PrintStack()
				log.Printf("Recovered: %v\n", r)
				shell.NewShell("(recovered)", os.Stdin, os.Stdout, os.Stderr).Loop(Circuit(mem))
			}
		}()
		b.MaterializeAddress(see.ParseAddress(flagMain))
		select {} // wait forever
	}
}

func compile(dir string) Memory {
	if dir != "" {
		Load(Root(), dir)
	}
	return Root()
}
