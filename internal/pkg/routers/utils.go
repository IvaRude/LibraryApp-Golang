package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"homework-3/internal/pkg/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type StatusInt int

type LibraryApp interface {
	CreateAuthor(ctx context.Context, updateAuthorData *UpdateAuthorRequest) StatusInt
	GetAuthor(ctx context.Context, id int64) ([]byte, StatusInt)
	UpdateAuthor(ctx context.Context, updateAuthorData *UpdateAuthorRequest) StatusInt
	DeleteAuthor(ctx context.Context, id int64) StatusInt

	CreateBook(ctx context.Context, updateBookData *UpdateBookRequest) StatusInt
	GetBook(ctx context.Context, id int64) ([]byte, StatusInt)
}

func ParseID(req *http.Request) (int64, StatusInt) {
	id, ok := mux.Vars(req)[queryParamKey]
	if !ok {
		return 0, http.StatusBadRequest
	}
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, http.StatusBadRequest
	}
	return idInt, http.StatusOK
}

func BuildHandlerMessage(body []byte, eventType string, method string) (*models.HandlerMessage, error) {
	var request models.Request
	if len(body) > 0 {
		var item models.Item
		err := json.Unmarshal(body, &item)
		if err != nil {
			return nil, err
		}
		request = models.Request{Method: method, Body: fmt.Sprint(item)}
	} else {
		request = models.Request{Method: method, Body: ""}
	}
	return &models.HandlerMessage{Timestamp: time.Now(), Req: request, EventType: eventType}, nil
}
