package repo

import (
	"github.com/gianlucapastori/nausicaa/internal/entities"
	"github.com/gianlucapastori/nausicaa/internal/packages/groups"
	"github.com/jmoiron/sqlx"
)

type groupsRepo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) groups.Repo {
	return &groupsRepo{db: db}
}

// InsertGroup implements groups.Repo
func (gR *groupsRepo) InsertGroup(group *entities.Group) error {
	INSERT_GROUP_SQL := `INSERT INTO groups(name,description,creator_id) VALUES ($1,$2,$3) `

	_, err := gR.db.Queryx(INSERT_GROUP_SQL,
		group.Name,
		group.Description,
		group.CreatorId,
	)
	if err != nil {
		return err
	}

	return nil
}

// FetchGroupById implements groups.Repo
func (gR *groupsRepo) FetchGroupById(id string) (*entities.Group, error) {
	SELECT_GROUP_BY_ID_SQL := `SELECT * FROM groups WHERE id = $1`

	row, err := gR.db.Queryx(SELECT_GROUP_BY_ID_SQL, id)
	if err != nil {
		return nil, err
	}

	g := &entities.Group{}
	defer row.Close()

	for row.Next() {
		if err := row.Scan(&g.Id, &g.Name, &g.Description, &g.CreatorId); err != nil {
			return nil, err
		}
	}

	return g, nil
}
