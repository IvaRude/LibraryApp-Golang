package app

import (
	"context"
	"encoding/json"
	"errors"
	"homework-3/internal/pkg/models"
	"homework-3/internal/pkg/repository"
	"homework-3/internal/pkg/routers"
	"log"
	"net/http"
)

func (a *App) CreateBook(ctx context.Context, updateBookData *routers.UpdateBookRequest) routers.StatusInt {
	bookRepo := &models.Book{
		Name:     updateBookData.Name,
		AuthorId: updateBookData.AuthorId,
	}
	_, err := a.BookRepo.Add(ctx, bookRepo)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return http.StatusNotFound
		}
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

func (a *App) GetBook(ctx context.Context, id int64) ([]byte, routers.StatusInt) {
	book, err := a.BookRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return nil, http.StatusNotFound
		}
		log.Print(err)
		return nil, http.StatusInternalServerError
	}
	bookJson, err := json.Marshal(book)
	if err != nil {
		log.Print(err)
		return nil, http.StatusInternalServerError
	}
	return bookJson, http.StatusOK
}
