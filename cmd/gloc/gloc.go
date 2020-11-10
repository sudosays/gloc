package main

import (
	"flag"
	"fmt"
	"time"
	//    "github.com/sudosays/gloc/pkg/convert"
	"github.com/sudosays/gloc/pkg/localise"
)

var interactive bool

func init() {
	flag.BoolVar(&interactive, "i", false, "Run in interactive mode, continuously updating calculations")
}

func main() {
	longitude := 22.45
	//latitude := -33.966667
	flag.Parse()
	fmt.Printf("hello, universe!\n\n")
	localSiderealTime := localise.CalculateCurrentSiderealTimeFromEpoch(longitude)
	if interactive {
		for {
			EraseOutputLine()
			localSiderealTime = localise.CalculateCurrentSiderealTimeFromEpoch(longitude)
			fmt.Println("Current sidereal time is ", localSiderealTime, " degrees.")
			time.Sleep(50 * time.Millisecond)
		}
	}

}

func EraseOutputLine() {
	fmt.Print("\u001b[F\u001b[0K")
}
