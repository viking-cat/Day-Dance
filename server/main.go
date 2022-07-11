package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

type ServerConfig struct {
	Name string `json:"serverName"`
	Port uint   `json:"port"`
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path)
}

func getJson(file string) (*ServerConfig, error) {
	log.Println("Reading " + file)

	rawFile, fileError := ioutil.ReadFile(file)
	if fileError != nil {
		log.Println(fileError)
		return nil, fileError
	}

	var config ServerConfig

	marshalError := json.Unmarshal([]byte(rawFile), &config)
	if marshalError != nil {
		log.Println(marshalError)
		return nil, marshalError
	}

	log.Println("All: ", config)
	log.Println("port: ", config.Port)
	log.Println("port type: ", reflect.TypeOf(config.Port))
	log.Println("name type: ", config.Name)

	return &config, nil
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got \"/\" request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got \"/hello\" request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {

	// Get JSON config
	config, jsonError := getJson("settings.json")
	if jsonError != nil {
		log.Println(jsonError)
	}

	// Define Routes
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	// Cast uint to string
	port := ":" + fmt.Sprint(config.Port)

	// Start server listener
	serverError := http.ListenAndServe(port, nil)
	if jsonError != serverError {
		log.Println(serverError)
	}
}
