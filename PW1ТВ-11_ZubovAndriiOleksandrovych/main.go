package main

import (
	"log"
	"net/http"

	"myapp/calculators" 
)

func main() {
	// Прив'язка обробників до маршрутів
	http.HandleFunc("/task1", calculators.CalculatorOneHandler)
	http.HandleFunc("/task2", calculators.CalculatorTwoHandler)

	// Налаштування обслуговування статичних файлів (CSS, зображення тощо)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Сервер запущено на порті :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
