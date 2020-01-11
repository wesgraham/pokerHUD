package behaviour

import (
	"fmt"
)

type State struct {
	hero     string
	villain  string
	handID   int
}

type StateData struct {
	heroData []byte
	villainData []byte
	textureData []byte
}

func Predict(state State) (string, error) {

	stateData, err := aggregate(state)
	if err != nil {
		return "", fmt.Errorf("error aggregating state: %s", err)
	}
	strength, err := strength(stateData)
	if err != nil {
		return "", fmt.Errorf("error determining strength: %s", err)
	}

	return strength, nil
}