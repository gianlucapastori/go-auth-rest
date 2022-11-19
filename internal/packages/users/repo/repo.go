package repo

import (
	"errors"

	"github.com/gianlucapastori/nausicaa/internal/entities"
	"github.com/gianlucapastori/nausicaa/internal/packages/users"
	"github.com/google/uuid"
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
	SELECT_USER_BY_EMAIL_SQL := `SELECT * FROM users WHERE email = $1`

	row, err := uR.db.Queryx(SELECT_USER_BY_EMAIL_SQL, email)
	if err != nil {
		return nil, err
	}

	u := &entities.User{}
	defer row.Close()

	for row.Next() {
		if err := row.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Username, &u.Email, &u.Password); err != nil {
			return nil, err
		}
	}

	if u.Id != uuid.Nil {
		return nil, errors.New("email already in use")
	}

	return u, nil
}

func (uR *userRepo) FetchUserByUsername(username string) (*entities.User, error) {
	SELECT_USER_BY_USERNAME_SQL := `SELECT * FROM users WHERE username = $1`

	row, err := uR.db.Queryx(SELECT_USER_BY_USERNAME_SQL, username)
	if err != nil {
		return nil, err
	}

	u := &entities.User{}
	defer row.Close()

	for row.Next() {
		if err := row.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Username, &u.Email, &u.Password); err != nil {
			return nil, err
		}
	}

	if u.Id != uuid.Nil {
		return nil, errors.New("username already in use")
	}

	return u, nil
}
