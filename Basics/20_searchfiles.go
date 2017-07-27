package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	searchFiles("P:/test")

	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)

}

func searchFiles(dir string) { // dir is the parent directory you what to search
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
