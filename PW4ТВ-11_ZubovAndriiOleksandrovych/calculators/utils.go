package calculators

import "fmt"

// formatFloat округлює число до заданої кількості знаків
func formatFloat(val float64, precision int) string {
	return fmt.Sprintf("%.*f", precision, val)
}
