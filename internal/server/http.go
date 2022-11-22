package server

import (
	"fmt"
	"net/http"

	"github.com/gianlucapastori/nausicaa/internal/middleware"
	groupsHttp "github.com/gianlucapastori/nausicaa/internal/packages/groups/http"
	groupsRepo "github.com/gianlucapastori/nausicaa/internal/packages/groups/repo"
	groupsServices "github.com/gianlucapastori/nausicaa/internal/packages/groups/services"
	usersHttp "github.com/gianlucapastori/nausicaa/internal/packages/users/http"
	usersRepo "github.com/gianlucapastori/nausicaa/internal/packages/users/repo"
	usersServices "github.com/gianlucapastori/nausicaa/internal/packages/users/services"
	"github.com/gianlucapastori/nausicaa/pkg/utils"
	"github.com/gianlucapastori/nausicaa/pkg/validator"
	"github.com/gorilla/mux"
)

func (s *Server) Map(mux *mux.Router) error {
	validator.New()

	// repositories
	uRepo := usersRepo.New(s.db)
	gRepo := groupsRepo.New(s.db)

	// services
	uServ := usersServices.New(uRepo, s.cfg, s.sugar)
	gServ := groupsServices.New(gRepo, s.sugar, s.cfg)

	// controllers
	uCont := usersHttp.New(uServ, s.cfg, s.sugar)
	gCont := groupsHttp.New(gServ, s.cfg, s.sugar)

	mw := middleware.New(s.cfg, uServ, s.sugar)

	// base middlewares
	mux.Use(mw.JSON)
	mux.Use(mw.CORS)

	vRoute := mux.PathPrefix(fmt.Sprintf("/api/v%s", s.cfg.SERVER.VERSION)).Subrouter()

	// healthcheck route
	vRoute.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		utils.Respond(w, http.StatusOK, "pong! :)")
	}).Methods("GET")

	uRoute := vRoute.PathPrefix("/users").Subrouter()
	gRoute := vRoute.PathPrefix("/groups").Subrouter()

	usersHttp.Map(uRoute, uCont, mw)
	groupsHttp.Map(gRoute, gCont, mw)

	return nil
}
