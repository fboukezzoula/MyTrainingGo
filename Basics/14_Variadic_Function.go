package main

import "fmt"

func main() {

	computeTotal := add(5, 8, 8, 10, 26, 265)
	fmt.Println(computeTotal)

	computeMultiplication := multiplication(2.235, 2.32, 3.432124654564210, 10.23254E02)
	fmt.Println(computeMultiplication)
}

func add(mynumbers ...int) int {
	var total int
	for _, n := range mynumbers {
		total += n
	}
	return total
}

func multiplication(mynumbers ...float64) float64 {
	var total float64
	for _, n := range mynumbers {
		total += n
	}
	return total
}
