// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package url installs a faculty for URL manipulations.
package url

import (
	// "fmt"
	"net/url"

	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/kit/plumb"
	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/think"
)

func init() {
	faculty.Root.AddTerminal("Client", Client{})
}

// Client ...
type Client struct{}

func (Client) Materialize() think.Reflex {
	consumerEndo, consumerExo := think.NewSynapse()
	accessEndo, accessExo := think.NewSynapse()
	userTimelineQueryEndo, userTimelineQueryExo := think.NewSynapse()
	userTimelineAnswerEndo, userTimelineAnswerExo := think.NewSynapse()
	go func() {
		h := &client{
			userTimelineQuery: plumb.NewHear(),
			userTimelineAnswer: plumb.NewSpeak(),
			consumer: plumb.NewCondition(),
			access: plumb.NewCondition(),
		}
		consumerEndo.Focus(p.consumer.Determine) // Consumer
		accessEndo.Focus(p.access.Determine) // Access
		userTimelineQueryEndo.Focus(h.userTimelineQuery.Cognize) // UserTimelineQuery
		h.userTimelineAnswer.Connect(userTimelineAnswerEndo.Focus(think.DontCognize)) // UserTimelineAnswer
		go h.loop()
	}()
	return think.Reflex{
		"Consumer":   consumerExo, // key and secret
		"Access":   accessExo, // access token and secret
		"UserTimelineQuery": userTimelineQueryExo,
		"UserTimelineAnswer": userTimelineAnswerExo,
	}
}

type client struct {
	sync.Mutex
	//
	userTimelineQuery *plumb.Hear
	userTimelineAnswer *plumb.Speak
	//
	consumer *plumb.Condition
	access *plumb.Condition
}

func (h *client) loop() {
	userTimelineAnswer := h.userTimelineAnswer.Connected()
	//
	consumer := h.consumer.Image() // wait for consumer key and secret information
	anaconda.SetConsumerKey(consumer.String("Key")) // dial API server
	anaconda.SetConsumerSecret(consumer.String("Secret"))
	//
	access := h.access.Image() // wait for access token and access token secret
	api := anaconda.NewTwitterApi(access.String("Token"), access.String("Secret"))	
	//
	for {
	}
}
