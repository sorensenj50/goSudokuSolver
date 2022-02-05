package main

import (
	"math/rand"
	"time"
)

type RandomGenerator struct {
	possibilities [gridSize]bool
}

func makeGenerator() RandomGenerator {
	var generator RandomGenerator
	possibilities := [9]bool{}
	for index := range [gridSize]bool{} {
		possibilities[index] = true
	}
	generator.possibilities = possibilities
	return generator
}

func (wrapper *RandomGenerator) get() int {
	num := randomNumber(0, len(wrapper.possibilities))
	for range [gridSize]int{} {
		if wrapper.possibilities[num] {
			return num + 1
		} else {
			num = (num + 1) % 9
		}
	}
	return 0
}

func (wrapper *RandomGenerator) setInvalid(num int) {
	wrapper.possibilities[num] = false
}

func (wrapper *RandomGenerator) reset() {
	trueValues := [gridSize]bool{true, true, true, true, true, true, true, true, true}
	wrapper.possibilities = trueValues
}

func randomNumber(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}
