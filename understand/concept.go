// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package understand

import (
	"fmt"

	"github.com/gocircuit/escher/see"
)

// TODO: Eventually Circuit can be represented as an image.Image
type Circuit struct {
	Genus []*see.Circuit // origin
	Name  string
	Peer map[string]*Peer // peers and self; self corresponds to the empty string
}

type Peer struct {
	Name   string
	Design interface{}
	Valve  map[string]*Valve
}

type Valve struct {
	Of       *Peer
	Name     string
	Matching *Valve
}

/*
	circuit {
		Genus *see.Circuit
		Name string
		Peer { 
			0 {
				Name string
				Design interface{}
				Valve {
					0 int
					...
				}
			}
			...
		}
		Valve {
			0 {
				Peer int
				Name string
			}
			...
		}
		Matching {
			0 {
				Left int // left valve index
				Right int
			}
			...
		}
	}
*/
