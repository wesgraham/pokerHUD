package profile

import (
	"fmt"
	"github.com/wesgraham/pokerHUD/pkg/store"
	)

type Profile struct {
	username     string
	stats		 Stats
}

type Stats struct {
	vpip         int
	pfr          int
	averageRaise int
}

func Get(username string) (Profile, error) {
	// TODO: Remove need to hardcode query
	query := "http://localhost:3000/hands?username=eq." + username
	data, err := store.Get(query)
	if err != nil {
		return Profile{}, fmt.Errorf("error retrieving query data: %s", err)
	}

	stats, err := statsagg(data)
	if err != nil {
		return Profile{}, fmt.Errorf("error aggregating stats: %s", err)
	}

	return Profile{username, stats}, nil
}
