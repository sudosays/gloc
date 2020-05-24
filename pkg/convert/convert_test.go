package convert

import(
    "testing"
)

func TestConvertSexagesimalToDecimal(t *testing.T) {

    got := SexagesimalToDecimal(2,30,45)
    if got != 2.5125 {
        t.Errorf("SexagesimalToDecimal(2,30,45) = %f; want 2.5123", got)
    }
}
