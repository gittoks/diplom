package main

import (
	"github.com/gittoks/diplom/server/database"
	"github.com/gittoks/diplom/server/routes"
)

func main() {
	database.Start()
	routes.Start()
}
