package main

import (
	"fmt"
	"os"
)

type Config struct {
    Database struct {
		Host string `json:"host"`
		Port string `json:"port"`	  
	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}


func LoadConfiguration(filename string) (Config, error) {
    var config Config
    configFile, err := os.Open(filename)
    defer configFile.Close()
    if err != nil {
        fmt.Println(err.Error())
 		return config, err
    }
    jsonParser := json.NewDecoder(configFile)
    jsonParser.Decode(&config)
    return config, err
}

func main () {

	fmt.Println("Starting reading json file and parse it ...")
	config, _= LoadConfiguration("config.json")
	fmt.Println(config)

}