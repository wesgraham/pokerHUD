package store

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/wesgraham/pokerHUD/pkg/types"
	"io/ioutil"
	"net/http"
)

func Post(entry types.Hand) (string, error) {

	requestBody, err := json.Marshal(entry)
	fmt.Println(requestBody)
	fmt.Println(string(requestBody))
	if err != nil {
		return "", fmt.Errorf("error marshalling json: %s", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:3000/hands", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("error posting: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return string(body), nil
}
