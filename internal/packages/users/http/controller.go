package http

import (
	"fmt"
	"net/http"

	"github.com/gianlucapastori/nausicaa/cfg"
	"github.com/gianlucapastori/nausicaa/internal/entities"
	"github.com/gianlucapastori/nausicaa/internal/packages/users"
	"github.com/gianlucapastori/nausicaa/pkg/jwt"
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

		u, err := uC.serv.Register(&entities.User{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Username:  req.Username,
			Email:     req.Email,
			Password:  req.Password,
		})

		if err != nil {
			if err.Error() == "email already in use" || err.Error() == "username already in use" {
				utils.Respond(w, http.StatusBadRequest, err.Error())
				return
			}
			utils.Respond(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := uC.SendTokensToCookie(w, r, u); err != nil {
			utils.Respond(w, http.StatusInternalServerError, err.Error())
		}

		utils.Respond(w, http.StatusOK, fmt.Sprintf("user %s created", u.Username))
	}
}

func (uC *userController) Protected() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Respond(w, http.StatusOK, "access granted!")
	}
}

func (uC *userController) RequestNewAccess() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if acc, _, err := jwt.RequestTokens(&entities.User{}, uC.cfg); err != nil {
			utils.Respond(w, 500, err.Error())
		} else {
			acccookie := &http.Cookie{
				Name:  uC.cfg.SERVER.JWT.ACCESS_COOKIE_NAME,
				Value: acc,
				Path:  "/",
			}

			http.SetCookie(w, acccookie)
			r.AddCookie(acccookie)
		}

		utils.Respond(w, http.StatusOK, "")
	}
}

func (uC *userController) SendTokensToCookie(w http.ResponseWriter, r *http.Request, user *entities.User) error {
	if acc, ref, err := jwt.RequestTokens(user, uC.cfg); err != nil {
		return err
	} else {
		acccookie := &http.Cookie{
			Name:  uC.cfg.SERVER.JWT.ACCESS_COOKIE_NAME,
			Value: acc,
			Path:  "/",
		}

		refcookie := &http.Cookie{
			Name:     uC.cfg.SERVER.JWT.REFRESH_COOKIE_NAME,
			Value:    ref,
			HttpOnly: true,
			Path:     "/",
		}

		http.SetCookie(w, acccookie)
		http.SetCookie(w, refcookie)
		r.AddCookie(acccookie)
		r.AddCookie(refcookie)
		return nil
	}
}
