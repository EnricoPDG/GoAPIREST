package main

import (
	"fmt"
	"github.com/EnricoPDG/GoAPIREST/database"
	"github.com/EnricoPDG/GoAPIREST/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando servidor Rest Go")
	routes.HandleRequest()
}
