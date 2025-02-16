package main

import (
	"log"
	"net/http"

	"fifth/calculators" 
)

func main() {
	// Маршрути для обробників
	http.HandleFunc("/calc1", calculators.Calc1)
	http.HandleFunc("/calc2", calculators.Calc2)

	// Обслуговування статичних файлів (CSS)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Сервер запущено на порту :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
