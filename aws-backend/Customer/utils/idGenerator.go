package utils

import (
	"math/rand"
	"time"
)

func IdGenerator() int {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(99999999)
	return id
}
