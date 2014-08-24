// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package twitter installs a faculty for access to the Twitter API.
package twitter

import (
	// "fmt"
	"net/url"
	"strconv"
	"sync"

	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/faculty/basic"
	"github.com/gocircuit/escher/kit/plumb"
	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/think"

	"github.com/gocircuit/escher/github.com/ChimeraCoder/anaconda"
)

func init() {
	ns := faculty.Root.Refine("web").Refine("twitter")
	ns.AddTerminal("Client", Client{})
	ns.AddTerminal("ForkAnswer", ForkAnswer{})
	ns.AddTerminal("ForkConsumer", ForkConsumer{})
	ns.AddTerminal("ForkAccess", ForkAccess{})
	ns.AddTerminal("ForkUserTimelineQuery", ForkUserTimelineQuery{})
}

// ForkAnswer…
type ForkAnswer struct{}

func (ForkAnswer) Materialize() think.Reflex {
	return basic.MaterializeFork("_", "Name", "Sentence")
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
		consumerEndo.Focus(h.consumer.Determine) // Consumer
		accessEndo.Focus(h.access.Determine) // Access
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
		select {
		case t := <-h.userTimelineQuery.Chan():
			q := t.(Image)
			uv := url.Values{}
			uv.Set("user_id", q.OptionalString("UserId"))
			uv.Set("screen_name", q.OptionalString("ScreenName"))
			uv.Set("since_id", strconv.Itoa(q.OptionalInt("AfterId"))) // return results indexed greater than since_id
			uv.Set("max_id", strconv.Itoa(q.OptionalInt("NotAfterId"))) // return results indexed no greater than max_id
			uv.Set("count", strconv.Itoa(q.OptionalInt("Count")))
			timeline, err := api.GetUserTimeline(uv)
			if err != nil {
				panic(err)
			}
			userTimelineAnswer.ReCognize(
				Image{
					"Name": q.Interface("Name"),
					"Sentence": Imagine(timeline),
				},
			)
		}
	}
}
