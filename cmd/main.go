package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"homework-3/internal/pkg/db"
	"homework-3/internal/pkg/repository"
	"homework-3/internal/pkg/repository/postgresql"

	"github.com/gorilla/mux"
)

const port = ":9000"
const queryParamKey = "key"
const errorNotFound = "Not Found"
const errorBadRequest = "Bad Request"
const errorServerError = "Server Error"

type server struct {
	authorRepo *postgresql.AuthorRepo
}

type addAuthorRequest struct {
	Name string `json:"name"`
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

	router.HandleFunc(fmt.Sprintf("/author/{%s:[0-9]+}", queryParamKey), func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			implemetation.Get(w, req)
		case http.MethodDelete:
			implemetation.Delete(w, req)
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
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	authorRepo := &repository.Author{
		Name: unm.Name,
	}
	id, err := s.authorRepo.Add(req.Context(), authorRepo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	authorRepo.Id = id
	authorJson, _ := json.Marshal(authorRepo)
	w.Write(authorJson)
}

func (s *server) Get(w http.ResponseWriter, req *http.Request) {
	key, ok := mux.Vars(req)[queryParamKey]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		body, _ := json.Marshal(map[string]string{"Error message": errorBadRequest})
		w.Write(body)
		return
	}
	keyInt, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		body, _ := json.Marshal(map[string]string{"Error message": errorBadRequest})
		w.Write(body)
		return
	}
	author, err := s.authorRepo.GetByID(req.Context(), keyInt)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			// w.WriteHeader(http.StatusNotFound)
			// body, _ := json.Marshal(map[string]string{"Error message": errorNotFound})
			w.Write([]byte(errorNotFound))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		body, _ := json.Marshal(map[string]string{"Error message": errorServerError})
		w.Write(body)
		return
	}
	authorJson, _ := json.Marshal(author)
	w.Write(authorJson)
}

func (s *server) UpdateAuthor(_ http.ResponseWriter, req *http.Request) {
	fmt.Println("update")
}
func (s *server) Delete(w http.ResponseWriter, req *http.Request) {
	fmt.Println("delete")

	//key, ok := mux.Vars(req)[queryParamKey]
	//if !ok {
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//_, ok = s.data[key]
	//if !ok {
	//	w.WriteHeader(http.StatusNotFound)
	//	return
	//}
	//
	//delete(s.data, key)

}
