package handlers

import (
	"fmt"
	"github.com/wesgraham/pokerHUD/pkg/profile"
	"io/ioutil"
	"net/http"
	"os"
)

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
