// Package weather provides tools to check for current weather condition.
package weather

var (
// CurrentCondition holds the current condition.
	CurrentCondition string
// CurrentLocation holds the location of the you want to check the weather condition.
	CurrentLocation  string
)
// Forecast returns a string value of the current condition of a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}