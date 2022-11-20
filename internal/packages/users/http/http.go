package http

import (
	"github.com/gianlucapastori/nausicaa/internal/middleware"
	"github.com/gianlucapastori/nausicaa/internal/packages/users"
	"github.com/gorilla/mux"
)

func Map(mux *mux.Router, cont users.Controller, mw *middleware.Middleware) {
	r := mux.PathPrefix("/protected").Subrouter()
	r.Use(mw.AuthJWT)
	r.HandleFunc("", cont.Protected())

	r = mux.PathPrefix("/refresh-token").Subrouter()
	r.Use(mw.AuthRefresh)
	r.HandleFunc("", cont.RequestNewAccess())

	mux.HandleFunc("/register", cont.RegisterUser()).Methods("POST")
}
