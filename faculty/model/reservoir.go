// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package model

import (
	"container/list"
	"sync"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
)

type Reservoir struct{
	sync.Mutex
	stop sync.Once
	y, focus Circuit // output and current focus
	path list.List
}

func (r *Reservoir) Is() {
	r.y = New()
	r.focus = r.y
}

func (r *Reservoir) CognizeY(eye *be.Eye, v interface{}) {}

/*
	{
		Command string
		Gate string
		Meaning *
	}
*/
func (r *Reservoir) CognizeX(eye *be.Eye, v interface{}) {
	r.Lock()
	defer r.Unlock()
	u := v.(Circuit)
	switch u.StringAt("Command") {
	case "Open":
		r.path.PushBack(r.focus)
		r.focus = r.focus.CircuitAt(u.At("Gate"))

	case "Close":
		r.focus = r.path.Remove(r.path.Back()).(Circuit)

	case "Include":
		if _, over := r.focus.Include(u.At("Gate"), u.At("Meaning")); over {
			panic("over including")
		}

	case "Exclude":
		if _, forgotten := r.focus.Exclude(u.At("Gate")); !forgotten {
			panic("nothing to exclude")
		}

	case "Link":
		a, b := Vector(u.CircuitAt(0)), Vector(u.CircuitAt(1))
		r.focus.Link(a, b)

	case "Unlink":
		a, b := Vector(u.CircuitAt(0)), Vector(u.CircuitAt(1))
		r.focus.Unlink(a, b)

	case "Yield":
		r.stop.Do(func() { eye.Show("Y", r.y) })
	}
}

//
type OpenCommand Circuit

func NewOpenCommand(gate Name) OpenCommand {
	return OpenCommand(New().Grow("Command", "Open").Grow("Gate", gate))
}

func (x OpenCommand) Reduce() (gate Name) {
	return x.At("Gate")
}

//
type CloseCommand Circuit

func NewCloseCommand(gate Name) CloseCommand {
	return CloseCommand(New().Grow("Command", "Close").Grow("Gate", gate))
}

//
type IncludeCommand Circuit

func NewIncludeCommand(gate Name, meaning Meaning) IncludeCommand {
	return IncludeCommand(New().Grow("Command", "Include").Grow("Gate", gate).Grow("Meaning", meaning))
}

func (x IncludeCommand) Reduce() (gate Name, meaning Meaning) {
	return x.At("Gate"), x.At("Meaning")
}

//
type ExcludeCommand Circuit

func NewExcludeCommand(gate Name) ExcludeCommand {
	return ExcludeCommand(New().Grow("Command", "Exclude").Grow("Gate", gate))
}

func (x ExcludeCommand) Reduce() (gate Name) {
	return x.At("Gate")
}

//
type LinkCommand Circuit

func NewLinkCommand(u, v Vector) LinkCommand {
	return LinkCommand(New().Grow("Command", "Link").Grow("U", Circuit(u)).Grow("V", Circuit(v)))
}

func (x LinkCommand) Reduce() (u, v Vector) {
	return Vector(Circuit(x).CircuitAt("U")), Vector(Circuit(x).CircuitAt("V"))
}

//
type UnlinkCommand Circuit

func NewUnlinkCommand(u, v Vector) UnlinkCommand {
	return UnlinkCommand(New().Grow("Command", "Unlink").Grow("U", Circuit(u)).Grow("V", Circuit(v)))
}

func (x UnlinkCommand) Reduce() (u, v Vector) {
	return Vector(Circuit(x).CircuitAt("U")), Vector(Circuit(x).CircuitAt("V"))
}

//
type YieldCommand Circuit

func NewYieldCommand() YieldCommand {
	return YieldCommand(New().Grow("Command", "Yield"))
}
