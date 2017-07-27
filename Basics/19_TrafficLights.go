// An example of declaring a constant grouping with an iota enumerator.
package main

import (
	"fmt"
)

const (
	_ = iota
	TrafficLightStateRedLight
	TrafficLightStateYellowLight
	TrafficLightStateGreenLight
)

func main() {

	fmt.Println("Red Light State Code: ", TrafficLightStateRedLight)
	fmt.Println("Green Light State Code: ", TrafficLightStateGreenLight)
	fmt.Println("Yellow Light State Code: ", TrafficLightStateYellowLight)
}
