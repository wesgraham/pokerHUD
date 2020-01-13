package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Root)
	http.HandleFunc("/getProfile", Profile)
	http.HandleFunc("/profiles", Profiles)
	http.HandleFunc("/getStats", Stat)
	http.HandleFunc("/stats", Stats)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
