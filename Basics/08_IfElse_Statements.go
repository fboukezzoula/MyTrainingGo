package main

import "fmt"

func main() {

	condition := 5 < 3 //return bool false
	fmt.Println(condition)

	if condition {
		fmt.Println("En effet, 5 est plus grand que 3 !")
	} else {

		fmt.Println("Never ! 3 n'est pas plus grand que 5 .. voyons !!!")
	}

	// AND &&

	if 6 > 3 && 2 < 1 {
		fmt.Println("Les 2 conditions sont respectées !")
	} else {
		fmt.Println("Une des 2 conditions n'est pas respectée !")
	}

	// OR ||
	if 6 > 3 || 2 < 1 {
		fmt.Println("Une des 2 conditions sont respectées !")
	} else {
		fmt.Println("Aucune des 2 conditions n'est pas respectée !")
	}

}
