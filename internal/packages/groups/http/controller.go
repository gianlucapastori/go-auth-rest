package http

import (
	"github.com/gianlucapastori/nausicaa/cfg"
	"github.com/gianlucapastori/nausicaa/internal/packages/groups"
	"go.uber.org/zap"
)

type groupsController struct {
	serv  groups.Services
	cfg   *cfg.Config
	sugar *zap.SugaredLogger
}

func New(serv groups.Services, cfg *cfg.Config, sugar *zap.SugaredLogger) groups.Controller {
	return &groupsController{serv: serv, cfg: cfg, sugar: sugar}
}
