package main

import "fmt"

func main() {

	add(5, 8, 8, 10, 26, 265)
}

func add(mynumbers ...int) {

	var total int

	for _, n := range mynumbers {
		total += n
	}
	fmt.Println(total)
}
