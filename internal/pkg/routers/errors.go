package routers

import (
	"encoding/json"
	"net/http"
)

const (
	errorNotFound    = "Not Found"
	errorBadRequest  = "Bad Request"
	errorServerError = "Server Error"
)

func AnswerError(w http.ResponseWriter, statusCode statusInt) {
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
