package server

import (
	"fmt"
	"net/http"

	"github.com/gianlucapastori/nausicaa/pkg/utils"
	"github.com/gorilla/mux"
)

func (s *Server) Map(mux *mux.Router) error {
	vRoute := mux.PathPrefix(fmt.Sprintf("/api/v%s", s.cfg.SERVER.VERSION)).Subrouter()

	vRoute.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(w, r)
		})
	})

	vRoute.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		utils.Respond(w, http.StatusOK, "pong! :)")
	}).Methods("GET")

	return nil
}
