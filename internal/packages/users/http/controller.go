package http

import (
	"net/http"

	"github.com/gianlucapastori/nausicaa/cfg"
	"github.com/gianlucapastori/nausicaa/internal/entities"
	"github.com/gianlucapastori/nausicaa/internal/packages/users"
	"github.com/gianlucapastori/nausicaa/pkg/utils"
	"go.uber.org/zap"
)

type userController struct {
	serv  users.Services
	cfg   *cfg.Config
	sugar *zap.SugaredLogger
}

func New(serv users.Services, cfg *cfg.Config, sugar *zap.SugaredLogger) users.Controller {
	return &userController{serv: serv, cfg: cfg, sugar: sugar}
}

func (uC *userController) RegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := &entities.User{}
		if err := utils.ReadRequest(r, u); err != nil {
			uC.sugar.Errorf("error while reading request: %v", err.Error())
			utils.Respond(w, 500, err.Error())
			return
		}

		utils.Respond(w, 200, "oi")
	}
}
