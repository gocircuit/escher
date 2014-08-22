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

	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/think"
)

func init() {
	ns := faculty.Root.Refine("text")
	ns.AddTerminal("Merge", Merge{})
	ns.AddTerminal("Form", Form{})
}

// Merge …
type Merge struct{}

func (Merge) Materialize() think.Reflex {
	_Endo, _Exo := think.NewSynapse()
	firstEndo, firstExo := think.NewSynapse()
	secondEndo, secondExo := think.NewSynapse()
	thirdEndo, thirdExo := think.NewSynapse()
	go func() {
		h := &merge{
			ready: make(chan struct{}),
		}
		h.reply = _Endo.Focus(think.DontCognize)
		close(h.ready)
		firstEndo.Focus(func(v interface{}) { h.CognizeArm(0, v) })
		secondEndo.Focus(func(v interface{}) { h.CognizeArm(1, v) })
		thirdEndo.Focus(func(v interface{}) { h.CognizeArm(2, v) })
	}()
	return think.Reflex{
		"_":      _Exo,
		"First":  firstExo,
		"Second": secondExo,
		"Third":  thirdExo,
	}
}

type merge struct {
	ready chan struct{}
	reply *think.ReCognizer
	sync.Mutex
	arm [3]*bytes.Buffer
}

func (h *merge) CognizeArm(index int, v interface{}) {
	<-h.ready
	h.Lock()
	defer h.Unlock()
	switch t := v.(type) {
	case string:
		h.arm[index] = bytes.NewBufferString(t)
	case []byte:
		h.arm[index] = bytes.NewBuffer(t)
	case byte:
		h.arm[index] = bytes.NewBuffer([]byte{t})
	case rune:
		h.arm[index] = bytes.NewBuffer(nil)
		h.arm[index].WriteRune(t)
	case io.Reader:
		h.arm[index] = bytes.NewBuffer(nil)
		io.Copy(h.arm[index], t)
	default:
		panic("unsupported")
	}
	// merge
	if h.arm[0] == nil || h.arm[1] == nil || h.arm[2] == nil {
		return
	}
	var a bytes.Buffer
	a.Write(h.arm[0].Bytes())
	a.Write(h.arm[1].Bytes())
	a.Write(h.arm[2].Bytes())
	h.reply.ReCognize(a.String())
}
