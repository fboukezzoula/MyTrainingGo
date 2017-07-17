package main

import "fmt"

func main() {

	var j int
	fmt.Println(j)

	condition := j > 5
	fmt.Println(condition)

	if condition {
		fmt.Println("golang")
		j++
	}
}
