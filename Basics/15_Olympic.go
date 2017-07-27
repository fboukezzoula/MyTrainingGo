// An example of using Go's iota enumerator to model the year and city where the summer Olympics took place.
package main

import "fmt"

func main() {

	const (
		LosAngeles = 1984 + (iota * 4)
		Seoul
		Barcelona
		Atlanta
		Sydney
		Athens
		Beijing
		London
		Rio
		Tokyo
	)

	fmt.Println("These cities hosted or will host the Summer Olympics in the provided year...")
	fmt.Printf("%-38s %-38v \n", "City", "Year")
	fmt.Printf("%-38s %-38v \n", "Los Angeles", LosAngeles)
	fmt.Printf("%-38s %-38v \n", "Atlanta", Atlanta)
	fmt.Printf("%-38s %-38v \n", "Tokyo", Tokyo)

}
