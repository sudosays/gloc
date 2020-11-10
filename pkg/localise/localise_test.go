package localise

import (
    "testing"
    "time"
    "math"
)

const errorTolerance = 0.001

func TestSiderealTimeCalculation(t *testing.T) {}

func TestDecimalDaysCalculation(t *testing.T) {

    testTime, _ := time.Parse(time.RFC3339,"1998-08-10T23:10:00Z")
    expectedAnswer := -508.534722
    actualAnswer := DecimalDaysSinceEpoch(testTime)

    if math.Abs(actualAnswer - expectedAnswer) > errorTolerance {
        t.Errorf("DecimalDaysSinceEpoch = %f; want %f", actualAnswer, expectedAnswer)
    }

}
