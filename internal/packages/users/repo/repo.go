package repo

import (
	"github.com/gianlucapastori/nausicaa/internal/entities"
	"github.com/gianlucapastori/nausicaa/internal/packages/users"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) users.Repo {
	return &userRepo{db: db}
}

func (uR *userRepo) InsertUser(user *entities.User) error {
	panic("")
}

func (uR *userRepo) FetchUserByEmail(email string) (*entities.User, error) {
	panic("")
}

func (uR *userRepo) FetchUserByUsername(username string) (*entities.User, error) {
	panic("")
}
