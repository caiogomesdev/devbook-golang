package controller

import "net/http"


type userController struct {}

var User userController;

func (user userController) GetAll(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("TODO"))
}

func (user userController) Find(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("TODO"))
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
