package services

import (
	"github.com/caiogomesdev/devbook-golang/internal/controllers/dto"
	"github.com/caiogomesdev/devbook-golang/internal/models"
	"github.com/caiogomesdev/devbook-golang/internal/repositories"
	"github.com/caiogomesdev/devbook-golang/internal/utils"
)

type userService struct {}

var UserService userService

func (_ userService) GetById(id uint64, userDto *dto.UserDto) (error, bool) {
  var userModel models.User
  err := repositories.User.Find(id, &userModel)
  if err != nil {
    return err, false
  }
  if userModel.ID == 0 {
    return nil, false
  }
  *userDto = mapUserToDto(userModel)
  return nil, true
}

func (_ userService) GetAll(usersDto *[]dto.UserDto) error {
  var usersModel []models.User
  err := repositories.User.FindAll(&usersModel)
  if err != nil {
    return err
  }
  *usersDto = utils.SliceMap(usersModel, mapUserToDto)
  return nil
}

func (_ userService) Create(userDto *dto.UserDto) (error) {
  userModel := models.User{
    Name: userDto.Name,
    Nick: userDto.Nick,
    Email: userDto.Email,
    Password: userDto.Password,
  }
  err := repositories.User.Create(&userModel);
  if err != nil {
    return err
  }
  *userDto = mapUserToDto(userModel)
  return nil
}

func mapUserToDto(item models.User) dto.UserDto {
  return dto.UserDto{
    ID: item.ID,
    Name: item.Name,
    Nick: item.Nick,
    Email: item.Email,
    Password: "",
  }
}
