package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "{{ }}"

	strArray := strings.Fields(str)
	
	lenght := len(strArray)

	
	fmt.Println(lenght)
	
	if lenght == 2 {
		fmt.Println("Your delimiters definition & placeholders have this format : ", strArray[0],"KeyName",strArray[1]  )
	} else {
		fmt.Println("ko")
	}
	
	
	fmt.Println(strArray[0])
	fmt.Println(strArray[1])

}
