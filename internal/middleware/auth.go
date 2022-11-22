package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gianlucapastori/go-auth-jwt/cfg"
	"github.com/gianlucapastori/go-auth-jwt/pkg/jwt"
	"github.com/gianlucapastori/go-auth-jwt/pkg/utils"
	jwtgo "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/context"
)

func (mw *Middleware) AuthJWT(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		acctoken, err := jwt.ExtractBearer(r)
		if err != nil {
			utils.Respond(w, http.StatusInternalServerError, err.Error())
			return
		}

		if token, err := validateAccessToken(acctoken, mw.cfg); err != nil {
			if err.Error() == "invalid jwt token" {
				utils.Respond(w, 403, err.Error())
			}
			utils.Respond(w, http.StatusInternalServerError, err.Error())
			return
		} else {
			context.Set(r, "user_id", token.Claims.(jwtgo.MapClaims)["user_id"].(string))
		}

		h.ServeHTTP(w, r)
	})
}

func (mw *Middleware) AuthRefresh(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		refcookie, err := r.Cookie(mw.cfg.SERVER.JWT.REFRESH_COOKIE_NAME)
		if err != nil {
			utils.Respond(w, http.StatusInternalServerError, err.Error())
			return
		}

		reftoken := refcookie.Value

		if err := validateRefreshToken(reftoken, mw.cfg); err != nil {
			if err.Error() == "invalid jwt token" {
				utils.Respond(w, 403, err.Error())
			}
			utils.Respond(w, http.StatusInternalServerError, err.Error())
			return
		}

		h.ServeHTTP(w, r)
	})
}

func (mw *Middleware) AuthPwd(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type Request struct {
			Token                   string `json:"token" validate:"required"`
			Email                   string `json:"email" validate:"required,email"`
			NewPassword             string `json:"new_password" validate:"gte=5,required"`
			NewPasswordConfirmation string `json:"new_password_confirmation" validate:"gte=5,required"`
		}

		req := &Request{}

		if err := utils.ReadRequest(r, req); err != nil {
			mw.sugar.Errorf("error while reading request: %v", err.Error())
			utils.Respond(w, http.StatusInternalServerError, err.Error())
			return
		}

		u, err := mw.serv.FetchByEmail(req.Email)
		if err != nil {
			mw.sugar.Errorf(err.Error())
			utils.Respond(w, 500, err.Error())
			return
		}

		if err := validatePwdToken(req.Token, u.Password, mw.cfg); err != nil {
			if err.Error() == "invalid jwt token" {
				utils.Respond(w, 403, err.Error())
			}
			utils.Respond(w, http.StatusInternalServerError, err.Error())
			return
		}

		context.Set(r, "token", req.Token)
		context.Set(r, "email", req.Email)
		context.Set(r, "pwd", req.NewPassword)
		context.Set(r, "pwd_conf", req.NewPasswordConfirmation)
		h.ServeHTTP(w, r)
	})

}

func validateAccessToken(tokenStr string, cfg *cfg.Config) (*jwtgo.Token, error) {
	if tokenStr == "" {
		return nil, errors.New("invalid jwt token")
	}

	token, err := jwtgo.Parse(tokenStr, func(token *jwtgo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtgo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secret := []byte(cfg.SERVER.JWT.SECRET_KEY_ACCESS)
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid jwt token")
	}

	if claims, ok := token.Claims.(jwtgo.MapClaims); ok && token.Valid {
		id := claims["user_id"].(string)

		if claims["user_id"] == nil || id == "" {
			return nil, errors.New("could not get claims properties")
		}
	} else {
		return nil, errors.New("invalid claims")
	}

	return token, nil
}

func validatePwdToken(tokenStr string, db_hash string, cfg *cfg.Config) error {
	if tokenStr == "" {
		return errors.New("invalid jwt token")
	}

	token, err := jwtgo.Parse(tokenStr, func(token *jwtgo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtgo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.SERVER.JWT.SECRET_KEY_ACCESS), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("invalid jwt token")
	}

	if token.Claims.(jwtgo.MapClaims)["hash"] == nil {
		return errors.New("invalid claims")
	}

	if token.Claims.(jwtgo.MapClaims)["hash"] != db_hash {
		return errors.New("invalid jwt token")
	}

	return nil
}

func validateRefreshToken(tokenStr string, cfg *cfg.Config) error {
	if tokenStr == "" {
		return errors.New("invalid jwt token")
	}

	token, err := jwtgo.Parse(tokenStr, func(token *jwtgo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtgo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.SERVER.JWT.SECRET_KEY_REFRESH), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("invalid jwt token")
	}

	return nil
}
