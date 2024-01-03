//go:generate mockgen -source ./repository.go -destination=./mocks/repository/repository.go -package=mock_repository
package app

import (
	"context"
	"homework-3/internal/pkg/models"
)

type AuthorsRepo interface {
	Add(ctx context.Context, author *models.Author) (int64, error)
	GetByID(ctx context.Context, id int64) (*models.Author, error)
	Update(ctx context.Context, author *models.Author) error
	DeleteById(ctx context.Context, id int64) error
}

type BooksRepo interface {
	Add(ctx context.Context, book *models.Book) (int64, error)
	GetByID(ctx context.Context, id int64) (*models.Book, error)
}
