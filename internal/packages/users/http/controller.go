package http

import (
	"net/http"

	"github.com/gianlucapastori/nausicaa/cfg"
	"github.com/gianlucapastori/nausicaa/internal/packages/users"
	"github.com/gianlucapastori/nausicaa/pkg/utils"
	"go.uber.org/zap"
)

type userController struct {
	services users.Services
	cfg      *cfg.Config
	sugar    *zap.SugaredLogger
}

func New(services users.Services, cfg *cfg.Config, sugar *zap.SugaredLogger) users.Controller {
	return &userController{services: services, cfg: cfg, sugar: sugar}
}

func (*userController) RegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Respond(w, 200, "oi")
	}
}
