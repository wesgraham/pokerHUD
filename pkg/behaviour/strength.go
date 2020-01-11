package behaviour

import (
	"encoding/json"
	"fmt"
	"github.com/wesgraham/pokerHUD/pkg/types"
)

func strength(data StateData) (string, error) {

	var heroEntries []types.HandEntry
	err := json.Unmarshal(data.heroData, &heroEntries)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling data: %s", err)
	}

	var villainEntries []types.HandEntry
	err = json.Unmarshal(data.villainData, &villainEntries)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling data: %s", err)
	}

	var textureEntries []types.HandEntry
	err = json.Unmarshal(data.villainData, &textureEntries)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling data: %s", err)
	}

	var heroAgg string
	numraise, numcall, numfold := 0, 0, 0
	for i := 0; i < len(heroEntries); i++ {
		if len(heroEntries[i].Board) == len(textureEntries[len(textureEntries)-1].Board) {
			switch heroEntries[i].Action {
			case "fold":
				numfold += 1
			case "call":
				numcall += 1
			case "raise":
				numraise += 1
			}
		}
	}
	if numraise > numcall && numraise > numfold {
		heroAgg = "aggressive"
	} else if numcall > numraise && numcall > numfold {
		heroAgg = "caller"
	} else {
		heroAgg = "folder"
	}

	var villainAgg string
	numraise, numcall, numfold = 0, 0, 0
	for i := 0; i < len(villainEntries); i++ {
		if len(villainEntries[i].Board) == len(textureEntries[len(textureEntries)-1].Board) {
			switch villainEntries[i].Action {
			case "fold":
				numfold += 1
			case "call":
				numcall += 1
			case "raise":
				numraise += 1
			}
		}
	}
	if numraise > numcall && numraise > numfold {
		villainAgg = "aggressive"
	} else if numcall > numraise && numcall > numfold {
		villainAgg = "caller"
	} else {
		villainAgg = "folder"
	}

	// TODO: Determine metric for texture agg
	textureAgg := "TBD"

	ans := "hero is:" + heroAgg + "villain is:" + villainAgg + "texture is:" + textureAgg
	return ans, nil
}
