package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConfigureDb() (*gorm.DB, error){
  host := os.Getenv("DB_HOST")
  user := os.Getenv("DB_USER")
  password := os.Getenv("DB_PASS")
  port := os.Getenv("DB_PORT")
  database := os.Getenv("DB_DATABASE")

  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, database, port)
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    return nil, err
  }
  return db, nil
}
