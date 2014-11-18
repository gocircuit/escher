// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	"fmt"

	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

func init() {
	faculty.Register("Ignore", be.Ignore{})
	faculty.Register("Idle", be.NewIdleMaterializer())
	faculty.Register("Scanln", Scanln{})
	//
	faculty.Register("Grow", be.NewNativeMaterializer(&Grow{}))
	faculty.Register("Fork", be.NewNativeMaterializer(&be.Union{}))
	faculty.Register("Lens", be.NewNativeMaterializer(&Lens{}))
	//
	faculty.Register("Alternate", be.NewNativeMaterializer(&Alternate{}))
	faculty.Register("Alt", be.NewNativeMaterializer(&Alternate{}))
	faculty.Register("OneWayDoor", be.NewNativeMaterializer(&OneWayDoor{}))
	//
	faculty.Register("Repeat", be.NewNativeMaterializer(Repeat{}))
	//
	faculty.Register("OnEnd", be.NewNativeMaterializer(&OnHash{}, "#End")) // ??
}

// Scanln
type Scanln struct{}

func (Scanln) Materialize(*be.Matter) (be.Reflex, Value) {
	s, t := be.NewSynapse()
	go func() {
		r := s.Focus(be.DontCognize)
		go func() {
			for {
				var em string
				fmt.Scanln(&em)
				r.ReCognize(em)
			}
		}()
	}()
	return be.Reflex{DefaultValve: t}, Scanln{}
}

// Println
type Println struct{}

func (Println) Materialize(*be.Matter) (be.Reflex, Value) {
	s, t := be.NewSynapse()
	go func() {
		s.Focus(
			func(v interface{}) {
				fmt.Printf("%v\n", v)
			},
		)
	}()
	return be.Reflex{DefaultValve: t}, Println{}
}
