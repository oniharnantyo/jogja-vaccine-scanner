package random_test

import (
	"testing"

	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/util/random"
)

func TestRandomByRange(t *testing.T) {
	randomizer := random.NewRandomizer()
	res := randomizer.RandomByRange(1, 2000)

	t.Log(res)

	if res <= 1 && res >= 10 {
		t.Fail()
	}
}
