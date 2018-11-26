package core

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server holds the shared dependencies of the service
type Server struct {
	Router *mux.Router
	DB     *sql.DB
	// todo logger
}

// Init initializes the server instance
func (s *Server) Init(routes Routes) {
	s.routes(routes)
	s.serve()
}

func InitRouter() *mux.Router {
	return mux.NewRouter().StrictSlash(true)
}

// Serve serves the application :)
func (s *Server) serve() {
	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(originsOk, headersOk, methodsOk)(s.Router)))
}

func (s *Server) routes(routes Routes) {
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc(s)
		// handler = Logger(handler, route.Name)

		s.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
}