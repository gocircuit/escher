// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package os

import (
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/kit/plumb"
)

func Init(a string) {
	args = make(map[string]string) // n1=v1,n2=v2
	for _, p := range strings.Split(a, ",") {
		if p == "" {
			continue
		}
		nv := strings.Split(p, "=")
		if len(nv) != 2 {
			panic("command-line argument syntax")
		}
		v, err := url.QueryUnescape(nv[1])
		if err != nil {
			panic(err)
		}
		args[nv[0]] = v
		log.Printf("Argument %s=%s", nv[0], v)
	}
	ns := faculty.Root.Refine("os")
	ns.AddTerminal("Arg", Arg{})
	ns.AddTerminal("Env", Env{})
	ns.AddTerminal("Exit", Exit{})
	ns.AddTerminal("Fatal", Fatal{})
	ns.AddTerminal("Stdin", Stdin{})
	ns.AddTerminal("Stdout", Stdout{})
	ns.AddTerminal("Stderr", Stderr{})
	//
	ns.AddTerminal("LookPath", LookPath{})
	ns.AddTerminal("Process", Process{})
	ns.AddTerminal("ForkCommand", ForkCommand{})
	ns.AddTerminal("ForkExit", ForkExit{})
	ns.AddTerminal("ForkIO", ForkIO{})
	//
}

var args map[string]string

// Arg
type Arg struct{}

func (Arg) Materialize(*be.Matter) be.Reflex {
	valueEndo, valueExo := be.NewSynapse()
	nameEndo, nameExo := be.NewSynapse()
	go func() {
		h := &arg{}
		h.valueRe = valueEndo.Focus(be.DontCognize)
		nameEndo.Focus(h.CognizeName)
	}()
	return be.Reflex{
		"Name":  nameExo,
		"Value": valueExo,
	}
}

type arg struct {
	valueRe *be.ReCognizer
}

func (h *arg) CognizeName(v interface{}) {
	n, ok := v.(string)
	if !ok {
		panic("non-string name perceived by os.arg")
	}
	h.valueRe.ReCognize(args[n])
}

// Env
type Env struct{}

func (Env) Materialize(*be.Matter) be.Reflex {
	valueEndo, valueExo := be.NewSynapse()
	nameEndo, nameExo := be.NewSynapse()
	go func() {
		h := &env{}
		h.valueRe = valueEndo.Focus(be.DontCognize)
		nameEndo.Focus(h.CognizeName)
	}()
	return be.Reflex{
		"Name":  nameExo,
		"Value": valueExo,
	}
}

type env struct {
	valueRe *be.ReCognizer
}

func (h *env) CognizeName(v interface{}) {
	n, ok := v.(string)
	if !ok {
		panic("non-string name perceived by os.env")
	}
	ev := os.Getenv(n)
	log.Printf("Environment %s=%s", n, ev)
	h.valueRe.ReCognize(ev)
}

// Exit
type Exit struct{}

func (Exit) Materialize(*be.Matter) be.Reflex {
	reflex, _ := plumb.NewEyeCognizer(
		func(eye *plumb.Eye, valve string, value interface{}) {
			os.Exit(value.(int))
		}, 
		"_",
	)
	return reflex
}

// Fatal
type Fatal struct{}

func (Fatal) Materialize(*be.Matter) be.Reflex {
	reflex, _ := plumb.NewEyeCognizer(
		func(eye *plumb.Eye, valve string, value interface{}) {
			log.Fatalln(value)
		}, 
		"_",
	)
	return reflex
}

// LookPath
type LookPath struct{}

func (LookPath) Materialize(*be.Matter) be.Reflex {
	reflex, _ := plumb.NewEyeCognizer(
		func(eye *plumb.Eye, valve string, value interface{}) {
			if valve != "Name" {
				return
			}
			p, err := exec.LookPath(value.(string))
			if err != nil {
				log.Fatalf("no file path to %s", value.(string))
			}
			eye.Show("_", p)
		},
		"Name", "_",
	)
	return reflex
}
