package main

import (
	"fmt"
	"log"
	"os"

	"github.com/caiogomesdev/devbook-golang/internal/migration"

	"github.com/joho/godotenv"
)

type User struct {
  Name string
}

func executeArg(args []string){
    if args[1] == "migrate"{
      migration.Migrate()
      return
    }
    fmt.Println("Comando não reconhecido")
    return
}

func main(){
  err := godotenv.Load();
  if err != nil {
    message := fmt.Sprintf("Error to Load .env file")
    log.Fatal(message)
  }
	args := os.Args
  if len(args) >= 2 {
    executeArg(args);
  }
}
