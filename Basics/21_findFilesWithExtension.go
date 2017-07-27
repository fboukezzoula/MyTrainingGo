/* Go Lang find files with extension from the current working directory.
Copyright (c) 2010-2014 Alex Niderberg */

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {

	fmt.Println("Search PlaceHolders on these extensions files: ", os.Getenv("PAAS_DOCKER_EXTENSIONS"))

	var folders [10]string

	folders[0] = "P:/test"
	folders[1] = "P:/GoWorkspace/src/My_Go_Learning"

	fmt.Println(folders)

	filesxextensions := strings.Split(os.Getenv("PAAS_DOCKER_EXTENSIONS"), ",")

	for _, extension := range filesxextensions {
		fmt.Println(">> all files with the extension :", extension, "dsdqs")
		fmt.Println(checkExt(extension))
	}

}

func checkExt(ext string) []string {

	/*pathS, err := os.Getwd()
	if err != nil {
		panic(err)
	}*/
	var files []string
	filepath.Walk("P:/test", func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(ext, f.Name())
			if err == nil && r {
				files = append(files, f.Name())
			}
		}
		return nil
	})
	return files
}
