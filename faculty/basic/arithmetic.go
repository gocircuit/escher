// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	"strconv"

	"github.com/gocircuit/escher/be"
	cir "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	// faculty.Register("Sum", Sum{})
	faculty.Register(be.NewMaterializer(IntString{}), "e", "IntString")
}

// IntString
type IntString struct{}

func (IntString) Spark(*be.Eye, cir.Circuit, ...interface{}) cir.Value {
	return IntString{}
}

func (IntString) CognizeInt(eye *be.Eye, v interface{}) {
	eye.Show("String", strconv.Itoa(v.(int)))
}

func (IntString) CognizeString(eye *be.Eye, v interface{}) {
	i, err := strconv.Atoi(v.(string))
	if err != nil {
		panic(err)
	}
	eye.Show("Int", i)
}
