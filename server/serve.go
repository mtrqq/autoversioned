package main

import (
	"log"
	"net/http"
)

var version = "1.1.0"

func getVersion(w http.ResponseWriter, req *http.Request) {
	log.Printf("Responding with version")
	w.Write([]byte(version))
}

func main() {
	log.Printf("Starting http server at port 8080")
	http.HandleFunc("/", getVersion)
	http.ListenAndServe(":8080", nil)
}
