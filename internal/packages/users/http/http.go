package http

import (
	"github.com/gianlucapastori/nausicaa/internal/middleware"
	"github.com/gianlucapastori/nausicaa/internal/packages/users"
	"github.com/gorilla/mux"
)

func Map(mux *mux.Router, cont users.Controller, mw *middleware.Middleware) {
	r := mux.PathPrefix("/protected").Subrouter()
	r.Use(mw.AuthJWT)
	r.HandleFunc("", cont.Protected()).Methods("GET")

	r = mux.PathPrefix("/refresh-token").Subrouter()
	r.Use(mw.AuthRefresh)
	r.HandleFunc("", cont.RequestNewAccess()).Methods("GET")

	r = mux.PathPrefix("/change-password").Subrouter()
	r.Use(mw.AuthPwd)
	r.HandleFunc("", cont.ChangePassword()).Methods("POST")

	mux.HandleFunc("/login", cont.LoginUser()).Methods("POST")

	mux.HandleFunc("/register", cont.RegisterUser()).Methods("POST")

	mux.HandleFunc("/new-password", cont.RequestNewPassword()).Methods("POST")
}
