// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package browser

import (
	"net/http"
	"path"
	"sync"

	"github.com/petar/maymounkov.io/code.google.com/p/go.net/websocket"
)

// Server is a static file and websocket server.
type Server struct {
	sync.Mutex
	http *http.Server
}

// SessionPlayer is a handler for an incoming WebSocket session
type SessionPlayer func(*http.Request, *Session)

// Start a new websocket server
func NewServer(addr, dir string) *Server {
	var srv = &Server{
		http: &http.Server{
			Addr:           addr,
			Handler:        http.NewServeMux(),
			ReadTimeout:    0,
			WriteTimeout:   0,
			MaxHeaderBytes: 0,
			TLSConfig:      nil,
			TLSNextProto:   nil,
		},
	}
	// By default all paths are served statically from the local dir.
	srv.http.Handler.(*http.ServeMux).Handle("/", http.FileServer(http.Dir(dir)))
	return srv
}

// Run starts the web server and blocks.
func (srv *Server) Run() {
	if err := srv.http.ListenAndServe(); err != nil {
		panic(err)
	}
}

// Capture the request object and …
type playerHandler SessionPlayer

func (ph playerHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	websocket.Handler(func(conn *websocket.Conn) {
		ph(req, NewSession(conn))
	}).ServeHTTP(w, req)
}

// Add adds a handler for URL paths beginning with name which hand-off the resulting websocket to player.
func (srv *Server) Add(name string, player SessionPlayer) {
	srv.Lock()
	defer srv.Unlock()
	srv.http.Handler.(*http.ServeMux).Handle(
		path.Join("/", name), 
		playerHandler(player),
	)
}
