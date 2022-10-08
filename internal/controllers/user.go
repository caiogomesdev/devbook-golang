package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/caiogomesdev/devbook-golang/internal/models"
	"github.com/caiogomesdev/devbook-golang/internal/services"
	"github.com/gorilla/mux"
)

type userController struct {}

var User userController;

func (user userController) GetAll(w http.ResponseWriter, r *http.Request) {
  var users []models.User

  err := services.UserService.GetAll(&users)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  json.NewEncoder(w).Encode(users)
}

func (_ userController) Find(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  id, _ := strconv.ParseUint(params["id"], 10, 8)

  var user models.User;
  err, result := services.UserService.GetById(id, &user)

  if err != nil || !result {
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{ "message":  "Não foi possível encontrar usuário com esse id" })
    return
  }
  json.NewEncoder(w).Encode(user)
}

func (user userController) Create(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("TODO"))
}

func (user userController) Update(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("TODO"))
}

func (user userController) Delete(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("TODO"))
}
