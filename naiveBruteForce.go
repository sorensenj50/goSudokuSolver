package main

import (
	"math/rand"
	"time"
)

func setRandomSeed() {
	rand.Seed(time.Now().UnixNano())
}
