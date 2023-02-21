package conv

// Konverterer Kelvin til Celsius
func KelvinToCelsius(value float64) float64 {
	return (value + 273.15)
}

// Konverterer Celsius til Kelvin
func CelsiusToKelvin(value float64) float64 {
	return (value - 273.15)
}

// Konverterer Celsius til Farhenheit
func CelsiusToFarhenheit(value float64) float64 {
	return (value*(9/5) + 32)
}

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	return (value - 32) * (5 / 9)
}

// Konverterer Kelvin til Farhenheit
func KelvinToFarhenheit(value float64) float64 {
	return (value - 273.15*(9/5) + 32)
}

// Konverterer Farhenheit til Kelvin
func FarhenheitToKelvin(value float64) float64 {
	return (value - 32*(5/9) + 273.15)
}
