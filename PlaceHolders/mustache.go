package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/hoisie/mustache"
)

var commits = map[string]string{}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	_, err := fmt.Println("@@@@@@@@ Search PlaceHolders on these extensions files: ", os.Getenv("PAAS_DOCKER_EXTENSIONS"))
	check(err)

	filesxextensions := strings.Split(os.Getenv("PAAS_DOCKER_EXTENSIONS"), ",")

	folders := os.Getenv("PAAS_DOCKER_FOLDERS")
	cleaned := strings.Replace(folders, ",", " ", -1)

	// convert 'clened' comma separated string to slice
	strSlice := strings.Fields(cleaned)

	for i := 0; i <= 1; i++ {
		fmt.Println("\n... Searching in :", strSlice[i], "folder ...")
		for _, extension := range filesxextensions {
			fmt.Println("\nAll files with the extension : >>>", extension, "<<<\n------------------------------------------------------")
			_, err := fmt.Println(checkExt(extension, strSlice[i]))
			check(err)
		}
	}

	data := mustache.Render("hello {{c}}", map[string]string{"c": "world"})
	println(data)

	//func Render(data string, context ...interface{}) string
	//func RenderFile(filename string, context ...interface{}) string

	datafile := mustache.RenderFile("P:/GoWorkspace/src/My_Go_Learning/PlaceHolder.txt", map[string]string{"om": "psg"})
	println(datafile)

	commits["key"] = "thing_value"
	commits["om"] = "psg"

	for key, value := range commits {
		fmt.Println("Key:", key, "Value:", value)
	}

}

func checkExt(ext, rootparent string) []string {
	var files []string
	filepath.Walk(rootparent, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString("."+ext, f.Name())
			if err == nil && r {
				files = append(files, f.Name())
			}
		}
		return nil
	})
	return files
}
