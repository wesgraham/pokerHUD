package main

import (
	"log"
	"net/http"
)

// TODO: Get rid of localhost everywhere
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/getProfile", profileHandler)
	http.HandleFunc("/profiles", profilesHandler)
	http.HandleFunc("/getStats", statHandler)
	http.HandleFunc("/stats", statsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
