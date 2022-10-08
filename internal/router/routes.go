package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router{
  route := mux.NewRouter()

  route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{ "message": "hello world"})
  }).Methods(http.MethodGet)

  return route;
}
