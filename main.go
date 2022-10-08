package main

import (
	"devbook/src/migration"
	"fmt"
	"os"
)

func main(){
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
