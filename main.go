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

type Shape interface {
	Area() float64
 }
 
 type Rectangle struct {
	Length float64
	Width  float64
 }
 
 func (r Rectangle) Area() float64 {
	return r.Length * r.Width
 }
 
 type Square struct {
	Side float64
 }
 
 func (s Square) Area() float64 {
	return s.Side * s.Side
 }
 
 func GetTotalArea(shapes []Shape) float64 {
	var totalArea float64
	for _, shape := range shapes {
	   totalArea += shape.Area()
	}
	return totalArea
 }