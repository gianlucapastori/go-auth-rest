package users

import "github.com/gianlucapastori/nausicaa/internal/entities"

type Repo interface {
	InsertUser(*entities.User) (*entities.User, error)
	FetchUserByEmail(string) (*entities.User, error)
	FetchUserByUsername(string) (*entities.User, error)
}
