package postgresql

import (
	"database/sql"
	"context"
	"errors"
	
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	
	"tradier-fiber/internals"
	"tradier-fiber/internals/postgresql/db"
)


Type Repo struct {
	q *db.Queries
}

func NewRepo(d db.DBTX) *Repo {
	return &Repo{
		q: db.New(d)
	}	
}