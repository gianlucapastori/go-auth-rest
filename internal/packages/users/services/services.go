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

func (*userService) FetchByEmail(email string) (*entities.User, error) {
	panic("unimplemented")
}

func (*userService) FetchByUsername(username string) (*entities.User, error) {
	panic("unimplemented")
}
