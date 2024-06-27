package main

import (
	"ApiGo/database"
	"ApiGo/routes"
	"fmt"
)

func main() {
	database.ConectToDatabase()
	fmt.Println("Iniciando o servidor...")
	routes.HandleRequest()
}
