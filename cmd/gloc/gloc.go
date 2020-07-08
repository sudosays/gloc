package main

import (
    "fmt"
    "flag"
    "time"
//    "github.com/sudosays/gloc/pkg/convert"
    "github.com/sudosays/gloc/pkg/localise"
)

var interactive bool

func init() { 
    flag.BoolVar(&interactive, "i", false, "Run in interactive mode, continuously updating calculations")
}

func main() {
    
    flag.Parse()

    fmt.Printf("hello, universe!\n")

    localise.CalculateCurrentSiderealTimeFromEpoch()
    
    if (interactive) {
        for {
            fmt.Print("\u001b[F\u001b[0K")
            localise.CalculateCurrentSiderealTimeFromEpoch()
            time.Sleep(50 * time.Millisecond)
        }
    }
}
