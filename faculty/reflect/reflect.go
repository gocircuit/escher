// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package reflect

import (
	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/be"
)

func init() {
	faculty.Register("reflect.Index", be.NewNativeMaterializer(&Index{}))
	faculty.Register("reflect.Address", be.NewNativeMaterializer(Addresses{}))
}
