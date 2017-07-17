package main

import "fmt"

func main() {

	for i := 1; i < 10; i++ {
		fmt.Println(i)
		if i == 6 {
			break
		}

	}

	MyVarString := "Marseille"

	for _, char := range MyVarString {
		fmt.Print(string(char) + "-")
	}
	fmt.Println()

}
