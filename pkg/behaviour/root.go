package behaviour

import (
	"fmt"
)

// Used to collect data on the state of a board and the players in question
type State struct {
	hero    string
	villain string
	handID  int
}

// Represents data on state of the board and the players in question
type StateData struct {
	heroData    []byte
	villainData []byte
	textureData []byte
}

// Predicts a users next action given a current state
func Predict(state State) (string, error) {
	stateData, err := aggregateStateData(state)
	if err != nil {
		return "", fmt.Errorf("error aggregating state: %s", err)
	}
	strength, err := getStrength(stateData)
	if err != nil {
		return "", fmt.Errorf("error determining strength: %s", err)
	}

	return strength, nil
}
