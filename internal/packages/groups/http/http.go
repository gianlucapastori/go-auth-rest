package http

import (
	"github.com/gianlucapastori/nausicaa/internal/middleware"
	"github.com/gianlucapastori/nausicaa/internal/packages/groups"
	"github.com/gorilla/mux"
)

func Map(mux *mux.Router, cont groups.Controller, mw *middleware.Middleware) {

}
