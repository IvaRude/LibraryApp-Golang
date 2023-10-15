package routers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type statusInt int

func parseID(req *http.Request) (int64, statusInt) {
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
