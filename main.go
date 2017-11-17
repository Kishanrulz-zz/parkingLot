package main

import (

	"os"
	simulator "parking/simulator"
)


func main() {
	if len(os.Args) == 2 {
		simulator.SimulateParkingFromFile(os.Args[1])
	} else {
		simulator.SimulateParkingInteractive()
	}
}



