package models

import "gorm.io/gorm"

type User struct {
  gorm.Model
  Name string `gorm:"type:varchar(50);not null;default:null"`
  Nick string `gorm:"type:varchar(50);not null;default:null;unique"`
  Email string `gorm:"type:varchar(100);not null;default:null;unique"`
  Password string `gorm:"type:varchar(100);not null;default:null"`
}
