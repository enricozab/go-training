// Package weather implements the forecast based on the current condition and location.
package weather

// CurrentCondition stores the current condition value.
var CurrentCondition string

// CurrentLocation stores the current location value.
var CurrentLocation string

// Forecast returns the forecast based on the parameters: city, location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
