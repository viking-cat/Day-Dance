package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

type Config struct {
	Name string `json:"serverName"`
	Port uint   `json:"port"`
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path)
}

func getJson(file string) bool {
	log.Println("Reading " + file)

	rawFile, fileError := ioutil.ReadFile(file)
	if fileError != nil {
		log.Println(fileError)
		return false
	}

	var config Config

	marshalError := json.Unmarshal([]byte(rawFile), &config)
	if marshalError != nil {
		log.Println(marshalError)
		return false
	}

	log.Println("All: ", config)
	log.Println("port: ", config.Port)
	log.Println("port type: ", reflect.TypeOf(config.Port))
	log.Println("name type: ", config.Name)

	/*
			Fun experiments
			Golang really need a better JSON Unmarshal
			Maybe should write a library for fun
		// var jsonResult map[string]interface{}
		// https://stackoverflow.com/questions/28806951/accessing-nested-map-of-type-mapstringinterface-in-golang
		// https://gist.github.com/ChristopherThorpe/fd3720efe2ba83c929bf4105719ee967
		// https://irshadhasmat.medium.com/golang-simple-json-parsing-using-empty-interface-and-without-struct-in-go-language-e56d0e69968
		marshalError := json.Unmarshal([]byte(rawFile), &jsonResult)
		if marshalError != nil {
			log.Println(marshalError)
			return false
		}

		log.Println("port: ", jsonResult["port"])
		log.Println("port type: ", reflect.TypeOf(jsonResult["port"]))
		log.Println("name type: ", reflect.TypeOf(jsonResult["serverName"]))
		log.Println("All of it\n", jsonResult)
		log.Println("Final Experiment: ", jsonResult["something"])
		//thatThing := jsonResult["something"]
		//log.Println("Final Experiment: ", reflect.TypeOf(thatThing["abba"]))
	*/

	return true
}

func main() {

	// log.Println("Reading config.json")
	// serverConfig, err := os.Open("config.json")
	// if err != nil {
	//log.Println(err)
	//}

	if getJson("settings.json") {
		log.Println("success")
	} else {
		log.Println("failure")
	}

	// log.Printf("Starting Web Server")
	// http.HandleFunc("/", serveFile)
	// log.Fatal(http.ListenAndServe(":8081", nil))
}
