package router

import (
	"net/http"

	"github.com/caiogomesdev/devbook-golang/internal/controller"
)

var UserRoutes = []Route{
  {
    URI: "/users",
    method: http.MethodGet,
    function: controller.User.GetAll,
    authenticationRequired: false,
  },
  {
    URI: "/user/{id}",
    method: http.MethodGet,
    function: controller.User.Find,
    authenticationRequired: false,
  },
  {
    URI: "/users",
    method: http.MethodPost,
    function: controller.User.Create,
    authenticationRequired: false,
  },
  {
    URI: "/user/{id}",
    method: http.MethodPut,
    function: controller.User.Update,
    authenticationRequired: false,
  },
  {
    URI: "/user/{id}",
    method: http.MethodDelete,
    function: controller.User.Delete,
    authenticationRequired: false,
  },
}
