package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"homework-3/internal/pkg/db"
	"homework-3/internal/pkg/repository"
	"homework-3/internal/pkg/repository/postgresql"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const port = ":9000"
const queryParamKey = "key"
const (
	errorNotFound    = "Not Found"
	errorBadRequest  = "Bad Request"
	errorServerError = "Server Error"
)

type server struct {
	authorRepo *postgresql.AuthorRepo
}

type addAuthorRequest struct {
	Name string `json:"name"`
}

type updateAuthorRequest struct {
	addAuthorRequest
	Id int64 `json:"id"`
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	database, err := db.NewDB(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer database.GetPool(ctx).Close()

	authorRepo := postgresql.NewAuthors(database)

	implemetation := server{authorRepo: authorRepo}
	http.Handle("/", createRouter(implemetation))
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func createRouter(implemetation server) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/author", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			implemetation.CreateAuthor(w, req)
		case http.MethodPut:
			implemetation.UpdateAuthor(w, req)
		default:
			fmt.Println("error")
		}
	})

	router.HandleFunc(fmt.Sprintf("/author/{%s:[0-9]*}", queryParamKey), func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			implemetation.GetAuthor(w, req)
		case http.MethodDelete:
			implemetation.DeleteAuthor(w, req)
		default:
			fmt.Println("error")
		}
	})
	return router
}

func (s *server) CreateAuthor(w http.ResponseWriter, req *http.Request) {
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
	_, err = s.authorRepo.Add(req.Context(), authorRepo)
	if err != nil {
		AnswerError(w, http.StatusInternalServerError)
		return
	}
}

func (s *server) GetAuthor(w http.ResponseWriter, req *http.Request) {
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
	author, err := s.authorRepo.GetByID(req.Context(), keyInt)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			AnswerError(w, http.StatusNotFound)
			return
		}
		AnswerError(w, http.StatusInternalServerError)
		return
	}
	authorJson, _ := json.Marshal(author)
	w.Write(authorJson)
}

func (s *server) UpdateAuthor(w http.ResponseWriter, req *http.Request) {
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
	err = s.authorRepo.Update(req.Context(), authorRepo)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			AnswerError(w, http.StatusNotFound)
			return
		}
		AnswerError(w, http.StatusInternalServerError)
		return
	}
}
func (s *server) DeleteAuthor(w http.ResponseWriter, req *http.Request) {
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
	err = s.authorRepo.DeleteById(req.Context(), keyInt)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			AnswerError(w, http.StatusNotFound)
			return
		}
		AnswerError(w, http.StatusInternalServerError)
		return
	}
}

func AnswerError(w http.ResponseWriter, statusCode int) {
	if statusCode == http.StatusNotFound {
		w.WriteHeader(http.StatusNotFound)
		body, _ := json.Marshal(map[string]string{"Error message": errorNotFound})
		w.Write([]byte(body))
	} else if statusCode == http.StatusInternalServerError {
		w.WriteHeader(http.StatusInternalServerError)
		body, _ := json.Marshal(map[string]string{"Error message": errorServerError})
		w.Write(body)
	} else if statusCode == http.StatusBadRequest {
		w.WriteHeader(http.StatusBadRequest)
		body, _ := json.Marshal(map[string]string{"Error message": errorBadRequest})
		w.Write(body)
	}
}
