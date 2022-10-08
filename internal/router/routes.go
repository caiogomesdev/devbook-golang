package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
  URI string
  method string
  function func(w http.ResponseWriter, r *http.Request)
  authenticationRequired bool
}

func GetRoutes() *mux.Router{
  muxRoute := mux.NewRouter()

  var routesArray = [][]Route {
    UserRoutes,
  }

  for _, sliceRoute := range routesArray {
    executeSliceRoutes(sliceRoute, muxRoute)
  }

  return muxRoute;
}

func executeSliceRoutes(sliceRoute []Route, muxRoute *mux.Router){
  for _, route := range sliceRoute {
    muxRoute.HandleFunc(route.URI, route.function).Methods(route.method)
  }
}
