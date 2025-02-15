package calculators

import (
	"html/template"
	"net/http"
	"strconv"
	"fmt"
)

func roundValue(num float64, decimals int) string {
	format := fmt.Sprintf("%%.%df", decimals)
	return fmt.Sprintf(format, num)
}

func Secondcalc(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		energyInput, _ := strconv.ParseFloat(r.FormValue("energyInput"), 64)
		areaRatio, _ := strconv.ParseFloat(r.FormValue("areaRatio"), 64)
		gasOutput, _ := strconv.ParseFloat(r.FormValue("gasOutput"), 64)
		efficiencyFactor, _ := strconv.ParseFloat(r.FormValue("efficiencyFactor"), 64)
		heatOutput, _ := strconv.ParseFloat(r.FormValue("heatOutput"), 64)
		coefficient, _ := strconv.ParseFloat(r.FormValue("coefficient"), 64)
		tempFactor, _ := strconv.ParseFloat(r.FormValue("tempFactor"), 64)

		heatCoeff := (1e6/energyInput)*heatOutput*(areaRatio/(100-gasOutput))*(1-efficiencyFactor) + tempFactor
		energyResult := 1e-6 * heatCoeff * energyInput * coefficient

		tmpl, _ := template.ParseFiles("templates/calc1.html")
		tmpl.Execute(w, map[string]string{
			"heatCoeff":    roundValue(heatCoeff, 2),
			"energyResult": roundValue(energyResult, 2),
		})
		return
	}

	tmpl, _ := template.ParseFiles("templates/calc1.html")
	tmpl.Execute(w, nil)
}
