package postgresql

import (
	"context"
	"database/sql"

	"homework-3/internal/pkg/db"
	"homework-3/internal/pkg/repository"
)

type AuthorRepo struct {
	db *db.Database
}

func NewAuthors(database *db.Database) *AuthorRepo {
	return &AuthorRepo{db: database}
}

func (r *AuthorRepo) Add(ctx context.Context, author *repository.Author) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, `INSERT INTO authors(name) VALUES($1) RETURNING id;`, author.Name).Scan(&id)
	return id, err
}

func (r *AuthorRepo) GetByID(ctx context.Context, id int64) (*repository.Author, error) {
	var a repository.Author
	err := r.db.Get(ctx, &a, "SELECT id,name FROM authors WHERE id=$1", id)
	if err == sql.ErrNoRows {
		return nil, repository.ErrObjectNotFound
	}
	return &a, nil
}
