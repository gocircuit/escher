// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	"github.com/gocircuit/circuit/client"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/plumb"
)

// Joining
type Joining struct{}

func (x Joining) Materialize() be.Reflex {
	return MaterializeSubscription("Joining")
}

// Leaving
type Leaving struct{}

func (x Leaving) Materialize() be.Reflex {
	return MaterializeSubscription("Leaving")
}

func MaterializeSubscription(kind string) be.Reflex {
	reflex, eye := be.NewEye("Server", "_")
	go func() {
		var server string
		for {
			valve, value := eye.See()
			?? // use containerlike spawn param with server and name!!
			if valve != "Server" || server != "" {
				continue
			}
			server = value.(string)
			go func() {
				id := ChooseID()
				anchor := program.Client.Walk(
					??
					[]string{
						server, 
						"escher", 
						program.Name, 
						"circuit." + kind, 
						id,
					},
				)
				var ss client.Subscription
				var err error
				switch kind {
				case "Joining":
					ss, err = anchor.MakeOnJoin()
				case "Leaving":
					ss, err = anchor.MakeOnLeave()
				default:
					panic(2)
				}
				if err != nil {
					panic("plugging")
				}
				defer anchor.Scrub()
				for {
					v, ok := ss.Consume()
					if !ok {
						panic("subscription should not be closing ever")
					}
					eye.Show("_", v)
				}
			}()
		}
	}()
	return reflex
}
