package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gianlucapastori/nausicaa/cfg"
	"github.com/gianlucapastori/nausicaa/pkg/jwt"
	"github.com/gianlucapastori/nausicaa/pkg/utils"
	jwtgo "github.com/golang-jwt/jwt/v4"
)

func (mw *Middleware) AuthJWT(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		acctoken, err := jwt.ExtractBearer(r)
		if err != nil {
			utils.Respond(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := validateAccessToken(acctoken, mw.cfg); err != nil {
			if err.Error() == "invalid jwt token" {
				utils.Respond(w, 403, err.Error())
			}
			utils.Respond(w, http.StatusInternalServerError, err.Error())
			return
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

func validateAccessToken(tokenStr string, cfg *cfg.Config) error {
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
