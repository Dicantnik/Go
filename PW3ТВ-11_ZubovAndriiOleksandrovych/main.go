package main

import (
	"log"
	"net/http"

	"third/calculators" // Замініть "your_project" на ім'я вашого модуля
)

func main() {
	// Обробник для головного маршруту
	http.HandleFunc("/", calculators.ProfitHandler)

	// Обслуговування статичних файлів (CSS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Сервер запущено на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
