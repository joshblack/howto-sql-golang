package api

import (
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()
	return r
}
