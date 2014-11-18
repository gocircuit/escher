// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package http

import (
	// "fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/kit/plumb"
	"github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/circuit"
)

func init() {
	faculty.Register("http.Server", be.NewNativeMaterializer(&Server{}))
}

type Server struct {
	eye *be.Eye
	sync.Mutex
	server *http.Server
	throttle chan struct{}
}

func (s *Server) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	s.eye = eye
	const throttle = 50
	s.throttle = make(chan struct{}, throttle)
	for i := 0; i < throttle; i++ {
		s.throttle <- struct{}{}
	}
	return nil
}

func (s *Server) CognizeRequestResponse(eye *be.Eye, value interface{}) {}

func (s *Server) CognizeStart(eye *be.Eye, value interface{}) {
	s.Lock()
	defer s.Unlock()
	//
	u := value.(Circuit)
	if s.server != nil {
		panic("server running")
	}
	s.server = &http.Server{
		Addr: u.StringAt("Address"),
		Handler: s,
	}
	go func() {
		s.server.ListenAndServe()
	}()
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	<-s.throttle
	defer func() {
		s.throttle <- struct{}{}
	}()
	//
	mx, my := be.NewEntanglement()
	ch := make(chan struct{}, 1)
	go mx.Synapse().Focus(  // MUST throttle
		func (v interface{}) {
			resp := v.(Circuit)
			h := w.Header()
			g := resp.CircuitAt("Header")
			for _, k := range g.SortedLetters() {
				h[k] = circuitSlice(g.CircuitAt(k))
			}
			w.WriteHeader(resp.IntAt("Status"))
			w.Write([]byte(plumb.AsString(resp.At("Body"))))
			ch <- struct{}{}
		},
	)
	s.eye.Show(
		"RequestResponse", 
		New().
			Grow("Request", requestCircuit(req)).
			Grow("Respond", my),
	)
	<-ch
}

func requestCircuit(req *http.Request) Circuit {
	x := New()

	// HTTP method
	x.Gate["Method"] = req.Method

	// URL path
	var nn []Name
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) > 0 && parts[0] == "" {
		parts = parts[1:]
	}
	for _, n := range parts {
		nn = append(nn, n)
	}
	x.Gate["Path"] = NewAddress(nn...)

	// URL query
	v := New()
	for k, ss := range req.URL.Query() {
		v.Gate[k] = sliceCircuit(ss)
	}
	x.Gate["Query"] = v

	return x
}

func sliceCircuit(ss []string) Circuit {
	x := New()
	for i, v := range ss {
		x.Gate[i] = v
	}
	return x
}

func circuitSlice(u Circuit) []string {
	var ss []string
	for _, j := range u.SortedNumbers() {
		ss = append(ss, u.StringAt(j))
	}
	return ss
}
