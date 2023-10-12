package routers

import (
	"encoding/json"
	"errors"
	"fmt"
	"homework-3/internal/pkg/repository"
	"homework-3/internal/pkg/server"
	"io"
	"log"
	"net/http"
	"strconv"

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

func CreateAuthorRouter(router *mux.Router, s server.Server) *mux.Router {
	router.HandleFunc("/author", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			CreateAuthor(s, w, req)
		case http.MethodPut:
			UpdateAuthor(s, w, req)
		default:
			fmt.Println("error")
		}
	})

	router.HandleFunc(fmt.Sprintf("/author/{%s:[0-9]*}", queryParamKey), func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			GetAuthor(s, w, req)
		case http.MethodDelete:
			DeleteAuthor(s, w, req)
		default:
			fmt.Println("error")
		}
	})
	return router
}

func CreateAuthor(s server.Server, w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var unm addAuthorRequest
	if err = json.Unmarshal(body, &unm); err != nil {
		AnswerError(w, http.StatusInternalServerError)
		return
	}
	authorRepo := &repository.Author{
		Name: unm.Name,
	}
	_, err = s.AuthorRepo.Add(req.Context(), authorRepo)
	if err != nil {
		AnswerError(w, http.StatusInternalServerError)
		return
	}
}

func GetAuthor(s server.Server, w http.ResponseWriter, req *http.Request) {
	key, ok := mux.Vars(req)[queryParamKey]
	if !ok {
		AnswerError(w, http.StatusBadRequest)
		return
	}
	keyInt, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		AnswerError(w, http.StatusBadRequest)
		return
	}
	author, err := s.AuthorRepo.GetByID(req.Context(), keyInt)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			AnswerError(w, http.StatusNotFound)
			return
		}
		log.Fatal(err)
		AnswerError(w, http.StatusInternalServerError)
		return
	}
	authorJson, err := json.Marshal(author)
	if err != nil {
		log.Fatal(err)
		AnswerError(w, http.StatusInternalServerError)
		return
	}
	w.Write(authorJson)
}

func UpdateAuthor(s server.Server, w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		AnswerError(w, http.StatusInternalServerError)
		return
	}
	var unm updateAuthorRequest
	if err = json.Unmarshal(body, &unm); err != nil {
		AnswerError(w, http.StatusInternalServerError)
		return
	}
	authorRepo := &repository.Author{
		Name: unm.Name,
		Id:   unm.Id,
	}
	err = s.AuthorRepo.Update(req.Context(), authorRepo)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			AnswerError(w, http.StatusNotFound)
			return
		}
		AnswerError(w, http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
}
func DeleteAuthor(s server.Server, w http.ResponseWriter, req *http.Request) {
	key, ok := mux.Vars(req)[queryParamKey]
	if !ok {
		AnswerError(w, http.StatusBadRequest)
		return
	}
	keyInt, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		AnswerError(w, http.StatusBadRequest)
		return
	}
	err = s.AuthorRepo.DeleteById(req.Context(), keyInt)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			AnswerError(w, http.StatusNotFound)
			return
		}
		AnswerError(w, http.StatusInternalServerError)
		return
	}
}
