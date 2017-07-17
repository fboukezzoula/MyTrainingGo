package main

import "fmt"

func main() {

	mystring1 := "Marseille champions project !"
	mystring2 := " PSG en demi cette année ?"

	fmt.Println(mystring1 + mystring2)

	fmt.Println("Vive le GO !!" + mystring1 + mystring2 + " LILLE en 2eme division ?")

	mySliceChars := mystring1[0]
	fmt.Println(string(mySliceChars))

}
