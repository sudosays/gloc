// Package for the conversion of units within and between celestial coordinate
// systems
package convert

// Converts units from sexagesimal to decimal equivalent.
func SexagesimalToDecimal(hours, minutes, seconds float64) float64 {

    return (hours + minutes/60 + seconds/3600)

}
