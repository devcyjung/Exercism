// Package weather provides tools for weather forecast.
package weather

// CurrentCondition is the current condition.
var CurrentCondition string

// CurrentLocation is the current location.
var CurrentLocation string

//Forecast returns the current location and current condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
