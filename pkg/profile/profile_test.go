package profile

import (
	"encoding/json"
	"fmt"
	"github.com/wesgraham/pokerHUD/pkg/store"
	"github.com/wesgraham/pokerHUD/pkg/types"
	"testing"
)

func TestGet(t *testing.T) {
	profileBytes, err := Get("wg")
	if err != nil {
		t.Errorf("Get(wgraham): %s", err)
		return
	}

	var profile Profile
	err = json.Unmarshal(profileBytes, &profile)
	if err != nil {
		t.Errorf("unmarshal: %s", err)
		return
	}

	if profile.Username == "" {
		t.Errorf("Empty Username")
		return
	}

	entry := types.Hand{"wgraham", 1, 100000, "KJo", 2000, "call", 200, nil, false, false}
	store.Post(entry)

	profileBytes, err = Get("wgraham")
	if err != nil {
		t.Errorf("Get(wgraham): %s", err)
		return
	}

	err = json.Unmarshal(profileBytes, &profile)
	if err != nil {
		t.Errorf("unmarshal: %s", err)
		return
	}

	if profile.Username == "" {
		t.Errorf("Empty Username")
		return
	}

	if profile.Stats.Vpip == 0 {
		t.Errorf("Stats should be non-empty")
	}
}

func TestGetAll(t *testing.T) {
	profileArrayBytes, err := GetAll()
	if err != nil {
		t.Errorf("Get(wgraham): %s", err)
		return
	}

	var profileArray []Profile
	err = json.Unmarshal(profileArrayBytes, &profileArray)
	if err != nil {
		t.Errorf("unmarshal: %s", err)
		return
	}

	for i:=0; i<len(profileArray);i++ {
		fmt.Println(profileArray[i].Username)
		fmt.Println(profileArray[i].Stats)
	}

}