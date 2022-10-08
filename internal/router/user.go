package router

import (
	"net/http"

	"github.com/caiogomesdev/devbook-golang/internal/controllers"
)

var UserRoutes = []Route{
  {
    URI: "/users",
    method: http.MethodGet,
    function: controllers.User.GetAll,
    authenticationRequired: false,
  },
  {
    URI: "/user/{id}",
    method: http.MethodGet,
    function: controllers.User.Find,
    authenticationRequired: false,
  },
  {
    URI: "/users",
    method: http.MethodPost,
    function: controllers.User.Create,
    authenticationRequired: false,
  },
  {
    URI: "/user/{id}",
    method: http.MethodPut,
    function: controllers.User.Update,
    authenticationRequired: false,
  },
  {
    URI: "/user/{id}",
    method: http.MethodDelete,
    function: controllers.User.Delete,
    authenticationRequired: false,
  },
}
