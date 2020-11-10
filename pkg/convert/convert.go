// Package for the conversion of units within and between celestial coordinate
// systems
package convert

import (
    "math"
    "time"
)

type Declination struct {
    Degree, Arcminute, Arcsecond, Posneg float64
}

type RightAscension struct {
    Hour, Minute, Second float64
}

type CelestialCoordinates struct {
    Ra RightAscension
    Dec Declination
}

// Converts units from sexagesimal to decimal equivalent.
func DeclinationToDecimal(n Declination) float64 {
    return (n.Degree + n.Posneg*n.Arcminute/60 + n.Posneg*n.Arcsecond/3600)
}

// Converts units from decimal to sexagesimal equivalent.
func DecimalToDeclination(d float64) Declination {
    var posneg float64
    if d >= 0 {
        posneg = 1.0
    } else {
        posneg = -1.0
    }

    degrees := math.Trunc(d)
    arcminutes := math.Trunc((d - degrees*posneg)*60.0*posneg)
    arcseconds := math.Trunc((d - posneg*degrees - posneg*arcminutes/60.0)*3600.0*posneg)
    return Declination{degrees, arcminutes, arcseconds, posneg}
}

// Converts units from right ascension to decimal
func RightAscensionToDecimal(r RightAscension) float64 {
    return (r.Hour*15.0 + r.Minute/4.0 + r.Second/240.0)
}

// Converts units from decimal to right ascension 
func DecimalToRightAscension(d float64) RightAscension {
    hours := math.Trunc(d/15.0)
    minutes := math.Trunc((d - hours*15.0)*4.0)
    seconds := math.Trunc((d - hours*15.0 - minutes/4.0)*240.0)

    return RightAscension{hours, minutes, seconds}
}

func ToAltAzimuth(coordinates CelestialCoordinates, observationLocation ) {


}
