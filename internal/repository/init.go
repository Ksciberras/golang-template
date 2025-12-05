package repository

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"

	schema "auth-node/internal/repository/database"
)

type Repo struct {
	PGSQL *PGSQL
}

type PGSQL struct {
	Db      *sql.DB
	Queries *schema.Queries
}

func InitRepo(cfg *config.Config) (*Repo, error) {
	db, err := sql.Open("postgres", cfg.Database.URL)
	if err != nil {
		return nil, err
	}

	queries := schema.New(db)

	return &Repo{
		PGSQL: &PGSQL{
			Db:      db,
			Queries: queries,
		},
	}, nil
}
