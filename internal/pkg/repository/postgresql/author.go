package postgresql

import (
	"context"

	"homework-3/internal/pkg/db"
	"homework-3/internal/pkg/repository"
)

type AuthorRepo struct {
	db db.DBops
}

func NewAuthors(database db.DBops) *AuthorRepo {
	return &AuthorRepo{db: database}
}

func (r *AuthorRepo) Add(ctx context.Context, author *repository.Author) (int64, error) {
	var id int64
	err := r.db.Get(ctx, &id, `INSERT INTO authors(name) VALUES($1) RETURNING id;`, author.Name)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (r *AuthorRepo) GetByID(ctx context.Context, id int64) (*repository.Author, error) {
	var a repository.Author
	err := r.db.Get(ctx, &a, "SELECT id,name FROM authors WHERE id=$1;", id)
	if err != nil {
		return nil, repository.ErrObjectNotFound
	}
	var books []repository.Book
	err = r.db.Select(ctx, &books, "SELECT id, name, author_id FROM books WHERE author_id=$1;", id)
	a.Books = books
	return &a, nil
}

func (r *AuthorRepo) Update(ctx context.Context, author *repository.Author) error {
	var id int64
	err := r.db.Get(ctx, &id, `UPDATE authors SET name = $1 WHERE id = $2 RETURNING id;`, author.Name, author.Id)
	if err != nil {
		return repository.ErrObjectNotFound
	}
	return nil
}

func (r *AuthorRepo) DeleteById(ctx context.Context, id int64) error {
	var deletedId int64
	err := r.db.Get(ctx, &deletedId, `WITH q AS (DELETE FROM books WHERE author_id = $1) DELETE FROM authors WHERE id = $1 RETURNING id;`, id)
	if err != nil {
		return repository.ErrObjectNotFound
	}
	return nil
}
