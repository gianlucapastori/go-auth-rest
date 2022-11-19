package services

import (
	"github.com/gianlucapastori/nausicaa/cfg"
	"github.com/gianlucapastori/nausicaa/internal/entities"
	"github.com/gianlucapastori/nausicaa/internal/packages/users"
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

func (*userService) Register(user *entities.User) (*entities.User, error) {
	panic("unimplemented")
}

func (uS *userService) FetchByEmail(email string) (*entities.User, error) {
	return uS.repo.FetchUserByEmail(email)
}

func (uS *userService) FetchByUsername(username string) (*entities.User, error) {
	return uS.repo.FetchUserByUsername(username)
}
