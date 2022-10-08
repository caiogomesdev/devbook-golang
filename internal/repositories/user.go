package repositories

import (
	"fmt"

	"github.com/caiogomesdev/devbook-golang/internal/config"
	"github.com/caiogomesdev/devbook-golang/internal/models"
)

type userRepository struct {}

var User userRepository;

func (_ userRepository) Find(id uint64, user *models.User) error {
  db, err := config.ConfigureDb()
  if err != nil {
    return err
  }
  db.Model(models.User{}).Where("id = ?", id).Find(&user)
  return nil
}

func (_ userRepository) FindAll(users *[]models.User) error {
  db, err := config.ConfigureDb()
  if err != nil {
    return err
  }
  db1 := db.Find(&users)
  fmt.Println(db1)
  return nil
}



