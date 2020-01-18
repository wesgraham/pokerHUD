package main

import (
	"encoding/json"
	"fmt"
	"github.com/wesgraham/pokerHUD/pkg/profile"
	"github.com/wesgraham/pokerHUD/pkg/store"
	"github.com/wesgraham/pokerHUD/pkg/types"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func Root(w http.ResponseWriter, r *http.Request) {
	// Load file of homepage with basic HTML
	enableCors(&w)
	pwd, _ := os.Getwd()
	dat, err := ioutil.ReadFile(pwd + "/app/home.html")
	if err != nil {
		fmt.Printf("error: %s", err)
	}

	fmt.Fprintf(w, string(dat))
}

func Profiles(w http.ResponseWriter, r *http.Request) {
	// Load file of homepage with basic HTML
	pwd, _ := os.Getwd()
	dat, err := ioutil.ReadFile(pwd + "/app/profile.html")
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	fmt.Fprintf(w, string(dat))
}

func Profile(w http.ResponseWriter, r *http.Request) {
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

func Stats(w http.ResponseWriter, r *http.Request) {
	// Load file of homepage with basic HTML
	pwd, _ := os.Getwd()
	dat, err := ioutil.ReadFile(pwd + "/app/stats.html")
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	fmt.Fprintf(w, string(dat))
}

func Stat(w http.ResponseWriter, r *http.Request) {
	data, err := profile.GetAll()
	fmt.Println(data)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// If all goes well, write the JSON list of profiles to the response
	w.Write(data)
}

func Post(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	body, err := ioutil.ReadAll(r.Body)
	var hand types.Hand
	err = json.Unmarshal(body, &hand)
	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := store.Post(hand)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// If all goes well, write the JSON list of birds to the response
	fmt.Fprintf(w, resp)
}

func main() {
	http.HandleFunc("/", Root)
	http.HandleFunc("/getProfile", Profile)
	http.HandleFunc("/profiles", Profiles)
	http.HandleFunc("/getStats", Stat)
	http.HandleFunc("/stats", Stats)
	http.HandleFunc("/post", Post)
	fmt.Printf("Listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
