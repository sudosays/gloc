package main

import (
    "fmt"
    "github.com/sudosays/gloc/pkg/convert"
)

func main() {

    fmt.Printf("hello, universe!\n")
    s := convert.RightAscension{2,30,40}

    result := convert.RightAscensionToDecimal(s)

    fmt.Printf("RA 2h30m45s converted to decimal is: %f\n", result )

}
