// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package text provides gates for manipulating text.
package text

import (
	"bytes"
	"io"
	// "log"
	"sync"

	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	ns := faculty.Root.Refine("text")
	ns.AddTerminal("Merge", Merge{})
}

// Merge …
type Merge struct{}

func (Merge) Materialize() think.Reflex {
	_Endo, _Exo := think.NewSynapse()
	firstEndo, firstExo := think.NewSynapse()
	secondEndo, secondExo := think.NewSynapse()
	go func() {
		h := &merge{
			ready: make(chan struct{}),
		}
		h.reply = _Endo.Focus(think.DontCognize)
		close(h.ready)
		firstEndo.Focus(func(v interface{}) { h.CognizeArm(0, v) })
		secondEndo.Focus(func(v interface{}) { h.CognizeArm(1, v) })
	}()
	return think.Reflex{
		"_": _Exo, 
		"First": firstExo, 
		"Second": secondExo, 
	}
}

type merge struct {
	ready chan struct{}
	reply *think.ReCognizer
	sync.Mutex
	arm [2]bytes.Buffer
}

func (h *merge) CognizeArm(index int, v interface{}) {
	<-h.ready
	h.Lock()
	defer h.Unlock()
	h.arm[index].Reset()
	switch t := v.(type) {
	case string:
		h.arm[index].WriteString(t)
	case []byte:
		h.arm[index].Write(t)
	case byte:
		h.arm[index].WriteByte(t)
	case rune:
		h.arm[index].WriteRune(t)
	case io.Reader:
		io.Copy(&h.arm[index], t)
	default:
		panic("unsupported")
	}
	// merge
	var a bytes.Buffer
	a.Write(h.arm[0].Bytes())
	a.Write(h.arm[1].Bytes())
	h.reply.ReCognize(a)
}
