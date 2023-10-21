package routers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

const queryParamKey = "key"

type AddAuthorRequest struct {
	Name string `json:"name"`
}

type UpdateAuthorRequest struct {
	AddAuthorRequest
	Id int64 `json:"id"`
}

func CreateAuthorRouter(router *mux.Router, libraryApp LibraryApp) *mux.Router {
	router.HandleFunc("/author", func(w http.ResponseWriter, req *http.Request) {
		updateAuthorData, status := parseUpdateAuthorRequest(req)
		if status != http.StatusOK {
			AnswerError(w, status)
			return
		}
		switch req.Method {
		case http.MethodPost:
			if status = libraryApp.CreateAuthor(req.Context(), updateAuthorData); status != http.StatusOK {
				AnswerError(w, status)
			} else {
				w.WriteHeader(int(status))
			}
		case http.MethodPut:
			if status = libraryApp.UpdateAuthor(req.Context(), updateAuthorData); status != http.StatusOK {
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
			authorJson, status := libraryApp.GetAuthor(req.Context(), id)
			if status != http.StatusOK {
				AnswerError(w, status)
			} else {
				w.WriteHeader(int(status))
				w.Write(authorJson)
			}
		case http.MethodDelete:
			if status = libraryApp.DeleteAuthor(req.Context(), id); status != http.StatusOK {
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

func parseUpdateAuthorRequest(req *http.Request) (*UpdateAuthorRequest, StatusInt) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, http.StatusBadRequest
	}
	var unm UpdateAuthorRequest
	if err = json.Unmarshal(body, &unm); err != nil {
		return nil, http.StatusBadRequest
	}
	return &unm, http.StatusOK
}
