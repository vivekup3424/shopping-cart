package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vivekup3424/shopping-cart/services/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (A *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	userService := user.NewHandler()
	userService.RegisterRoutes(subrouter) //what the heck is this shit
	//I dont even understand, I am just copying pasting
	log.Printf("listening on: %s", A.addr)
	return http.ListenAndServe(A.addr, router)
}
