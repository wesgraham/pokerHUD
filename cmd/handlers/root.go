package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

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
