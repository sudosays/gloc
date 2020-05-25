package convert

import(
    "testing"
)


func TestConvertDeclinationToDecimal(t *testing.T) {
    s := Declination{2,30,45,1}
    want := 2.5125
    got := DeclinationToDecimal(s)
    if got != want {
        t.Errorf("DeclinationToDecimal(2,30,45) = %f; want %f", got, want)
    }
}

func TestConvertDecimalToDeclination(t *testing.T) {
    got := DecimalToDeclination(2.5125)
    want := Declination{2, 30, 45, 1}
    if got != want {
        t.Errorf("DecimalToDeclination(2.5125) = %f; want %f", got, want)
    }
}

func TestConvertRightAscensionToDecimal(t *testing.T) {
    r := RightAscension{1,37,25}
    want := 24.354167
    got := RightAscensionToDecimal(r)
    if got != want {
        t.Errorf("Converting RightAscension to Decimal RA(1,37,25) = %f; want %f", got, want)
    }

    r = RightAscension{23,59,59}
    want = 359.9958
    got = RightAscensionToDecimal(r)
    if got != want {
        t.Errorf("Converting RightAscension to Decimal RA(23,59,59) = %f; want %f", got, want)
    }
}

func TestConvertDecimalToRightAscension(t *testing.T) {
    d := 24.354167
    want := RightAscension{1,37,25}
    got := DecimalToRightAscension(d)
    if got != want {
        t.Errorf("Converting DecimalToRightAscension(%f)= %f; want %f", d, got, want)
    }
}
