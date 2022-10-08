package services

import (
	"github.com/caiogomesdev/devbook-golang/internal/models"
	"github.com/caiogomesdev/devbook-golang/internal/repositories"
)

type userService struct {}

var UserService userService

func (_ userService) GetById(id uint64, user *models.User) (error, bool) {
  err := repositories.User.Find(id, user)

  if err != nil {
    return err, false
  }
  if user.ID == 0 {
    return nil, false
  }
  return nil, true
}
