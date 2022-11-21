package services

import (
	"github.com/gianlucapastori/nausicaa/cfg"
	"github.com/gianlucapastori/nausicaa/internal/packages/groups"
	"go.uber.org/zap"
)

type groupsServices struct {
	repo  groups.Repo
	sugar *zap.SugaredLogger
	cfg   *cfg.Config
}

func New(repo groups.Repo, sugar *zap.SugaredLogger, cfg *cfg.Config) groups.Services {
	return &groupsServices{repo: repo, sugar: sugar, cfg: cfg}
}
