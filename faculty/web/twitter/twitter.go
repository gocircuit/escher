// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package twitter installs a faculty for access to the Twitter API.
package twitter

import (
	// "fmt"
	"log"
	"net/url"
	"strconv"
	"sync"

	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/kit/plumb"
	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/think"

	"github.com/gocircuit/escher/github.com/ChimeraCoder/anaconda"
)

func init() {
	ns := faculty.Root.Refine("web").Refine("twitter")
	ns.AddTerminal("Client", Client{})
	ns.AddTerminal("Answer", AnswerMaterializer{})
	ns.AddTerminal("Consumer", ConsumerMaterializer{})
	ns.AddTerminal("Access", AccessMaterializer{})
	ns.AddTerminal("UserTimelineQuery", UserTimelineQueryMaterializer{})
}

// Client ...
type Client struct{}

func (Client) Materialize() think.Reflex {
	var c1, a1 sync.Once
	api, consumer, access := make(chan *anaconda.TwitterApi), make(chan Image, 1), make(chan Image, 1)
	go func() { // start connecting monad
		var c, a Image
		for i := 0; i < 2; i++ {
			select {
			case c = <-consumer:
				consumer = nil
			case a = <-access:
				access = nil
			}
		}
		anaconda.SetConsumerKey(c.String("Key")) // dial API server
		anaconda.SetConsumerSecret(c.String("Secret"))
		y := anaconda.NewTwitterApi(a.String("Token"), a.String("Secret"))	
		for {
			api <- y // give out api server to all endpoint goroutines
		}
	}()
	userTimelineQuery := make(chan Image, 5)
	reflex, eye := plumb.NewEyeCognizer(
		func (eye *plumb.Eye, valve string, value interface{}) {
			switch valve {
			case "Consumer":
				c1.Do(func () {consumer <- value.(Image)})
			case "Access":
				a1.Do(func() {access <- value.(Image)})
			case "UserTimelineQuery":
				userTimelineQuery <- value.(Image)
 			}
		},
		"Consumer", "Access", // set to start connection
		"UserTimelineQuery", "UserTimelineResult", // UserTimeline
	)
	go func() { // UserTimeline loop
		y := <-api
		for {
			q := <-userTimelineQuery
			uv := url.Values{}
			uv.Set("user_id", q.OptionalString("UserId"))
			uv.Set("screen_name", q.OptionalString("ScreenName"))
			if q.Has("AfterId") {
				uv.Set("since_id", strconv.Itoa(q.OptionalInt("AfterId"))) // return results indexed greater than since_id
			}
			if q.Has("NotAfterId") {
				uv.Set("max_id", strconv.Itoa(q.OptionalInt("NotAfterId"))) // return results indexed no greater than max_id
			}
			if q.Has("Count") {
				uv.Set("count", strconv.Itoa(q.OptionalInt("Count")))
			}
			log.Printf("Twitter query %v", ImagineWithMaps(uv).(Image).PrintLine())
			timeline, err := y.GetUserTimeline(uv)
			if err != nil {
				log.Fatalf("Problem with Twitter (%v)", err)
			}
			eye.Show(
				"UserTimelineResult",
				Pretty(
					Image{
						"Name": q.Interface("Name"),
						"Sentence": Imagine(timeline),
					},
				),
			)
		}
	}()
	return reflex
}
