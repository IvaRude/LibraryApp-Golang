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

func (a *App) CreateAuthor(ctx context.Context, updateAuthorData *routers.UpdateAuthorRequest) routers.StatusInt {
	authorRepo := &models.Author{
		Name: updateAuthorData.Name,
	}
	_, err := a.AuthorRepo.Add(ctx, authorRepo)
	if err != nil {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

func (a *App) GetAuthor(ctx context.Context, id int64) ([]byte, routers.StatusInt) {
	author, err := a.AuthorRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return nil, http.StatusNotFound
		}
		log.Print(err)
		return nil, http.StatusInternalServerError
	}
	authorJson, err := json.Marshal(author)
	if err != nil {
		log.Print(err)
		return nil, http.StatusInternalServerError
	}
	return authorJson, http.StatusOK
}

func (a *App) UpdateAuthor(ctx context.Context, updateAuthorData *routers.UpdateAuthorRequest) routers.StatusInt {
	authorRepo := &models.Author{
		Name: updateAuthorData.Name,
		Id:   updateAuthorData.Id,
	}
	if err := a.AuthorRepo.Update(ctx, authorRepo); err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return http.StatusNotFound
		}
		log.Print(err)
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

func (a *App) DeleteAuthor(ctx context.Context, id int64) routers.StatusInt {
	err := a.AuthorRepo.DeleteById(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return http.StatusNotFound
		}
		return http.StatusInternalServerError
	}
	return http.StatusOK
}
