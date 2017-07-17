package main

import (
	"fmt"
)

func main() {

	const myConstante float64 = 6.142562
	fmt.Println(myConstante)

	const (
		myFirst = iota
		mySecond
		myThird
	)

	fmt.Println(myFirst, mySecond, myThird)
}
