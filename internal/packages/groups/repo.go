package groups

import "github.com/gianlucapastori/nausicaa/internal/entities"

type Repo interface {
	InsertGroup(*entities.Group) error
	FetchGroupById(id string) (*entities.Group, error)
}
