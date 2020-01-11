package behaviour

import "fmt"

func aggregate(state State) (StateData, error) {
	heroQuery := "http://localhost:3000/hands?username=" + state.hero
	heroData, err := utils.QueryHandler(heroQuery)
	if err != nil {
		return StateData{}, fmt.Errorf("error retrieving query data: %s", err)
	}

	villainQuery := "http://localhost:3000/hands?username=" + state.villain
	villainData, err := utils.QueryHandler(villainQuery)
	if err != nil {
		return StateData{}, fmt.Errorf("error retrieving query data: %s", err)
	}

	textureQuery := "http://localhost:3000/hands?username=" + state.hero
	textureData, err := utils.QueryHandler(textureQuery)
	if err != nil {
		return StateData{}, fmt.Errorf("error retrieving query data: %s", err)
	}

	return StateData{heroData, villainData, textureData}, nil
}
