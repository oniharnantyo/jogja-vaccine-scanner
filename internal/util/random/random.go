package random

import (
	"math/rand"
	"time"
)

type Random interface {
	RandomByRange(min, max int) int
}

type Randomizer struct {
}

func (r *Randomizer) RandomByRange(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func NewRandomizer() Random {
	return &Randomizer{}
}
