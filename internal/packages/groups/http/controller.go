package http

import (
	"fmt"
	"net/http"

	"github.com/gianlucapastori/nausicaa/cfg"
	"github.com/gianlucapastori/nausicaa/internal/entities"
	"github.com/gianlucapastori/nausicaa/internal/packages/groups"
	"github.com/gianlucapastori/nausicaa/pkg/utils"
	"github.com/gorilla/context"
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

// CreateGroup implements groups.Controller
func (gC *groupsController) CreateGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		g := &entities.Group{}

		if err := utils.ReadRequest(r, g); err != nil {
			gC.sugar.Errorf("error while reading request: %v", err.Error())
			utils.Respond(w, http.StatusInternalServerError, err.Error())
			return
		}

		if context.Get(r, "user_id") == nil {
			gC.sugar.Errorf("malformed request")
			utils.Respond(w, http.StatusInternalServerError, "malformed request")
			return
		}

		user_id := context.Get(r, "user_id")

		g.CreatorId = user_id.(string)

		if err := gC.serv.CreateGroup(g); err != nil {
			utils.Respond(w, 500, err.Error())
			return
		}

		utils.Respond(w, 200, fmt.Sprintf("group %s created!", g.Name))
	}
}

// CreateTask implements groups.Controller
func (*groupsController) CreateTask() http.HandlerFunc {
	panic("unimplemented")
}

// RemoveGroup implements groups.Controller
func (*groupsController) RemoveGroup() http.HandlerFunc {
	panic("unimplemented")
}

// RemoveTask implements groups.Controller
func (*groupsController) RemoveTask() http.HandlerFunc {
	panic("unimplemented")
}

// UpdateGroup implements groups.Controller
func (*groupsController) UpdateGroup() http.HandlerFunc {
	panic("unimplemented")
}
