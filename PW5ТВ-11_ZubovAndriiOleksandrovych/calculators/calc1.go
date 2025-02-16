package calculators

import (
	"html/template"
	"math"
	"net/http"
	"strconv"
)

// LossHandler обчислює втрати за введеними параметрами
func Calc1(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Зчитування параметрів з форми
		angFrequency, _ := strconv.ParseFloat(r.FormValue("omega"), 64)
		timeInterval, _ := strconv.ParseFloat(r.FormValue("tb"), 64)
		nominalPower, _ := strconv.ParseFloat(r.FormValue("pNom"), 64)
		timeMultiplier, _ := strconv.ParseFloat(r.FormValue("tm"), 64)
		powerCoefficient, _ := strconv.ParseFloat(r.FormValue("kp"), 64)
		baseLoss1, _ := strconv.ParseFloat(r.FormValue("zPer0"), 64)
		baseLoss2, _ := strconv.ParseFloat(r.FormValue("zPlan"), 64)

		// Обчислення втрат
		avarLoss := angFrequency * nominalPower * timeInterval * timeMultiplier
		planLoss := powerCoefficient * nominalPower * timeMultiplier
		totalLoss := baseLoss1 + (avarLoss * baseLoss1) + (planLoss * baseLoss2)

		// Відображення результатів у шаблоні calc1.html
		tmpl, _ := template.ParseFiles("templates/calc1.html")
		tmpl.Execute(w, map[string]string{
			"LossAvar":  roundTo(avarLoss, 4),
			"LossPlan":  roundTo(planLoss, 4),
			"TotalLoss": roundTo(totalLoss, 4),
		})
		return
	}

	// Відображення форми, якщо метод не POST
	tmpl, _ := template.ParseFiles("templates/calc1.html")
	tmpl.Execute(w, nil)
}

// Структура параметрів для розрахунку надійності
type ReliabilityMetric struct {
	FailureProbability float64
	RepairDuration     int
	OccurrenceRate     float64
	RecoveryTime       int
}

// Дані для розрахунку надійності
var reliabilityMetrics = map[string]ReliabilityMetric{
	"T-110 kV":                     {FailureProbability: 0.015, RepairDuration: 100, OccurrenceRate: 1.0, RecoveryTime: 43},
	"T-35 kV":                      {FailureProbability: 0.02, RepairDuration: 80, OccurrenceRate: 1.0, RecoveryTime: 28},
	"T-10 kV (Cable Network)":      {FailureProbability: 0.005, RepairDuration: 60, OccurrenceRate: 0.5, RecoveryTime: 10},
	"T-10 kV (Overhead Network)":   {FailureProbability: 0.05, RepairDuration: 60, OccurrenceRate: 0.5, RecoveryTime: 10},
	"B-110 kV (Gas-Insulated)":     {FailureProbability: 0.01, RepairDuration: 30, OccurrenceRate: 0.1, RecoveryTime: 30},
	"B-10 kV (Oil)":                {FailureProbability: 0.02, RepairDuration: 15, OccurrenceRate: 0.33, RecoveryTime: 15},
	"B-10 kV (Vacuum)":             {FailureProbability: 0.05, RepairDuration: 15, OccurrenceRate: 0.33, RecoveryTime: 15},
	"Busbars 10 kV per Connection": {FailureProbability: 0.03, RepairDuration: 2, OccurrenceRate: 0.33, RecoveryTime: 15},
	"AV-0.38 kV":                   {FailureProbability: 0.05, RepairDuration: 20, OccurrenceRate: 1.0, RecoveryTime: 15},
	"ED 6,10 kV":                   {FailureProbability: 0.1, RepairDuration: 50, OccurrenceRate: 0.5, RecoveryTime: 0},
	"ED 0.38 kV":                   {FailureProbability: 0.1, RepairDuration: 50, OccurrenceRate: 0.5, RecoveryTime: 0},
	"PL-110 kV":                    {FailureProbability: 0.007, RepairDuration: 10, OccurrenceRate: 0.167, RecoveryTime: 35},
	"PL-35 kV":                     {FailureProbability: 0.02, RepairDuration: 8, OccurrenceRate: 0.167, RecoveryTime: 35},
	"PL-10 kV":                     {FailureProbability: 0.02, RepairDuration: 10, OccurrenceRate: 0.167, RecoveryTime: 35},
	"CL-10 kV (Trench)":            {FailureProbability: 0.03, RepairDuration: 44, OccurrenceRate: 1.0, RecoveryTime: 9},
	"CL-10 kV (Cable Channel)":     {FailureProbability: 0.005, RepairDuration: 18, OccurrenceRate: 1.0, RecoveryTime: 9},
}

// ReliabilityHandler виконує розрахунок надійності
func Calc2(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var totalFailProb float64
		var weightedRepair float64

		// Обхід усіх параметрів надійності
		for key, metric := range reliabilityMetrics {
			count, _ := strconv.Atoi(r.FormValue(key))
			if count > 0 {
				totalFailProb += float64(count) * metric.FailureProbability
				weightedRepair += float64(count) * metric.FailureProbability * float64(metric.RepairDuration)
			}
		}

		meanRepair := weightedRepair / totalFailProb
		unplannedDowntime := meanRepair * totalFailProb / 8760
		scheduledDowntime := 1.2 * 43 / 8760
		combinedFailProb := 2 * totalFailProb * (unplannedDowntime + scheduledDowntime)
		finalFailProb := combinedFailProb + 0.02

		// Відображення результатів у шаблоні calc2.html
		tmpl, _ := template.ParseFiles("templates/calc2.html")
		tmpl.Execute(w, map[string]string{
			"TotalFailure": roundTo(totalFailProb, 4),
			"MeanRepair":   roundTo(meanRepair, 4),
			"Unplanned":    roundTo(unplannedDowntime, 4),
			"Scheduled":    roundTo(scheduledDowntime, 4),
			"Combined":     roundTo(combinedFailProb, 4),
			"Final":        roundTo(finalFailProb, 4),
		})
		return
	}

	// Відображення форми, якщо метод не POST
	tmpl, _ := template.ParseFiles("templates/calc2.html")
	tmpl.Execute(w, nil)
}

// roundTo округлює число до заданої кількості знаків після коми і повертає рядок
func roundTo(val float64, precision int) string {
	factor := math.Pow10(precision)
	rounded := math.Round(val*factor) / factor
	return strconv.FormatFloat(rounded, 'f', precision, 64)
}
