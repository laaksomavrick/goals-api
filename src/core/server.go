package core

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/laaksomavrick/goals-api/src/config"
)

// Server holds the essential shared dependencies of the service
type Server struct {
	Router *mux.Router
	DB     *sql.DB
	Config *config.Config
}

// NewServer constructs a new instance of a server
func NewServer(router *mux.Router, db *sql.DB, config *config.Config) *Server {
	return &Server{
		Router: router,
		DB:     db,
		Config: config,
	}
}

// Init applies the middleware stack, registers route handlers, and serves the application
func (s *Server) Init(routes Routes) {
	s.Wire(routes)
	s.Serve()
}

// Serve serves the application :)
func (s *Server) Serve() {
	port := fmt.Sprintf(":%s", s.Config.Port)
	fmt.Printf("Server listening on port: %s\n", port)
	log.Fatal(http.ListenAndServe(port, s.Router))
}

// Wire applies middlewares to all routes and registers them to the Server.Router
func (s *Server) Wire(routes Routes) {
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc(s)
		handler = Logger(handler, route.Name)

		s.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

		headersOk := handlers.AllowedHeaders([]string{"Authorization"})
		originsOk := handlers.AllowedOrigins([]string{"*"})
		methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

		handlers.CORS(originsOk, headersOk, methodsOk)(s.Router)

	}
}
