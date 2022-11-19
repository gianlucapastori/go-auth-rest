package jwt

import (
	"math/rand"
	"time"

	"github.com/gianlucapastori/nausicaa/cfg"
	"github.com/gianlucapastori/nausicaa/internal/entities"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type UserClaims struct {
	UserId   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	jwt.RegisteredClaims
}

func RequestTokens(user *entities.User, cfg *cfg.Config) (string, string, error) {
	accclaims := UserClaims{
		UserId:   user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(rand.Int31n(cfg.SERVER.JWT.ACCESS_EXPIRES_AT)) * time.Minute)),
		},
	}

	refclaims := UserClaims{
		UserId:   user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(rand.Int31n(cfg.SERVER.JWT.REFRESH_EXPIRES_AT)) * time.Minute)),
		},
	}

	acctoken := jwt.NewWithClaims(jwt.SigningMethodHS256, accclaims)
	reftoken := jwt.NewWithClaims(jwt.SigningMethodHS256, refclaims)
	accTokenStr, err := acctoken.SignedString([]byte(cfg.SERVER.JWT.SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	refTokenStr, err := reftoken.SignedString([]byte(cfg.SERVER.JWT.SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	return accTokenStr, refTokenStr, nil
}
