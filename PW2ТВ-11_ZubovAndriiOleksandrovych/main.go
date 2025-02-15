package main

import (
	"log"
	"net/http"
	"second/calculators"
)

func main() {
	http.HandleFunc("/calc2", calculators.Secondcalc)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
