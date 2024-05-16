package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vivekup3424/shopping-cart/types"
)

type Handler struct {
}

// handler can take any dependency, hence dependency injection
// comes here
func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleLogin).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//get jsong payloaf, and check if user already exists
	var payload types.RegisterUserPayload

	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		w.WriteHeader(502)
		json.NewEncoder(w).Encode(`\
		{
			"Error":"Bad Request"
		}
		`)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(payload)

}
