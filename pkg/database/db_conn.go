package database

import (
	"fmt"

	"github.com/gianlucapastori/go-auth-jwt/cfg"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(cfg *cfg.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("user=%s host=%s dbname=%s password=%s port=%s sslmode=disable",
			cfg.DATABASE.USER,
			cfg.DATABASE.HOST,
			cfg.DATABASE.NAME,
			cfg.DATABASE.PSWD,
			cfg.DATABASE.PORT,
		))
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping database: %v", err)
	}

	return db, nil
}
