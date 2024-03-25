package main

import (
	"github.com/geblauth/GoRestApi/models"
	"github.com/geblauth/GoRestApi/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{

		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Hisotir 2"},
	}

	routes.HandleRequest()

}
