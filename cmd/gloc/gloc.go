package main

import (
    "fmt"
    "flag"
    "github.com/sudosays/gloc/pkg/convert"
    "github.com/sudosays/gloc/pkg/localise"
)

var interactive bool

func init() { 
    flag.BoolVar(&interactive, "i", false, "Run in interactive mode, continuously updating calculations")
}

func main() {
    
    flag.Parse()

    fmt.Printf("hello, universe!\n")
    s := convert.RightAscension{2.0,30.0,40.0}

    result := convert.RightAscensionToDecimal(s)

    fmt.Printf("RA 2h30m45s converted to decimal is: %f\n", result )
    
    if (interactive) {
        for {
            localise.CalculateCurrentSiderealTimeFromEpoch()
        }
    }
}
