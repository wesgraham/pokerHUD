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
	response, err := http.Get(query)
	if err != nil {
		return Profile{}, fmt.Errorf("error querying database: %s", err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Profile{}, fmt.Errorf("error reading response: %s", err)
	}

	stats, err := statsagg(data)
	if err != nil {
		return Profile{}, fmt.Errorf("error aggregating stats: %s", err)
	}

	return Profile{username, stats}, nil
}
