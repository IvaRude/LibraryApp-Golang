package postgresql

import (
	"context"
	"homework-3/internal/pkg/db"
	"homework-3/internal/pkg/repository"
)

type BookRepo struct {
	db *db.Database
}

func NewBooks(database *db.Database) *BookRepo {
	return &BookRepo{db: database}
}

func (r *BookRepo) Add(ctx context.Context, book *repository.Book) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, `SELECT id FROM authors WHERE id =$1;`, book.AuthorId).Scan(&id)
	if err != nil {
		return 0, repository.ErrObjectNotFound
	}
	err = r.db.ExecQueryRow(ctx, `INSERT INTO books(name, author_id) VALUES($1, $2) RETURNING id;`, book.Name, book.AuthorId).Scan(&id)
	return id, err
}

func (r *BookRepo) GetByID(ctx context.Context, id int64) (*repository.Book, error) {
	var a repository.Book
	err := r.db.Get(ctx, &a, "SELECT id, name, author_id FROM books WHERE id=$1;", id)
	if err != nil {
		return nil, repository.ErrObjectNotFound
	}
	return &a, nil
}

func (r *BookRepo) Update(ctx context.Context, author *repository.Book) error {
	var id int64
	err := r.db.ExecQueryRow(ctx, `UPDATE authors SET name = $1 WHERE id = $2 RETURNING id;`, author.Name, author.Id).Scan(&id)
	if err != nil {
		return repository.ErrObjectNotFound
	}
	return nil
}

func (r *BookRepo) DeleteById(ctx context.Context, id int64) error {
	var deletedId int64
	err := r.db.ExecQueryRow(ctx, `DELETE FROM authors WHERE id = $1 RETURNING id;`, id).Scan(&deletedId)
	if err != nil {
		return repository.ErrObjectNotFound
	}
	return nil
}
