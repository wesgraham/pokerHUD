package profile

import (
	"encoding/json"
	"fmt"
	"github.com/wesgraham/pokerHUD/pkg/store"
)

type Profile struct {
	Username string `json:"uname"`
	Stats    Stats  `json:"stats"`
}

type Stats struct {
	Vpip         int `json:"vpip"`
	Pfr          int `json:"pfr"`
	AverageRaise int `json:"averageRaise"`
}

func Get(username string) ([]byte, error) {
	// TODO: Remove need to hardcode query
	query := "http://localhost:3000/hands?uname=eq." + username
	data, err := store.Get(query)
	if err != nil {
		return nil, fmt.Errorf("error retrieving query data: %s", err)
	}

	stats, err := statsagg(data)
	if err != nil {
		return nil, fmt.Errorf("error aggregating stats: %s", err)
	}

	profileAsBytes, err := json.Marshal(Profile{username, stats})
	return profileAsBytes, nil
}
