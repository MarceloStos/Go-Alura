package main

import (
	"github.com/MarceloStos/Go-Alura/database"
	"github.com/MarceloStos/Go-Alura/routes"
)

func main() {
	database.ConectacomBancoDeDados()

	routes.HandlerRequests()
}
