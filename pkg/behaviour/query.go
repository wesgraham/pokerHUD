package behaviour

import (
	"fmt"
	"net/http"
)

func query(input string, c chan *http.Response) (bool, error) {
	response, err := http.Get(input)
	if err != nil {
		return false, fmt.Errorf("error querying database: %s", err)
	}
	c <- response
	return true, nil
}
