package behaviour

import(
	"fmt"
	"io/ioutil"
	"net/http"
)

type State struct {
	hero     string
	villain  string
	handID   int

}

func Predict(state State) (string, error) {
	heroQuery := "http://localhost:3000/hands?username=" + state.hero
	villainQuery := "http://localhost:3000/hands?username=" + state.villain
	textureQuery := "http://localhost:3000/hands?username=" + state.handID

	// TODO: Make this concurrent
	heroResponse, err := http.Get(heroQuery)
	if err != nil {
		return "", fmt.Errorf("error querying database: %s", err)
	}

	heroData, err := ioutil.ReadAll(heroResponse.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %s", err)
	}

	villainResponse, err := http.Get(villainQuery)
	if err != nil {
		return "", fmt.Errorf("error querying database: %s", err)
	}

	villainData, err := ioutil.ReadAll(villainResponse.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %s", err)
	}

	textureResponse, err := http.Get(textureQuery)
	if err != nil {
		return "", fmt.Errorf("error querying database: %s", err)
	}

	textureData, err := ioutil.ReadAll(textureResponse.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %s", err)
	}





}