package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)

func readFile(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
	return text
}
var json = readFile("./users.json")

func users(w http.ResponseWriter, req *http.Request) {	
	fmt.Fprintf(w, json)
}

func main() {
	http.HandleFunc("/users", users)
	http.ListenAndServe(":80", nil)
}