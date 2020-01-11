package profile

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

	query := "http://localhost:3000/hands?username=" + username
	data, err := utils.QueryHandler(query)
	if err != nil {
		return Profile{}, fmt.Errorf("error retrieving query data: %s", err)
	}

	stats, err := statsagg(data)
	if err != nil {
		return Profile{}, fmt.Errorf("error aggregating stats: %s", err)
	}

	return Profile{username, stats}, nil
}
