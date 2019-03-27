package util

import (
	"math/rand"
	"time"
)

func InitializeCommand() {
	rand.Seed(time.Now().UnixNano())
}
