// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	"github.com/gocircuit/circuit/client"
	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/faculty/basic"
)

// Joining
type Joining struct{}

func (x Joining) Materialize() think.Reflex {
	return MaterializeSubscription("Joining")
}

// Leaving
type Leaving struct{}

func (x Leaving) Materialize() think.Reflex {
	return MaterializeSubscription("Leaving")
}

func MaterializeSubscription(kind string) think.Reflex {
	_Endo, _Exo := think.NewSynapse()
	serverEndo, serverExo := think.NewSynapse()
	go func() {
		p := &subscription{
			kind: kind,
			id:    ChooseID(),
			z: basic.NewConnector(),
			server: basic.NewQuestion(),
		}
		p.z.Connect(_Endo.Focus(think.DontCognize))
		serverEndo.Focus(p.CognizeServer)
		p.loop()
	}()
	return think.Reflex{
		"_":  _Exo,
		"Server":  serverExo,
	}
}

// subscription is the materialized subscription reflex
type subscription struct {
	kind string // “Joining” or “Leaving”
	id string
	z *basic.Connector
	server *basic.Question
}

func (h *subscription) CognizeServer(v interface{}) {
	srv, ok := v.(string)
	if !ok {
		panic("process server anchor is non-string")
	}
	h.server.Lock()
	defer h.server.Unlock()
	h.server.Answer(srv)
}

func (h *subscription) loop() {
	z := h.z.Connected()
	anchor := program.Client.Walk([]string{h.server.String(), "escher", program.Name, "circuit." + h.kind, h.id})
	var ss client.Subscription
	var err error
	switch h.kind {
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
		z.ReCognize(v)
	}
}
