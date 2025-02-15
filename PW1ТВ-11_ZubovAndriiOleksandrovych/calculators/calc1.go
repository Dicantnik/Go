package calculators

import (
	"html/template"
	"net/http"
	"strconv"
)

// CalculatorOneHandler обробляє запити для першого калькулятора.
func CalculatorOneHandler(resp http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		// Зчитування значень форми
		hydrogen, _ := strconv.ParseFloat(req.FormValue("H"), 64)
		carbon, _ := strconv.ParseFloat(req.FormValue("C"), 64)
		sulfur, _ := strconv.ParseFloat(req.FormValue("S"), 64)
		nitrogen, _ := strconv.ParseFloat(req.FormValue("N"), 64)
		oxygen, _ := strconv.ParseFloat(req.FormValue("O"), 64)
		water, _ := strconv.ParseFloat(req.FormValue("W"), 64)
		ash, _ := strconv.ParseFloat(req.FormValue("A"), 64)

		// Обчислення коефіцієнтів перерахунку
		dryCoef := 100 / (100 - water)
		combCoef := 100 / (100 - water - ash)

		// Розрахунок нижчої теплоти згоряння для робочої маси
		lowerHeat := (339*carbon + 1030*hydrogen - 108.8*(oxygen-sulfur) - 25*water) / 1000

		// Розрахунок теплоти для сухої та горючої мас
		lhvDry := (lowerHeat + 0.025*water) * dryCoef
		lhvComb := (lowerHeat + 0.025*water) * combCoef

		// Перерахунок складу сухої маси
		hDry := hydrogen * dryCoef
		cDry := carbon * dryCoef
		sDry := sulfur * dryCoef
		nDry := nitrogen * dryCoef
		oDry := oxygen * dryCoef
		aDry := ash * dryCoef

		// Перерахунок складу горючої маси
		hComb := hydrogen * combCoef
		cComb := carbon * combCoef
		sComb := sulfur * combCoef
		nComb := nitrogen * combCoef
		oComb := oxygen * combCoef

		// Форматування чисел
		precShort := 2
		precLong := 4
		tmpl, _ := template.ParseFiles("templates/task1.html")
		tmpl.Execute(resp, map[string]string{
			"DryCoef":    roundValue(dryCoef, precShort),
			"CombCoef":   roundValue(combCoef, precShort),
			"LowerHeat":  roundValue(lowerHeat, precLong),
			"LhvDry":     roundValue(lhvDry, precLong),
			"LhvComb":    roundValue(lhvComb, precLong),
			"HDry":       roundValue(hDry, precShort),
			"CDry":       roundValue(cDry, precShort),
			"SDry":       roundValue(sDry, precShort),
			"NDry":       roundValue(nDry, precShort),
			"ODry":       roundValue(oDry, precShort),
			"ADry":       roundValue(aDry, precShort),
			"HComb":      roundValue(hComb, precShort),
			"CComb":      roundValue(cComb, precShort),
			"SComb":      roundValue(sComb, precShort),
			"NComb":      roundValue(nComb, precShort),
			"OComb":      roundValue(oComb, precShort),
		})
		return
	}

	// Відображення сторінки при GET-запиті
	tmpl, _ := template.ParseFiles("templates/task1.html")
	tmpl.Execute(resp, nil)
}
