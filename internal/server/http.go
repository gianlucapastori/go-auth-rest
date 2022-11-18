package server

import (
	"fmt"
	"net/http"

	"github.com/gianlucapastori/nausicaa/internal/middleware"
	"github.com/gianlucapastori/nausicaa/pkg/utils"
	"github.com/gorilla/mux"
)

func (s *Server) Map(mux *mux.Router, mw *middleware.Middleware) error {
	mux.Use(mw.JSON)
	mux.Use(mw.CORS)

	vRoute := mux.PathPrefix(fmt.Sprintf("/api/v%s", s.cfg.SERVER.VERSION)).Subrouter()

	vRoute.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		utils.Respond(w, http.StatusOK, "pong! :)")
	}).Methods("GET")

	return nil
}
