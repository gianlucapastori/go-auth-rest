package http

import (
	"github.com/gianlucapastori/nausicaa/internal/middleware"
	"github.com/gianlucapastori/nausicaa/internal/packages/users"
	"github.com/gorilla/mux"
)

func Map(mux *mux.Router, cont users.Controller, mw *middleware.Middleware) {
	mux.HandleFunc("/register", cont.RegisterUser()).Methods("POST")
}
