package middleware

import (
	"net/http"

	"github.com/gianlucapastori/go-auth-jwt/cfg"
	"github.com/gianlucapastori/go-auth-jwt/internal/packages/users"
	"go.uber.org/zap"
)

type Middleware struct {
	cfg   *cfg.Config
	serv  users.Services
	sugar *zap.SugaredLogger
}

func New(cfg *cfg.Config, serv users.Services, sugar *zap.SugaredLogger) *Middleware {
	return &Middleware{cfg: cfg, serv: serv, sugar: sugar}
}

// default "generic" middlewares
func (mw *Middleware) JSON(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		h.ServeHTTP(w, r)
	})
}
