package repo

import (
	"github.com/gianlucapastori/nausicaa/internal/packages/groups"
	"github.com/jmoiron/sqlx"
)

type groupsRepo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) groups.Repo {
	return &groupsRepo{db: db}
}
