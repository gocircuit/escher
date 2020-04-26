// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package http

import (
	"io"
	"net/http"
	"sync"

	"github.com/hoijui/escher/be"
	cir "github.com/hoijui/escher/circuit"
	"github.com/hoijui/escher/faculty"
)

func init() {
	faculty.Register(be.NewMaterializer(&Server{}), "http", "Server")
}

type Server struct {
	eye    *be.Eye
	matter cir.Circuit
	sync.Mutex
	server   *http.Server
	throttle chan struct{}
}

func (s *Server) Spark(eye *be.Eye, matter cir.Circuit, aux ...interface{}) cir.Value {
	s.eye, s.matter = eye, matter
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
	u := value.(cir.Circuit)
	if s.server != nil {
		panic("server running")
	}
	s.server = &http.Server{
		Addr:    u.StringAt("Address"),
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
	xx, yy := be.NewEntanglement()
	ch := make(chan struct{}, 1)
	go xx.Synapse().Connect( // connect to the server-side of the entanglement
		func(v interface{}) {
			defer func() {
				ch <- struct{}{} // release throttle token when request/response complete
			}()
			status, body, ok := s.cognizeResponse(w.Header(), v.(cir.Circuit))
			if !ok {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Escher web server: App error."))
				return
			}
			defer body.Close()
			w.WriteHeader(status)
			io.Copy(w, body)
		},
	)
	s.eye.Show(
		"RequestResponse",
		cir.New().
			Grow("Request", requestCircuit(req)).
			Grow("Respond", yy),
	)
	<-ch
}
