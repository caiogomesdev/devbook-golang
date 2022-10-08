package entity

import "gorm.io/gorm"

type User struct {
  gorm.Model
  Name string `gorm:"type:varchar(100);not null;default:null"`
  Email string `gorm:"type:varchar(150);not null;default:null;unique"`
  Password string `gorm:"type:varchar(150);not null;default:null"`
}
