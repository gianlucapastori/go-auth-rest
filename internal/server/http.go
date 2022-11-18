package server

import (
	"fmt"
	"net/http"

	"github.com/gianlucapastori/nausicaa/internal/middleware"
	usersHttp "github.com/gianlucapastori/nausicaa/internal/packages/users/http"
	usersRepo "github.com/gianlucapastori/nausicaa/internal/packages/users/repo"
	usersServices "github.com/gianlucapastori/nausicaa/internal/packages/users/services"
	"github.com/gianlucapastori/nausicaa/pkg/utils"
	"github.com/gianlucapastori/nausicaa/pkg/validator"
	"github.com/gorilla/mux"
)

func (s *Server) Map(mux *mux.Router, mw *middleware.Middleware) error {
	validator.New()

	// base middlewares
	mux.Use(mw.JSON)
	mux.Use(mw.CORS)

	// repositories
	uRepo := usersRepo.New(s.db)

	// services
	uServ := usersServices.New(uRepo, s.cfg, s.sugar)

	// controllers
	uCont := usersHttp.New(uServ, s.cfg, s.sugar)

	vRoute := mux.PathPrefix(fmt.Sprintf("/api/v%s", s.cfg.SERVER.VERSION)).Subrouter()

	// healthcheck route
	vRoute.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		utils.Respond(w, http.StatusOK, "pong! :)")
	}).Methods("GET")

	uRoute := vRoute.PathPrefix("/users").Subrouter()

	usersHttp.Map(uRoute, uCont, mw)

	return nil
}
