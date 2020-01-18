package profile

import (
	"encoding/json"
	"fmt"
	"github.com/wesgraham/pokerHUD/pkg/types"
)

func getVPiP(data []byte) (float32, error) {
	var entryArray []types.Hand
	err := json.Unmarshal(data, &entryArray)
	if err != nil {
		return -1, fmt.Errorf("error unmarshalling data: %s", err)
	}

	voluntaryPuts := 0
	for i := 0; i < len(entryArray); i++ {
		if entryArray[i].Action == "call" || entryArray[i].Action == "raise" {
			voluntaryPuts += 1
		}
	}

	vpip := float32(voluntaryPuts) / float32(max(len(entryArray), 1))
	return vpip, nil
}

func getPFR(data []byte) (float32, error) {
	var entryArray []types.Hand
	err := json.Unmarshal(data, &entryArray)
	if err != nil {
		return -1, fmt.Errorf("error unmarshalling data: %s", err)
	}

	pfr := 0
	for i := 0; i < len(entryArray); i++ {
		if entryArray[i].Action == "raise" {
			pfr += 1
		}
	}

	vpip := float32(pfr) / float32(max(len(entryArray), 1))
	return vpip, nil
}

func getAvgRaise(data []byte) (int, error) {
	var entryArray []types.Hand
	err := json.Unmarshal(data, &entryArray)
	if err != nil {
		return -1, fmt.Errorf("error unmarshalling data: %s", err)
	}

	raiseSum := 0
	raiseTot := 0
	for i := 0; i < len(entryArray); i++ {
		if entryArray[i].Action == "raise" {
			raiseSum += entryArray[i].Amount
			raiseTot += 1
		}
	}

	avgRaise := raiseSum / max(raiseTot, 1)
	return avgRaise, nil
}

func statsagg(data []byte) (Stats, error) {
	vpip, err := getVPiP(data)
	if err != nil {
		return Stats{}, fmt.Errorf("error retrieving vpip: %s", err)
	}

	pfr, err := getPFR(data)
	if err != nil {
		return Stats{}, fmt.Errorf("error retrieving pfr: %s", err)
	}

	avgraise, err := getAvgRaise(data)
	if err != nil {
		return Stats{}, fmt.Errorf("error retrieving avgraise: %s", err)
	}

	stats := Stats{vpip, pfr, avgraise}
	return stats, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
