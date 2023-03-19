package main

import (
	"application/database"
	"application/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
