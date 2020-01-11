package store

import (
	"github.com/wesgraham/pokerHUD/pkg/types"
)

func Post(entry types.HandEntry){
	query := "http://localhost:3000/hands?username=" + username
}
