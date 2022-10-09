package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var instanceDb *gorm.DB

func configureDb() (*gorm.DB, error){
  host := Env.Db.HOST
  user := Env.Db.USER
  password := Env.Db.PASSWORD
  port := Env.Db.PORT
  database := Env.Db.DATABASE

  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, database, port)
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic(err)
  }
  instanceDb = db
  return db, nil
}

func Db() *gorm.DB{
  return instanceDb
}
