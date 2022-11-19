package services

import (
	"errors"

	"github.com/gianlucapastori/nausicaa/cfg"
	"github.com/gianlucapastori/nausicaa/internal/entities"
	"github.com/gianlucapastori/nausicaa/internal/packages/users"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type userService struct {
	repo  users.Repo
	cfg   *cfg.Config
	sugar *zap.SugaredLogger
}

func New(repo users.Repo, cfg *cfg.Config, sugar *zap.SugaredLogger) users.Services {
	return &userService{repo: repo, cfg: cfg, sugar: sugar}
}

func (uS *userService) Register(user *entities.User) (*entities.User, error) {
	if err := uS.isEmailAvailable(user.Email); err != nil {
		return nil, err
	}

	if err := uS.isUsernameAvailable(user.Username); err != nil {
		return nil, err
	}

	return nil, nil
}

func (uS *userService) FetchByEmail(email string) (*entities.User, error) {
	return uS.repo.FetchUserByEmail(email)
}

func (uS *userService) FetchByUsername(username string) (*entities.User, error) {
	return uS.repo.FetchUserByUsername(username)
}

func (uS *userService) isEmailAvailable(email string) error {
	u, err := uS.FetchByEmail(email)
	if err != nil {
		uS.sugar.Errorf(err.Error())
		return err
	}
	if u.Id != uuid.Nil {
		return errors.New("email already in use")
	}
	return nil
}

func (uS *userService) isUsernameAvailable(username string) error {
	u, err := uS.FetchByUsername(username)
	if err != nil {
		uS.sugar.Errorf(err.Error())
		return err
	}
	if u.Id != uuid.Nil {
		return errors.New("username already in use")
	}
	return nil
}
