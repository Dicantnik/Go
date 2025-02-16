package calculators

import (
	"html/template"
	"math"
	"net/http"
	"strconv"
)

// Calc3Handler обробляє запити калькулятора 3
func Calc3Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Отримання значень з форми
		maxVoltageDrop, _ := strconv.ParseFloat(r.FormValue("maxVoltageDrop"), 64)
		transformerPower, _ := strconv.ParseFloat(r.FormValue("transformerPower"), 64)
		normalRes, _ := strconv.ParseFloat(r.FormValue("normalRes"), 64)
		normalReact, _ := strconv.ParseFloat(r.FormValue("normalReact"), 64)
		minimalRes, _ := strconv.ParseFloat(r.FormValue("minimalRes"), 64)
		minimalReact, _ := strconv.ParseFloat(r.FormValue("minimalReact"), 64)
		ratedSecondary, _ := strconv.ParseFloat(r.FormValue("ratedSecondary"), 64)
		primaryVoltage, _ := strconv.ParseFloat(r.FormValue("primaryVoltage"), 64)

		if maxVoltageDrop <= 0 || transformerPower <= 0 || normalRes <= 0 || normalReact <= 0 ||
			minimalRes <= 0 || minimalReact <= 0 || ratedSecondary <= 0 || primaryVoltage <= 0 {
			tmpl, _ := template.ParseFiles("templates/calc3.html")
			tmpl.Execute(w, map[string]interface{}{
				"error": "Будь ласка, введіть коректні значення для всіх полів.",
			})
			return
		}

		// Розрахунок імпедансу трансформатора
		xTransformer := (maxVoltageDrop * math.Pow(ratedSecondary, 2)) / (100 * transformerPower)

		// Розрахунки для нормального режиму
		xSheetNormal := normalReact + xTransformer
		impedanceNormal := math.Hypot(normalRes, xSheetNormal)
		threePhaseFaultNormal := (ratedSecondary * 1000) / (math.Sqrt(3) * impedanceNormal)
		twoPhaseFaultNormal := threePhaseFaultNormal * math.Sqrt(3) / 2

		// Розрахунки для мінімального режиму
		xSheetMinimal := minimalReact + xTransformer
		impedanceMinimal := math.Hypot(minimalRes, xSheetMinimal)
		threePhaseFaultMinimal := (ratedSecondary * 1000) / (math.Sqrt(3) * impedanceMinimal)
		twoPhaseFaultMinimal := threePhaseFaultMinimal * math.Sqrt(3) / 2

		// Коригування параметрів через коефіцієнт приведення
		transformCoeff := math.Pow(primaryVoltage, 2) / math.Pow(ratedSecondary, 2)
		adjustedNormalRes := normalRes * transformCoeff
		adjustedNormalReact := xSheetNormal * transformCoeff
		adjustedImpedanceNormal := math.Hypot(adjustedNormalRes, adjustedNormalReact)

		adjustedMinimalRes := minimalRes * transformCoeff
		adjustedMinimalReact := xSheetMinimal * transformCoeff
		adjustedImpedanceMinimal := math.Hypot(adjustedMinimalRes, adjustedMinimalReact)

		// Фактичні струми
		actualThreePhaseNormal := (primaryVoltage * 1000) / (math.Sqrt(3) * adjustedImpedanceNormal)
		actualTwoPhaseNormal := actualThreePhaseNormal * math.Sqrt(3) / 2
		actualThreePhaseMinimal := (primaryVoltage * 1000) / (math.Sqrt(3) * adjustedImpedanceMinimal)
		actualTwoPhaseMinimal := actualThreePhaseMinimal * math.Sqrt(3) / 2

		// Довжина лінії та параметри
		lineLength := 0.2 + 0.35 + 0.2 + 0.6 + 2 + 2.55 + 3.37 + 3.1
		baseRes := 0.64
		baseReact := 0.363

		lineRes := lineLength * baseRes
		lineReact := lineLength * baseReact

		// Розрахунки для точки 10 (нормальний режим)
		pointResNormal := lineRes + adjustedNormalRes
		pointReactNormal := lineReact + adjustedNormalReact
		pointImpedanceNormal := math.Hypot(pointResNormal, pointReactNormal)

		// Розрахунки для точки 10 (мінімальний режим)
		pointResMinimal := lineRes + adjustedMinimalRes
		pointReactMinimal := lineReact + adjustedMinimalReact
		pointImpedanceMinimal := math.Hypot(pointResMinimal, pointReactMinimal)

		// Струми короткого замикання в точці 10
		faultPointThreePhaseNormal := (primaryVoltage * 1000) / (math.Sqrt(3) * pointImpedanceNormal)
		faultPointTwoPhaseNormal := faultPointThreePhaseNormal * math.Sqrt(3) / 2
		faultPointThreePhaseMinimal := (primaryVoltage * 1000) / (math.Sqrt(3) * pointImpedanceMinimal)
		faultPointTwoPhaseMinimal := faultPointThreePhaseMinimal * math.Sqrt(3) / 2

		tmpl, _ := template.ParseFiles("templates/calc3.html")
		tmpl.Execute(w, map[string]interface{}{
			"xTransformer":               formatFloat(xTransformer, 2),
			"xSheetNormal":               formatFloat(xSheetNormal, 2),
			"impedanceNormal":            formatFloat(impedanceNormal, 2),
			"threePhaseFaultNormal":      formatFloat(threePhaseFaultNormal, 2),
			"twoPhaseFaultNormal":        formatFloat(twoPhaseFaultNormal, 2),
			"xSheetMinimal":              formatFloat(xSheetMinimal, 2),
			"impedanceMinimal":           formatFloat(impedanceMinimal, 2),
			"threePhaseFaultMinimal":     formatFloat(threePhaseFaultMinimal, 2),
			"twoPhaseFaultMinimal":       formatFloat(twoPhaseFaultMinimal, 2),
			"transformCoeff":             formatFloat(transformCoeff, 3),
			"adjustedNormalRes":          formatFloat(adjustedNormalRes, 2),
			"adjustedNormalReact":        formatFloat(adjustedNormalReact, 2),
			"adjustedImpedanceNormal":    formatFloat(adjustedImpedanceNormal, 2),
			"adjustedMinimalRes":         formatFloat(adjustedMinimalRes, 2),
			"adjustedMinimalReact":       formatFloat(adjustedMinimalReact, 2),
			"adjustedImpedanceMinimal":   formatFloat(adjustedImpedanceMinimal, 2),
			"actualThreePhaseNormal":     formatFloat(actualThreePhaseNormal, 2),
			"actualTwoPhaseNormal":       formatFloat(actualTwoPhaseNormal, 2),
			"actualThreePhaseMinimal":    formatFloat(actualThreePhaseMinimal, 2),
			"actualTwoPhaseMinimal":      formatFloat(actualTwoPhaseMinimal, 2),
			"lineLength":                 formatFloat(lineLength, 2),
			"lineRes":                    formatFloat(lineRes, 2),
			"lineReact":                  formatFloat(lineReact, 2),
			"pointResNormal":             formatFloat(pointResNormal, 2),
			"pointReactNormal":           formatFloat(pointReactNormal, 2),
			"pointImpedanceNormal":       formatFloat(pointImpedanceNormal, 2),
			"pointResMinimal":            formatFloat(pointResMinimal, 2),
			"pointReactMinimal":          formatFloat(pointReactMinimal, 2),
			"pointImpedanceMinimal":      formatFloat(pointImpedanceMinimal, 2),
			"faultPointThreePhaseNormal": formatFloat(faultPointThreePhaseNormal, 2),
			"faultPointTwoPhaseNormal":   formatFloat(faultPointTwoPhaseNormal, 2),
			"faultPointThreePhaseMinimal": formatFloat(faultPointThreePhaseMinimal, 2),
			"faultPointTwoPhaseMinimal":  formatFloat(faultPointTwoPhaseMinimal, 2),
		})
		return
	}
	tmpl, _ := template.ParseFiles("templates/calc3.html")
	tmpl.Execute(w, nil)
}
