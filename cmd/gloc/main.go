package main

import (
    "fmt"
    "github.com/sudosays/gloc/pkg/convert"
)

func main() {

    fmt.Printf("hello, universe!\n")
    result := convert.SexagesimalToDecimal(2,30,45)
    fmt.Printf("2h30m45s converted to decimal is: %f\n", result )

}
