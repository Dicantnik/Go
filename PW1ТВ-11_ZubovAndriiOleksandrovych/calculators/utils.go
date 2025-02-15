package calculators

import "fmt"

// roundValue повертає число типу float64 у вигляді рядка з заданою кількістю десяткових знаків.
func roundValue(num float64, decimals int) string {
	format := fmt.Sprintf("%%.%df", decimals)
	return fmt.Sprintf(format, num)
}
