package middleware

import (
	"net/http"

	"github.com/gianlucapastori/nausicaa/cfg"
	"go.uber.org/zap"
)

type Middleware struct {
	cfg   *cfg.Config
	sugar *zap.SugaredLogger
}

func New(cfg *cfg.Config, sugar *zap.SugaredLogger) *Middleware {
	return &Middleware{cfg: cfg, sugar: sugar}
}

// default "generic" middlewares
func (mw *Middleware) JSON(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		h.ServeHTTP(w, r)
	})
}
