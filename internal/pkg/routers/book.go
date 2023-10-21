package routers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type addBookRequest struct {
	Name     string `json:"name"`
	AuthorId int64  `json:"author_id"`
}

type UpdateBookRequest struct {
	addBookRequest
	Id int64 `json:"id"`
}

func CreateBookSubRouter(router *mux.Router, libraryApp LibraryApp) *mux.Router {
	router.HandleFunc("/book", func(w http.ResponseWriter, req *http.Request) {
		updateBookData, status := parseUpdateBookRequest(req)
		if status != http.StatusOK {
			AnswerError(w, status)
			return
		}
		switch req.Method {
		case http.MethodPost:
			if status = libraryApp.CreateBook(req.Context(), updateBookData); status != http.StatusOK {
				AnswerError(w, status)
			} else {
				w.WriteHeader(int(status))
			}
		default:
			fmt.Println("error")
		}
	})

	router.HandleFunc(fmt.Sprintf("/book/{%s:[0-9]*}", queryParamKey), func(w http.ResponseWriter, req *http.Request) {
		id, status := ParseID(req)
		if status != http.StatusOK {
			AnswerError(w, status)
			return
		}
		switch req.Method {
		case http.MethodGet:
			bookJson, status := libraryApp.GetBook(req.Context(), id)
			if status != http.StatusOK {
				AnswerError(w, status)
			} else {
				w.WriteHeader(int(status))
				w.Write(bookJson)
			}
		default:
			fmt.Println("error")
		}
	})
	return router
}

func parseUpdateBookRequest(req *http.Request) (*UpdateBookRequest, StatusInt) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, http.StatusBadRequest
	}
	var unm UpdateBookRequest
	if err = json.Unmarshal(body, &unm); err != nil {
		return nil, http.StatusBadRequest
	}
	return &unm, http.StatusOK
}
