package store

import (
	"github.com/wesgraham/pokerHUD/pkg/types"
	"testing"
)

func TestGet(t *testing.T) {
	data, err := Get("http://localhost:3000/hands?username=wgraham")
	if err != nil {
		t.Errorf("Get(wgraham): %s", err)
	}

	entry := types.Hand{"wgraham", 1, 100000, "KJo", 2000, "call", 200, nil, false, false }
	_, err = Post(entry)
	if err != nil {
		t.Errorf("Post(wgraham): %s", err)
	}

	data, err = Get("http://localhost:3000/hands?username=wgraham")
	if err != nil {
		t.Errorf("Get(wgraham): %s", err)
	}

	if len(data) == 0 {
		t.Errorf("Empty response")
	}
}

func TestPost(t *testing.T) {
	entry := types.Hand{"wgraham", 1, 100000, "KJo", 2000, "call", 200, nil, false, false }
	_, err := Post(entry)
	if err != nil {
		t.Errorf("Post(wgraham): %s", err)
	}

}
