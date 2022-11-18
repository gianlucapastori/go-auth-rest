package users

import "github.com/gianlucapastori/nausicaa/internal/entities"

type Services interface {
	Register(*entities.User) (*entities.User, error)
	FetchByEmail(string) (*entities.User, error)
	FetchByUsername(string) (*entities.User, error)
}
