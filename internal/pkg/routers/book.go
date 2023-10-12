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

type addBookRequest struct {
	Name     string `json:"name"`
	AuthorId int64  `json:"author_id"`
}

type updateBookRequest struct {
	addBookRequest
	Id int64 `json:"id"`
}

func CreateBookSubRouter(router *mux.Router, s server.Server) *mux.Router {
	router.HandleFunc("/book", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			CreateBook(s, w, req)
		default:
			fmt.Println("error")
		}
	})

	router.HandleFunc(fmt.Sprintf("/book/{%s:[0-9]*}", queryParamKey), func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			GetBook(s, w, req)
		default:
			fmt.Println("error")
		}
	})
	return router
}

func CreateBook(s server.Server, w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	var unm addBookRequest
	if err = json.Unmarshal(body, &unm); err != nil {
		AnswerError(w, http.StatusInternalServerError)
		return
	}
	bookRepo := &repository.Book{
		Name:     unm.Name,
		AuthorId: unm.AuthorId,
	}
	_, err = s.BookRepo.Add(req.Context(), bookRepo)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			AnswerError(w, http.StatusNotFound)
			return
		}
		AnswerError(w, http.StatusInternalServerError)
		return
	}
}

func GetBook(s server.Server, w http.ResponseWriter, req *http.Request) {
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
	book, err := s.BookRepo.GetByID(req.Context(), keyInt)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			AnswerError(w, http.StatusNotFound)
			return
		}
		AnswerError(w, http.StatusInternalServerError)
		return
	}
	bookJson, err := json.Marshal(book)
	if err != nil {
		log.Fatal(err)
		AnswerError(w, http.StatusInternalServerError)
		return
	}
	w.Write(bookJson)
}
