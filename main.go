package main

import (
	"devbook/src/migration"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type User struct {
  Name string
}

func main(){
  err := godotenv.Load();
  if err != nil {
    message := fmt.Sprintf("Error to Load .env file")
    log.Fatal(message)
  }
	args := os.Args
  if len(args) >= 2 {
    if args[1] == "migrate"{
      migration.Migrate()
      return
    }
    fmt.Println("Comando não reconhecido")
    return
  }
}
