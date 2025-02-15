package calculators

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

// calcNormalDistribution обчислює значення нормального розподілу для заданої точки.
func calcNormalDistribution(x, avg, sigma float64) float64 {
	return (1 / (sigma * math.Sqrt(2*math.Pi))) * math.Exp(-math.Pow(x-avg, 2)/(2*math.Pow(sigma, 2)))
}

// integrateTrapezoidal виконує інтеграцію функції нормального розподілу методом трапецій.
func integrateTrapezoidal(f func(float64, float64, float64) float64, avg, sigma, steps, deviation float64) float64 {
	lower := avg * (1 - deviation)
	upper := avg * (1 + deviation)
	h := (upper - lower) / steps
	var total float64

	for i := 0; i < int(steps); i++ {
		x1 := lower + float64(i)*h
		x2 := x1 + h
		total += 0.5 * (f(x1, avg, sigma) + f(x2, avg, sigma)) * h
	}

	return total
}

// calculateResults обчислює дохід, штраф та чистий прибуток.
func calculateResults(Pc, sigma, B float64) (revenue, fine, profit float64) {
	// Частка енергії, що генерується без небалансів (інтегруємо функцію нормального розподілу)
	balancedShare := integrateTrapezoidal(calcNormalDistribution, Pc, sigma, 10000, 0.05)

	revenue = Pc * 24 * balancedShare * B
	fine = Pc * 24 * (1 - balancedShare) * B
	profit = revenue - fine
	return
}

// ProfitHandler обробляє запити: показує форму та після POST‑запиту відображає результати.
func ProfitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Зчитування значень з форми
		Pc, err1 := strconv.ParseFloat(r.FormValue("Pc"), 64)
		sigma, err2 := strconv.ParseFloat(r.FormValue("Sigma"), 64)
		B, err3 := strconv.ParseFloat(r.FormValue("B"), 64)
		if err1 != nil || err2 != nil || err3 != nil {
			http.Error(w, "Невірні параметри", http.StatusBadRequest)
			return
		}

		revenue, fine, profit := calculateResults(Pc, sigma, B)

		// Завантаження шаблону
		tmpl, err := template.ParseFiles("templates/calc1.html")
		if err != nil {
			http.Error(w, "Не вдалося завантажити шаблон", http.StatusInternalServerError)
			return
		}

		data := map[string]string{
			"revenue": fmt.Sprintf("%.2f", revenue),
			"fine":    fmt.Sprintf("%.2f", fine),
			"profit":  fmt.Sprintf("%.2f", profit),
		}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "Помилка виконання шаблону", http.StatusInternalServerError)
		}
		return
	}

	// Для GET‑запиту просто показуємо форму
	tmpl, err := template.ParseFiles("templates/calc1.html")
	if err != nil {
		http.Error(w, "Не вдалося завантажити шаблон", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Помилка виконання шаблону", http.StatusInternalServerError)
	}
}
