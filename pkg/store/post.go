package store

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/wesgraham/pokerHUD/pkg/types"
	"io/ioutil"
	"net/http"
)

// TODO: Get rid of this - figure out why hand.go isnt displaying latest changes
type FormattedHand struct {
	Uname       string   `json:"uname"`
	HandID      int      `json:"handid"`
	Balance     int      `json:"balance"`
	Hand        string   `json:"hand"`
	PotSize     int      `json:"potsize"`
	Action      string   `json:"action"`
	Amount      int      `json:"amount"`
	Board       []string `json:"board"`
	ThreeBet    bool     `json:"threebet"`
	FourPlusBet bool     `json:"fourplusbet"`
}


func Post(entry types.Hand) (string, error) {

	formatted := FormattedHand{entry.Uname, entry.HandID, entry.Balance, entry.Hand, entry.PotSize, entry.Action, entry.Amount, entry.Board, entry.ThreeBet, entry.FourPlusBet}

	requestBody, err := json.Marshal(formatted)
	if err != nil {
		return "", fmt.Errorf("error marshalling json: %s", err)
	}
	fmt.Println(requestBody)

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

