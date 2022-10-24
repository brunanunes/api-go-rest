package main

import (
	"fmt"

	"github.com/brunanunes/go-rest-api/database"
	"github.com/brunanunes/go-rest-api/routes"
)

func main() {
	database.DatabaseConnect()
	fmt.Println("Iniciando o servidor Rest em Go")
	routes.HandleRequest()
}
