package main

import (
	"fmt"
	"github.com/wesgraham/pokerHUD/pkg/profile"
	"github.com/wesgraham/pokerHUD/pkg/store"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// TODO: Modularize Handlers
// TODO: Get rid of localhost everywhere
func handler(w http.ResponseWriter, r *http.Request) {
	// Load file of homepage with basic HTML
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	pwd, _ := os.Getwd()
	dat, err := ioutil.ReadFile(pwd + "/app/home.html")
	if err != nil {
		fmt.Printf("error: %s", err)
	}

	fmt.Fprintf(w, string(dat))
}


func profilesHandler(w http.ResponseWriter, r *http.Request) {
	// Load file of homepage with basic HTML
	pwd, _ := os.Getwd()
	dat, err := ioutil.ReadFile(pwd + "/app/profile.html")
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	fmt.Fprintf(w, string(dat))
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	query := "http://localhost:3000/hands"
	data, err := store.Get(query)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of birds to the response
	w.Write(data)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	// Load file of homepage with basic HTML
	pwd, _ := os.Getwd()
	dat, err := ioutil.ReadFile(pwd + "/app/stats.html")
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	fmt.Fprintf(w, string(dat))
}

func statHandler(w http.ResponseWriter, r *http.Request) {
	data, err := profile.Get("wg")

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of birds to the response
	w.Write(data)
}


func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/getProfile", profileHandler)
	http.HandleFunc("/profiles", profilesHandler)
	http.HandleFunc("/getStats", statHandler)
	http.HandleFunc("/stats", statsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
