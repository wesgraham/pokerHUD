package behaviour

import (
	"fmt"
	"github.com/wesgraham/pokerHUD/pkg/store"
)

func aggregateStateData(state State) (StateData, error) {
	heroQuery := "http://localhost:3000/hands?uname=eq." + state.hero
	heroData, err := store.Get(heroQuery)
	if err != nil {
		return StateData{}, fmt.Errorf("error retrieving query data: %s", err)
	}

	villainQuery := "http://localhost:3000/hands?uname=eq." + state.villain
	villainData, err := store.Get(villainQuery)
	if err != nil {
		return StateData{}, fmt.Errorf("error retrieving query data: %s", err)
	}

	textureQuery := "http://localhost:3000/hands?uname=eq." + state.hero
	textureData, err := store.Get(textureQuery)
	if err != nil {
		return StateData{}, fmt.Errorf("error retrieving query data: %s", err)
	}

	return StateData{heroData, villainData, textureData}, nil
}
