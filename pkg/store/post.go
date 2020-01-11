package store

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/wesgraham/pokerHUD/pkg/types"
	"io/ioutil"
	"net/http"
)

func Post(entry types.Hand) (string, error){

	requestBody, err := json.Marshal(entry)
	if err != nil {
		return "", fmt.Errorf("error marshalling json: %s", err)
	}

	resp, err := http.Post("http://localhost:3000/hands", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("error posting: %s", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error retreiving response: %s", err)
	}

	return string(body), nil
}
