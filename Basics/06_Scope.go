package main

import "fmt"

var iron int = 10

func main() {

	gold := 20

	fmt.Println(iron)
	fmt.Println(gold)

	anotherfunction()

}

func anotherfunction() {

	fmt.Println(iron)
	//fmt.Println(gold)

}
