package main

import (
	"log"
	"net/http"

	"fourth/calculators"
)

func main() {
	// Регістрація обробників для кожного калькулятора
	http.HandleFunc("/calc1", calculators.Calc1Handler)
	http.HandleFunc("/calc2", calculators.Calc2Handler)
	http.HandleFunc("/calc3", calculators.Calc3Handler)
	// Обслуговування статичних файлів (CSS, зображення тощо)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Сервер запущено на порті :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
