// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package twitter installs a faculty for access to the Twitter API.
package twitter

import (
	// "fmt"

	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/kit/plumb"
	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/think"

	"github.com/gocircuit/escher/ChimeraCoder/anaconda"
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
			userTimelineAnswer: plumb.NewMatching(),
			consumer: plumb.NewCondition(),
			access: plumb.NewCondition(),
		}
		h.userTimelineAnswer.Connect(userTimelineAnswerEndo.Focus(think.DontCognize))
		consumerEndo.Focus(p.CognizeConsumer)
		accessEndo.Focus(p.CognizeAccess)
		userTimelineQueryEndo.Focus(p.CognizeUserTimelineQuery)
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
	userTimelineAnswer *plumb.Speak
	userTimelineQuery *plumb.Hear
	//
	consumer *plumb.Condition
	access *plumb.Condition
}

func (h *client) CognizeConsumer(v interface{}) {
	h.consumer.Determine(v)
}

func (h *client) CognizeAccess(v interface{}) {
	h.access.Determine(v)
}

func (h *client) CognizeUserTimelineQuery(v interface{}) {
	h.userTimelineQueryChan <- v
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