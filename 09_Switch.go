package main

import "fmt"

func main() {

	myVariable := "create"

	switch myVariable {

	case "create":
		fmt.Println("Creating ...")
	case "open":
		fmt.Println("Openning ...")
	case "close":
		fmt.Println("Closing ...")

	default:
		fmt.Println("Unrecognised command ...")

	}

}
