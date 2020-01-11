package store

import (
"fmt"
"io/ioutil"
"net/http"
)

func Get(query string) ([]byte, error){
	response, err := http.Get(query)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %s", err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %s", err)
	}

	return data, nil
}
