package calculators

import (
	"html/template"
	"net/http"
	"strconv"
)

// CalculatorTwoHandler обробляє запити для другого калькулятора.
func CalculatorTwoHandler(resp http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		// Зчитування значень форми
		carbonVal, _ := strconv.ParseFloat(req.FormValue("C"), 64)
		hydrogenVal, _ := strconv.ParseFloat(req.FormValue("H"), 64)
		sulfurVal, _ := strconv.ParseFloat(req.FormValue("S"), 64)
		oxygenVal, _ := strconv.ParseFloat(req.FormValue("O"), 64)
		waterVal, _ := strconv.ParseFloat(req.FormValue("W"), 64)
		ashVal, _ := strconv.ParseFloat(req.FormValue("A"), 64)
		volatileVal, _ := strconv.ParseFloat(req.FormValue("V"), 64)
		qCombInput, _ := strconv.ParseFloat(req.FormValue("Q_comb"), 64)

		// Обчислення коефіцієнта переходу до робочої маси
		workCoef := (100 - waterVal - ashVal) / 100
		carbonWork := carbonVal * workCoef
		hydrogenWork := hydrogenVal * workCoef
		sulfurWork := sulfurVal * workCoef
		oxygenWork := oxygenVal * workCoef
		volatileWork := volatileVal * (100 - waterVal) / 100
		qResidual := qCombInput*(100-waterVal-ashVal)/100 - 0.025*waterVal

		// Форматування чисел
		prec := 2
		tmpl, _ := template.ParseFiles("templates/task2.html")
		tmpl.Execute(resp, map[string]string{
			"WorkCoef":     roundValue(workCoef, prec),
			"CarbonWork":   roundValue(carbonWork, prec),
			"HydrogenWork": roundValue(hydrogenWork, prec),
			"SulfurWork":   roundValue(sulfurWork, prec),
			"OxygenWork":   roundValue(oxygenWork, prec),
			"VolatileWork": roundValue(volatileWork, 1),
			"Ash":          roundValue(ashVal, prec),
			"QResidual":    roundValue(qResidual, prec),
		})
		return
	}

	// Відображення сторінки при GET-запиті
	tmpl, _ := template.ParseFiles("templates/task2.html")
	tmpl.Execute(resp, nil)
}
