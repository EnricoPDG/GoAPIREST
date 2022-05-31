package main

import (
	"fmt"
	"github.com/EnricoPDG/GoAPIREST/models"
	"github.com/EnricoPDG/GoAPIREST/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1,Nome:"Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando servidor Rest Go")
	routes.HandleRequest()
}
