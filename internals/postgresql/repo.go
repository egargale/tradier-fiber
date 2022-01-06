package postgresql

import (
	"database/sql"
	// "context"
	// "errors"

	// "github.com/google/uuid"
	// "github.com/jackc/pgx/v4"

	// "github.com/egargale/tradier-fiber/internals"
	"github.com/egargale/tradier-fiber/internals/postgresql/db"
)

type Repo struct {
	q *db.Queries
	DB *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		DB: db,
		q: db.New(db),}
}
