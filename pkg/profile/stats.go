package profile

import (
	"encoding/json"
	"fmt"
	types ""
)


// TODO: Speedup by querying directly?
func getVPiP(data []byte) (int, error){
	var entryArray []types.HandEntry
	err := json.Unmarshal(data, &entryArray)
	if err != nil {
		return -1, fmt.Errorf("error unmarshalling data: %s", err)
	}

	voluntaryPuts := 0
	for i := 0; i < len(entryArray); i++ {
		if (entryArray[i].action == "call" || entryArray[i].action == "raise") && entryArray[i].board == nil {
			voluntaryPuts += 1
		}
	}

	vpip := voluntaryPuts / len(entryArray)
	return vpip, nil
}


func getPFR(data []byte) (int, error) {
	var entryArray []types.HandEntry
	err := json.Unmarshal(data, &entryArray)
	if err != nil {
		return -1, fmt.Errorf("error unmarshalling data: %s", err)
	}

	pfr := 0
	for i := 0; i < len(entryArray); i++ {
		if entryArray[i].action == "raise" && entryArray[i].board == nil {
			pfr += 1
		}
	}

	vpip := pfr / len(entryArray)
	return vpip, nil
}

func getAvgRaise(data []byte) (int, error) {
	var entryArray []types.HandEntry
	err := json.Unmarshal(data, &entryArray)
	if err != nil {
		return -1, fmt.Errorf("error unmarshalling data: %s", err)
	}

	raiseSum := 0
	raiseTot := 0
	for i := 0; i < len(entryArray); i++ {
		if entryArray[i].action == "raise" {
			raiseSum += entryArray[i].amount
			raiseTot += 1
		}
	}

	avgRaise := raiseSum/raiseTot
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