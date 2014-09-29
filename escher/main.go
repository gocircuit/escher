// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package main

import (
	"flag"
	"fmt"
	"strings"
	"os"

	. "github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/circuit"
	. "github.com/gocircuit/escher/memory"
	. "github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/fs"
	"github.com/gocircuit/escher/shell"

	// Load faculties
	"github.com/gocircuit/escher/faculty/acid"
	"github.com/gocircuit/escher/faculty/circuit"
	facos "github.com/gocircuit/escher/faculty/os"
	
	_ "github.com/gocircuit/escher/faculty/basic"
	_ "github.com/gocircuit/escher/faculty/escher"
	// _ "github.com/gocircuit/escher/faculty/handbook"
	_ "github.com/gocircuit/escher/faculty/io"
	_ "github.com/gocircuit/escher/faculty/io/util"
	_ "github.com/gocircuit/escher/faculty/path"
	_ "github.com/gocircuit/escher/faculty/text"
	_ "github.com/gocircuit/escher/faculty/model"
	// _ "github.com/gocircuit/escher/faculty/think"
	_ "github.com/gocircuit/escher/faculty/time"
	// _ "github.com/gocircuit/escher/faculty/web/twitter"
	// _ "github.com/gocircuit/escher/faculty/xml"
)

var (
	flagMain     = flag.String("main", "main", "address of the startup circuit")
	flagShow     = flag.String("show", "", "print out an object at a given path; don't run")
	flagSvg     = flag.String("svg", "", "display a circuit as SVG; don't run")
	flagX        = flag.String("x", "", "program source directory X")
	flagY        = flag.String("y", "", "program source directory Y")
	flagZ        = flag.String("z", "", "program source directory Z")
	flagName     = flag.String("n", "", "execution name")
	flagArg      = flag.String("a", "", "program arguments")
	flagDiscover = flag.String("d", "", "multicast UDP discovery address for circuit faculty, if needed")
)

func main() {
	flag.Parse()
	// Initialize faculties
	facos.Init(*flagArg)
	loadCircuitFaculty(*flagName, *flagDiscover, *flagX, *flagY, *flagZ)
	//
	switch {
	case *flagSvg != "":
		walk := strings.Split(*flagSvg, ".")
		if len(walk) == 2 && walk[0] == "" && walk[1] == "" { // -svg .
			walk = nil
		}
		cd := compile(*flagX, *flagY, *flagZ).Lookup(NewAddressStrings(walk))
		switch t := cd.(type) {
		case Circuit:
			println("drawing not supported")
		// case Faulty:
		default:
			println(fmt.Sprintf("SVG display available only for circuits (%T)", t))
		}

	case *flagShow != "":
		walk := strings.Split(*flagShow, ".")
		if len(walk) == 2 && walk[0] == "" && walk[1] == "" { // -show .
			walk = nil
		}
		cd := compile(*flagX, *flagY, *flagZ).Lookup(NewAddressStrings(walk))
		switch t := cd.(type) {
		case Circuit:
			fmt.Println(t.Print("", "\t"))
		// case Faculty:
		default:
			fmt.Printf("%T/%v\n", t, t)
		}

	default:
		mem := compile(*flagX, *flagY, *flagZ)
		b := NewRenderer(mem)
		defer func() {
			if r := recover(); r != nil {
				panic(r)
				shell.NewShell(os.Stdin, os.Stdout, os.Stderr, mem).Loop()
			}
		}()
		b.MaterializeAddress(NewAddressParse(*flagMain))
		select {} // wait forever
	}
}

func compile(x, y, z string) Memory {
	if x != "" {
		Load(Root(), "X", x)
	}
	if y != "" {
		Load(Root(), "Y", y)
	}
	if z != "" {
		Load(Root(), "Z", z)
	}
	return Root()
}

func loadCircuitFaculty(name, discover, x, y, z string) {
	acid.Init(x, y, z)
	circuit.Init(discover)
}
