package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Root(w http.ResponseWriter, r *http.Request) {
	// Load file of homepage with basic HTML
	pwd, _ := os.Getwd()
	dat, err := ioutil.ReadFile(pwd + "/app/home.html")
	if err != nil {
		fmt.Printf("error: %s", err)
	}

	fmt.Fprintf(w, string(dat))
}
