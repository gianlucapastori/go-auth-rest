package http

import (
	"github.com/gianlucapastori/nausicaa/internal/middleware"
	"github.com/gianlucapastori/nausicaa/internal/packages/users"
	"github.com/gorilla/mux"
)

func Map(mux *mux.Router, controller users.Controller, mw *middleware.Middleware) {
	mux.HandleFunc("/register", controller.RegisterUser()).Methods("POST")
}
