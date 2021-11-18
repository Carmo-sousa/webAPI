package main

import (
	"github.com/Carmo-sousa/webAPI/database"
	"github.com/Carmo-sousa/webAPI/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
