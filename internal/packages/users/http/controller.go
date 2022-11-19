package http

import (
	"fmt"
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
		type Request struct {
			FirstName string `json:"first_name,omitempty" validate:"omitempty,lte=30,gte=2,alpha"`
			LastName  string `json:"last_name,omitempty" validate:"omitempty,lte=30,gte=2,alpha"`
			Username  string `json:"username" validate:"lte=30,gte=2,required"`
			Email     string `json:"email" validate:"email,required"`
			Password  string `json:"password" validate:"gte=5,required"`
			PwdConf   string `json:"password_confirmation" validate:"gte=5,required"`
		}

		req := &Request{}

		if err := utils.ReadRequest(r, req); err != nil {
			uC.sugar.Errorf("error while reading request: %v", err.Error())
			utils.Respond(w, http.StatusInternalServerError, err.Error())
			return
		}

		if _, err := uC.serv.Register(&entities.User{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Username:  req.Username,
			Email:     req.Email,
			Password:  req.Password,
		}); err != nil {
			if err.Error() == "email already in use" || err.Error() == "username already in use" {
				utils.Respond(w, http.StatusBadRequest, err.Error())
				return
			}
			utils.Respond(w, 200, err.Error())
			return
		}

		utils.Respond(w, http.StatusOK, fmt.Sprintf("user %s created", req.Username))
	}
}
