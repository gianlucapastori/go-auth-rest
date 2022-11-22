package groups

import "github.com/gianlucapastori/nausicaa/internal/entities"

type Services interface {
	CreateGroup(*entities.Group) error
	FetchGroupById(*entities.Group) (*entities.Group, error)
}
