package store

import (
	"github.com/wesgraham/pokerHUD/pkg/types"
	"testing"
)

func TestGet(t *testing.T) {
	data, err := Get("http://localhost:3000/hands?uname=eq.wgraham")
	if err != nil {
		t.Errorf("Get(wgraham) 1: %s", err)
		return
	}

	entry := types.Hand{"wgraham", 1, 100000, "KJo", 2000, "call", 200, nil, false, false}
	_, err = Post(entry)
	if err != nil {
		t.Errorf("Post(wgraham): %s", err)
		return
	}

	data, err = Get("http://localhost:3000/hands?uname=eq.wgraham")
	if err != nil {
		t.Errorf("Get(wgraham) 2: %s", err)
		return
	}

	if len(data) == 0 {
		t.Errorf("Empty response")
		return
	}

}

func TestPost(t *testing.T) {
	entry := types.Hand{"hubertblack", 1, 100000, "KJo", 2000, "call", 200, nil, false, false}
	_, err := Post(entry)
	if err != nil {
		t.Errorf("Post(wgraham): %s", err)
		return
	}
}
