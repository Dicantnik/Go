package calculators

import (
	"html/template"
	"math"
	"net/http"
	"strconv"
)

// Calc1Handler обробляє запити калькулятора 1
func Calc1Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Отримуємо дані форми
		areaValue, _ := strconv.ParseFloat(r.FormValue("area"), 64)
		faultCurrentInput, _ := strconv.ParseFloat(r.FormValue("faultCurrent"), 64)
		timeFactor, _ := strconv.ParseFloat(r.FormValue("timeFactor"), 64)
		ratedVolt, _ := strconv.ParseFloat(r.FormValue("ratedVolt"), 64)
		extraField, _ := strconv.ParseFloat(r.FormValue("extraField"), 64) // додаткове поле

		// Константи
		currentDensity := 1.4
		thermalConst := 92.0

		// Перевірка коректності введених даних
		if areaValue > 0 && faultCurrentInput > 0 && timeFactor > 0 && ratedVolt > 0 && extraField > 0 {
			currentCalc := areaValue / (2 * math.Sqrt(3) * ratedVolt)
			doubledCalc := 2 * currentCalc
			sectionCalc := currentCalc / currentDensity
			minimalSection := (faultCurrentInput * math.Sqrt(timeFactor)) / thermalConst

			tmpl, _ := template.ParseFiles("templates/calc1.html")
			tmpl.Execute(w, map[string]interface{}{
				"result":         true,
				"currentCalc":    formatFloat(currentCalc, 2),
				"doubledCalc":    formatFloat(doubledCalc, 2),
				"sectionCalc":    formatFloat(sectionCalc, 2),
				"minimalSection": formatFloat(minimalSection, 2),
			})
			return
		}
	}
	// Якщо GET-запит або дані не введено
	tmpl, _ := template.ParseFiles("templates/calc1.html")
	tmpl.Execute(w, nil)
}
