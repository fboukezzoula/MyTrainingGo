package main

import "fmt"

func main() {

	a, b := complexecalculate()
	fmt.Println(a + b)
}

func complexecalculate() (int, int) {

	e := 20
	f := 10
	return e, f

}
