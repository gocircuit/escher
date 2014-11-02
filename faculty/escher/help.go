// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	"fmt"
	"os"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Register("help", be.NewNativeMaterializer(&Help{}))
	faculty.Register("Help", be.NewNativeMaterializer(&Help{}))
}

type Help struct {
	idiom Circuit
}

func (h *Help) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	h.idiom = matter.Idiom
	return nil
}

func (h *Help) Cognize(eye *be.Eye, v interface{}) {
	switch u := v.(type) {
	case Address:
		h.address(u)
	case Circuit:
		h.circuit(u)
	case int:
		fmt.Fprintf(os.Stderr, "\nThis is  an integer constant equal to %v.\n\n", u)
	case float64:
		fmt.Fprintf(os.Stderr, "\nThis is  a float constant equal to %v.\n\n", u)
	case complex128:
		fmt.Fprintf(os.Stderr, "\nThis is  a complex constant equal to %v.\n\n", u)
	case string:
		fmt.Fprintf(os.Stderr, "\nThis is a string constant equal to %q.\n\n", u)
	default:
		fmt.Fprintf(os.Stderr, "\nThis is a value of uncommon type %T equal to %v.\n\n", u, u)
	}
}

func (h *Help) circuit(u Circuit) {
	fmt.Fprintf(os.Stderr, "\nWe are looking at a circuit design (in desugared syntax):\n%v\n", u)
}

func (h *Help) address(addr Address) {

	fmt.Fprintf(os.Stderr, "\nAddress %v resolves to ??", addr)
}
