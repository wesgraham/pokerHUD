package behaviour

import (
	"fmt"
	"github.com/wesgraham/pokerHUD/pkg/store"
	"github.com/wesgraham/pokerHUD/pkg/types"
	"testing"
)

func TestPredict(t *testing.T) {
	entry := types.Hand{"wgraham", 1, 100000, "KJo", 2000, "call", 200, nil, false, false}
	_, err := store.Post(entry)
	if err != nil {
		t.Errorf("error posting to db: %s", err)
	}

	entry = types.Hand{"wgraham", 2, 100000, "KJo", 2000, "call", 200, nil, false, false}
	_, err = store.Post(entry)
	if err != nil {
		t.Errorf("error posting to db: %s", err)
	}

	entry = types.Hand{"hubertblack", 1, 100000, "KJo", 4000, "raise", 500, nil, true, false}
	_, err = store.Post(entry)
	if err != nil {
		t.Errorf("error posting to db: %s", err)
	}

	state := State{"wgraham", "hubertblack", 2}
	ans, err := Predict(state)
	if err != nil {
		t.Errorf("Error predicting state: %s", err)
		return
	}
	fmt.Printf("Answer: %s", ans)
}
