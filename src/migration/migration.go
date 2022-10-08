package migration

import (
	"devbook/src/config"
	"devbook/src/entity"
	"fmt"
	"log"
)

func Migrate() {
  db, err := config.ConfigureDb()
  if err != nil {
    log.Fatal("Error configuring database")
  }
  err = db.Migrator().AutoMigrate(entity.User{})
  if err != nil {
    log.Fatal(fmt.Sprintf("Error to migrator: %v", err));
  }
	fmt.Println("Migration done");
}
