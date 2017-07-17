package main

import "fmt"

func main() {

	var MyNewVal int = 5

	MyVar1, MyVar2, Levier := 10, 20, true

	fmt.Println(MyVar1)
	fmt.Println(MyVar2)

	fmt.Println("-------MODIFIED--------------")

	MyVar1, MyVar2 = 11, 22

	fmt.Println(MyVar1)
	fmt.Println(MyVar2)

	fmt.Println(MyVar2 / MyVar1)

	fmt.Println(MyNewVal)

	if Levier {
		fmt.Println("marseille")
	} else {
		fmt.Println("psg")
	}
}
