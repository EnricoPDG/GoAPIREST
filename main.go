package main

import (
	"fmt"
	"github.com/EnricoPDG/GoAPIREST/models"
	"github.com/EnricoPDG/GoAPIREST/routes"
)

func main() {
	models.Personalidades = []models.Personalidade {
		{Nome: "Nome 1"},
		{Historia: "Historia 1"},
		{Nome: "Nome 2"},
		{Historia: "Historia 2"},
	}


	fmt.Println("Iniciando servidor Rest Go")
	routes.HandleRequest()
}
