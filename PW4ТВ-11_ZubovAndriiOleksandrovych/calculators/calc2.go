package calculators

import (
	"html/template"
	"math"
	"net/http"
	"strconv"
)

// Calc2Handler обробляє запити калькулятора 2
func Calc2Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		voltageInput, _ := strconv.ParseFloat(r.FormValue("voltage"), 64)
		impedanceInput, _ := strconv.ParseFloat(r.FormValue("impedance"), 64)

		voltageDropPercent := 10.5
		nominalTransformerPower := 6.3

		if voltageInput > 0 && impedanceInput > 0 {
			reactanceCircuit := (voltageInput * voltageInput) / impedanceInput
			transformerReactance := (voltageDropPercent / 100) * ((voltageInput * voltageInput) / nominalTransformerPower)
			totalReactance := reactanceCircuit + transformerReactance
			faultCurrentCalculated := (voltageInput * 1000) / (math.Sqrt(3) * totalReactance)

			tmpl, _ := template.ParseFiles("templates/calc2.html")
			tmpl.Execute(w, map[string]interface{}{
				"result":               true,
				"reactanceCircuit":     formatFloat(reactanceCircuit, 2),
				"transformerReactance": formatFloat(transformerReactance, 2),
				"totalReactance":       formatFloat(totalReactance, 2),
				"faultCurrent":         formatFloat(faultCurrentCalculated, 2),
			})
			return
		}
	}
	tmpl, _ := template.ParseFiles("templates/calc2.html")
	tmpl.Execute(w, nil)
}
