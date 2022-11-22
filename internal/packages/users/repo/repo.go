package repo

import (
	"github.com/gianlucapastori/go-auth-jwt/internal/entities"
	"github.com/gianlucapastori/go-auth-jwt/internal/packages/users"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) users.Repo {
	return &userRepo{db: db}
}

func (uR *userRepo) InsertUser(user *entities.User) (*entities.User, error) {
	INSERT_USER_SQL := `INSERT INTO users(first_name,last_name,username,email,password) VALUES ($1,$2,$3,$4,$5) `

	_, err := uR.db.Queryx(INSERT_USER_SQL,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Password,
	)
	if err != nil {
		return nil, err
	}

	u, err := uR.FetchUserByUsername(user.Username)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uR *userRepo) ChangeUserPasswordByEmail(email string, hash string) error {
	CHANGE_USER_PASSWORD_BY_EMAIL_SQL := `UPDATE users SET password = $1 WHERE email = $2`

	_, err := uR.db.Queryx(CHANGE_USER_PASSWORD_BY_EMAIL_SQL, hash, email)
	if err != nil {
		return err
	}

	return nil
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

	return u, nil
}
