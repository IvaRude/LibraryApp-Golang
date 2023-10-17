//go:generate mockgen -source ./repository.go -destination=./mocks/repository.go -package=mock_repository
package repository

import "context"

type AuthorsRepo interface {
	Add(ctx context.Context, author *Author) (int64, error)
	GetByID(ctx context.Context, id int64) (*Author, error)
	Update(ctx context.Context, author *Author) error
	DeleteById(ctx context.Context, id int64) error
}

type BooksRepo interface {
	Add(ctx context.Context, book *Book) (int64, error)
	GetByID(ctx context.Context, id int64) (*Book, error)
}
