package behaviour

import (
	"fmt"
	"github.com/wesgraham/pokerHUD/pkg/store"
	"github.com/wesgraham/pokerHUD/pkg/types"
	"testing"
)

func TestPredict(t *testing.T) {
	entry := types.Hand{"foldy", 4, 100000, "KJo", 2000, "fold", 0, nil, false, false}
	data, err := store.Post(entry)
	if err != nil {
		t.Errorf("error posting to db: %s", err)
	}
	fmt.Println(data)

	entry = types.Hand{"raisy", 4, 100000, "KJo", 4000, "raise", 500, nil, true, false}
	_, err = store.Post(entry)
	if err != nil {
		t.Errorf("error posting to db: %s", err)
	}

	state := State{"foldy", "raisy", 4}
	ans, err := Predict(state)
	if err != nil {
		t.Errorf("Error predicting state: %s", err)
		return
	}
	fmt.Printf("Answer: %s", ans)
}
