// Package for the terra-celestial localisation
package localise

import (
	"math"
	"time"
)

// Specified according to RFC3339 as a string to allow readability
const EpochTime = "2000-01-01T12:00:00.00Z"

// Geographical coordinates assume decimal format.
// TODO Add contversion of deg min sec to decimal in convert pkg
// East positive
// North positive
// GeographicalCoordinates are inn Decimal degrees (DD)
type GeographicalCoordinates struct {
	Latitude, Longitude float64
}

// CalculateCurrentSiderealTimeFromEpoch gives the current sidereal time in
// degrees at the given longitude as of time.Now().
func CalculateCurrentSiderealTimeFromEpoch(longitude float64) float64 {
	ut := time.Now().UTC()
	decimalDays := DecimalDaysSinceEpoch(ut)

	var t float64
	t = float64(ut.Hour())
	t += float64(ut.Minute()) / 60.0
	t += float64(ut.Second()) / 3600.0

	localSiderealTime := 100.46 + 0.985647*decimalDays
	localSiderealTime += longitude
	localSiderealTime += 15.0 * t

	localSiderealTime = math.Mod(localSiderealTime, 360)

	if localSiderealTime < 0 {
		localSiderealTime += 360
	}

	return localSiderealTime
}

func DecimalDaysSinceEpoch(t time.Time) float64 {

	epoch, _ := time.Parse(time.RFC3339, EpochTime)

	diff := t.Sub(epoch)

	decimalDays := diff.Hours() / 24

	return decimalDays

}
