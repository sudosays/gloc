// Package for the terra-celestial localisation 
package localise

import (
    "fmt"
    "time"
)

// Specified according to RFC3339 as a string to allow readability
const EpochTime = "2000-01-01T12:00:00.00Z"

// Geographical coordinates assume decimal format. 
// TODO Add contversion of deg min sec to decimal in convert pkg
type GeographicalCoordinates struct {
    Latitude, Lontitude float64
}

func CalculateCurrentSiderealTimeFromEpoch() {
    epoch, _ := time.Parse(time.RFC3339, EpochTime)
    t := time.Now()

    diff := t.Sub(epoch)

    fmt.Println("Time since epoch J2000: ", diff)
}
