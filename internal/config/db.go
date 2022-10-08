package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConfigureDb() (*gorm.DB, error){
  host := Env.Db.HOST
  user := Env.Db.USER
  password := Env.Db.PASSWORD
  port := Env.Db.PORT
  database := Env.Db.DATABASE

  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, database, port)
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    return nil, err
  }
  return db, nil
}
