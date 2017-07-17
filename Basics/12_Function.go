package main

import "fmt"

func main() {

	affiche()
	compute(6, 8, 25)

}

func affiche() {
	fmt.Println("The main function has called this affiche function !")

}

func compute(a int, b int, c int) {
	fmt.Println(a + b + c)

}
