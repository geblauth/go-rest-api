package main

import (
	"fmt"

	"github.com/geblauth/GoRestApi/database"
	"github.com/geblauth/GoRestApi/routes"
)

func main() {

	fmt.Println("Pera")
	database.ConectaBandoDeDados()
	fmt.Println("Sucesso")
	routes.HandleRequest()

}
