package postgresql

import (
	"context"
	"homework-3/internal/pkg/db"
	"homework-3/internal/pkg/models"
	"homework-3/internal/pkg/repository"
)

type BookRepo struct {
	db db.DBops
}

func NewBooks(database db.DBops) *BookRepo {
	return &BookRepo{db: database}
}

func (r *BookRepo) Add(ctx context.Context, book *models.Book) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, `SELECT id FROM authors WHERE id =$1;`, book.AuthorId).Scan(&id)
	if err != nil {
		return 0, repository.ErrObjectNotFound
	}
	err = r.db.ExecQueryRow(ctx, `INSERT INTO books(name, author_id) VALUES($1, $2) RETURNING id;`, book.Name, book.AuthorId).Scan(&id)
	return id, err
}

func (r *BookRepo) GetByID(ctx context.Context, id int64) (*models.Book, error) {
	var a models.Book
	err := r.db.Get(ctx, &a, "SELECT id, name, author_id FROM books WHERE id=$1;", id)
	if err != nil {
		return nil, repository.ErrObjectNotFound
	}
	return &a, nil
}
