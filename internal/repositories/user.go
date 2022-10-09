package repositories

import (
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
  db.Order("id desc").Omit("password").Find(&users)
  return nil
}

func (_ userRepository) Create(userModel *models.User) error {
  db, err := config.ConfigureDb()
  if err != nil {
    return err
  }
  gormDb := db.Create(&userModel)
  if gormDb.Error != nil {
    return gormDb.Error
  }
  return nil
}
