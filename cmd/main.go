package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/caiogomesdev/devbook-golang/internal/config"
	"github.com/caiogomesdev/devbook-golang/internal/migration"
	"github.com/caiogomesdev/devbook-golang/internal/router"
)

func executeArg(args []string){
    if args[1] == "migrate"{
      migration.Migrate()
      return
    }
    fmt.Println("Comando não reconhecido")
    return
}

func main(){
  config.Init()
	args := os.Args
  if len(args) >= 2 {
    executeArg(args);
    return;
  }
  port := fmt.Sprintf(":%s", config.Env.Api.PORT);
  fmt.Println(fmt.Sprintf("Server is running on PORT%s", port))
  log.Fatal(http.ListenAndServe(port, router.GetRoutes()))
}
