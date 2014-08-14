// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package os

import (
	"net/url"
	"os"
	"strings"

	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/faculty"
)

func Init(a string) {
	args = make(map[string]string)  // n1=v1:n2=v2
	for _, p := range strings.Split(a, ":") {
		nv := strings.Split(p, "=")
		if len(nv) != 2 {
			panic("command-line argument syntax")
		}
		v, err := url.QueryUnescape(nv[1])
		if err != nil {
			panic(err)
		}
		args[nv[0]] = v
	}
	faculty.Root.AddTerminal("arg", Arg{})
	faculty.Root.AddTerminal("env", Env{})
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
		"Name": nameExo,
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
		"Name": nameExo,
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
	h.valueRe.ReCognize(os.Getenv(n))
}
