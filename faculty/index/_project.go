// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package index

import (
	"fmt"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

// TODO: This is a first order graph projection. It can be extended to any order.
type Index struct {
	gv Circuit // Gate-Valve dictionary
	shadow Circuit
}

func (x *Index) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	x.gv, x.shadow = New(), New()
	return nil
}

func (x *Index) CognizeFlowFrame(eye *be.Eye, v interface{}) {

	// place involved gates in dictionary
	f := v.(Circuit)
	println(fmt.Sprintf("-> %v\n", f.Print("", "  ", 2)))

	ag := x.remember(f, 0)
	bg := x.remember(f, 1)

	// place flow links in shadow
	ak, bk := x.shadow.Degree(ag), x.shadow.Degree(bg)
	x.shadow.Link(Vector{ag, ak}, Vector{bg, bk})
}

//	….Gate.Valve.[Index]
func (x *Index) remember(frame Circuit, i int) (index int) {

	half := frame.CircuitAt(i)
	valve := half.NameAt("Valve")

	// remember gate in Gate-Valve dictionary
	var g Circuit
	switch t := half.At("Value").(type) {
	case Address:
		g = x.gv.Refine(t.Path...)
	case int:
		g = x.gv.Refine("int")
	case string:
		g = x.gv.Refine("string")
	case float64:
		g = x.gv.Refine("float64")
	case complex128:
		g = x.gv.Refine("complex128")
	case Circuit:
		g = x.gv.Refine("meaningless")
	case nil:
		g = x.gv.Refine("missing")
	default:
		g = x.gv.Refine("unknown")
	}

	// and valve
	var ok bool
	if index, ok = g.IntOptionAt(valve); !ok {
		index = x.shadow.Len()
		g.Gate[valve] = index
	}

	// place gate in index
	x.shadow.Include(valve, "GateValve")

	return index
}

func (x *Index) CognizeFlush(eye *be.Eye, v interface{}) {
	println("flush")
	gv, shadow := x.gv, x.shadow
	x.gv, x.shadow = New(), New()
	eye.Show(DefaultValve, New().Grow("GateValve", gv).Grow("Shadow", shadow))
}

func (x *Index) Cognize(eye *be.Eye, v interface{}) {}
