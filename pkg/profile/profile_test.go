package profile

import (
	"github.com/wesgraham/pokerHUD/pkg/store"
	"github.com/wesgraham/pokerHUD/pkg/types"
	"testing"
)

func TestGet(t *testing.T) {
	profile, err := Get("wgraham")
	if err != nil {
		t.Errorf("Get(wgraham): %s", err)
	}
	if profile.username == "" {
		t.Errorf("Empty Username")
	}

	entry := types.Hand{"wgraham", 1, 100000, "KJo", 2000, "call", 200, nil, false, false }
	store.Post(entry)

	profile, err = Get("wgraham")
	if err != nil {
		t.Errorf("Get(wgraham): %s", err)
	}
	if profile.username == "" {
		t.Errorf("Empty Username")
	}

	if profile.stats.vpip == 0 {
		t.Errorf("Stats should be non-empty")
	}
}