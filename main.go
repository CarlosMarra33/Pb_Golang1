package main

import (
	"application/database"
	"application/server"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	// http.ListenAndServe(":5000", nil)

	database.StartDB()
	s := server.NewServer()

	s.Run()
}
