package routers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"homework-3/internal/pkg/repository"
	"homework-3/internal/pkg/server"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const queryParamKey = "key"

type addAuthorRequest struct {
	Name string `json:"name"`
}

type updateAuthorRequest struct {
	addAuthorRequest
	Id int64 `json:"id"`
}

func CreateAuthorRouter(router *mux.Router, s *server.Server) *mux.Router {
	router.HandleFunc("/author", func(w http.ResponseWriter, req *http.Request) {
		updateAuthorData, status := parseUpdateAuthorRequest(req)
		if status != http.StatusOK {
			AnswerError(w, status)
			return
		}
		switch req.Method {
		case http.MethodPost:
			if status = CreateAuthor(req.Context(), s, updateAuthorData); status != http.StatusOK {
				AnswerError(w, status)
			} else {
				w.WriteHeader(int(status))
			}
		case http.MethodPut:
			if status = UpdateAuthor(req.Context(), s, updateAuthorData); status != http.StatusOK {
				AnswerError(w, status)
			} else {
				w.WriteHeader(int(status))
			}
		default:
			fmt.Println("error")
		}
	})

	router.HandleFunc(fmt.Sprintf("/author/{%s:[0-9]*}", queryParamKey), func(w http.ResponseWriter, req *http.Request) {
		id, status := ParseID(req)
		if status != http.StatusOK {
			AnswerError(w, status)
			return
		}
		switch req.Method {
		case http.MethodGet:
			authorJson, status := GetAuthor(req.Context(), s, id)
			if status != http.StatusOK {
				AnswerError(w, status)
			} else {
				w.WriteHeader(int(status))
				w.Write(authorJson)
			}
		case http.MethodDelete:
			if status = DeleteAuthor(req.Context(), s, id); status != http.StatusOK {
				AnswerError(w, status)
			} else {
				w.WriteHeader(int(status))
			}
		default:
			fmt.Println("error")
		}
	})
	return router
}

func CreateAuthor(ctx context.Context, s *server.Server, updateAuthorData *updateAuthorRequest) statusInt {
	authorRepo := &repository.Author{
		Name: updateAuthorData.Name,
	}
	_, err := s.AuthorRepo.Add(ctx, authorRepo)
	if err != nil {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

func GetAuthor(ctx context.Context, s *server.Server, id int64) ([]byte, statusInt) {
	author, err := s.AuthorRepo.GetByID(ctx, id)
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

func UpdateAuthor(ctx context.Context, s *server.Server, updateAuthorData *updateAuthorRequest) statusInt {
	authorRepo := &repository.Author{
		Name: updateAuthorData.Name,
		Id:   updateAuthorData.Id,
	}
	if err := s.AuthorRepo.Update(ctx, authorRepo); err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return http.StatusNotFound
		}
		log.Print(err)
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

func DeleteAuthor(ctx context.Context, s *server.Server, id int64) statusInt {
	err := s.AuthorRepo.DeleteById(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return http.StatusNotFound
		}
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

func parseUpdateAuthorRequest(req *http.Request) (*updateAuthorRequest, statusInt) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, http.StatusBadRequest
	}
	var unm updateAuthorRequest
	if err = json.Unmarshal(body, &unm); err != nil {
		return nil, http.StatusBadRequest
	}
	return &unm, http.StatusOK
}
