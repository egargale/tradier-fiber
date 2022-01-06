package postgresql

import (
	"context"
	"errors"

	// "github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	// "github.com/egargale/tradier-fiber/internals"
	"github.com/egargale/tradier-fiber/internals/postgresql/db"
)

type Repo struct {
	q *db.Queries
}

func NewRepo(d db.DBTX) *Repo {
	return &Repo{
		q: db.New(d)}
}

func (r *Repo) Find(ctx context.Context) ([]db.Todo, error) {
	res, err := r.q.GetAllTodos(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}
	}
	return res, err
}
