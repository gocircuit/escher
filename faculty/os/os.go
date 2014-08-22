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
	"strings"

	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/think"
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
	ns.AddTerminal("Stdin", Stdin{})
	ns.AddTerminal("Stdout", Stdout{})
	ns.AddTerminal("Stderr", Stderr{})
}

var args map[string]string

// Arg
type Arg struct{}

func (Arg) Materialize() think.Reflex {
	valueEndo, valueExo := think.NewSynapse()
	nameEndo, nameExo := think.NewSynapse()
	go func() {
		h := &arg{}
		h.valueRe = valueEndo.Focus(think.DontCognize)
		nameEndo.Focus(h.CognizeName)
	}()
	return think.Reflex{
		"Name":  nameExo,
		"Value": valueExo,
	}
}

type arg struct {
	valueRe *think.ReCognizer
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

func (Env) Materialize() think.Reflex {
	valueEndo, valueExo := think.NewSynapse()
	nameEndo, nameExo := think.NewSynapse()
	go func() {
		h := &env{}
		h.valueRe = valueEndo.Focus(think.DontCognize)
		nameEndo.Focus(h.CognizeName)
	}()
	return think.Reflex{
		"Name":  nameExo,
		"Value": valueExo,
	}
}

type env struct {
	valueRe *think.ReCognizer
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

func (Exit) Materialize() think.Reflex {
	_Endo, _Exo := think.NewSynapse()
	go func() {
		_Endo.Focus(cognizeExit)
	}()
	return think.Reflex{
		"_": _Exo,
	}
}

func cognizeExit(v interface{}) {
	log.Printf("Exit %v", v)
	os.Exit(v.(int))
}

// Stdin
type Stdin struct{}

func (Stdin) Materialize() think.Reflex {
	return think.NewNounReflex(os.Stdin)
}

// Stdout
type Stdout struct{}

func (Stdout) Materialize() think.Reflex {
	return think.NewNounReflex(os.Stdout)
}

// Stderr
type Stderr struct{}

func (Stderr) Materialize() think.Reflex {
	return think.NewNounReflex(os.Stderr)
}
