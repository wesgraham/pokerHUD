package profile

import (
	"encoding/json"
	"fmt"
	"github.com/wesgraham/pokerHUD/pkg/store"
	"github.com/wesgraham/pokerHUD/pkg/types"
)

// Profile represents a user and their stats
type Profile struct {
	Username string `json:"uname"`
	Stats    Stats  `json:"stats"`
}

// Stats represents core hud metrics
type Stats struct {
	Vpip         float32 `json:"vpip"`
	Pfr          float32 `json:"pfr"`
	AverageRaise int     `json:"averageRaise"`
}

// Get retrieves profile for a specific user
func Get(username string) ([]byte, error) {
	//TODO: Remove need to hardcode query
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

// GetAll retrieves profiles for all users
func GetAll() ([]byte, error) {
	query := "http://localhost:3000/hands"
	data, err := store.Get(query)
	if err != nil {
		return nil, fmt.Errorf("error retrieving query data: %s", err)
	}

	var hands []types.Hand
	err = json.Unmarshal(data, &hands)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling data: %s", err)
	}

	var profileArray []Profile
	var usernames []string
	for i := 0; i < len(hands); i++ {
		if !contains(usernames, hands[i].Uname) {
			query := "http://localhost:3000/hands?uname=eq." + hands[i].Uname
			data, err := store.Get(query)
			if err != nil {
				return nil, fmt.Errorf("error retrieving query data: %s", err)
			}

			stats, err := statsagg(data)
			if err != nil {
				return nil, fmt.Errorf("error aggregating stats: %s", err)
			}

			profileArray = append(profileArray, Profile{hands[i].Uname, stats})
			usernames = append(usernames, hands[i].Uname)
		}
	}

	profileArrayAsBytes, err := json.Marshal(profileArray)
	return profileArrayAsBytes, nil
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
