package main

import (
  "fmt"
  "strings"
)

func main() {
  
  str:= "pomme poire orange abricot"
  
  strArray:= strings.Fields(str)
  
  fmt.Println(strArray)
  fmt.Println(strArray[1:3])


}
