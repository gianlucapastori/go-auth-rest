package users

import "github.com/gianlucapastori/go-auth-jwt/internal/entities"

type Repo interface {
	InsertUser(*entities.User) (*entities.User, error)
	ChangeUserPasswordByEmail(string, string) error
	FetchUserByEmail(string) (*entities.User, error)
	FetchUserByUsername(string) (*entities.User, error)
}
