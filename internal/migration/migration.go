package migration

import (
	"fmt"
	"log"

	"github.com/caiogomesdev/devbook-golang/internal/config"
	"github.com/caiogomesdev/devbook-golang/internal/entity"
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
