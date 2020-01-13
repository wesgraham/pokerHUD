package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Load file of homepage with basic HTML
	pwd, _ := os.Getwd()
	dat, err := ioutil.ReadFile(pwd + "/app/home.html")
	if err != nil {
		fmt.Printf("error: %s", err)
	}

	fmt.Fprintf(w, string(dat))
}


func otherHandler(w http.ResponseWriter, r *http.Request) {
	// Load file of homepage with basic HTML
	pwd, _ := os.Getwd()
	dat, err := ioutil.ReadFile(pwd + "/app/profile.html")
	if err != nil {
		fmt.Printf("error: %s", err)
	}

	fmt.Fprintf(w, string(dat))
}


func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/profile", otherHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
