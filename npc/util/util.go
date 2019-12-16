package util

import (
	"math/rand"

	"github.com/badvassal/wllib/decode"
)

// dice rolls xDy and returns the sum.
func Dice(count int, sides int) int {
	total := 0

	for i := 0; i < count; i++ {
		total += rand.Intn(sides) + 1
	}

	return total
}

// randRange calculates a random number within the specified bounds
// (inclusive).
func RandRange(min int, max int) int {
	x := min

	delta := max - min
	if delta > 0 {
		x += rand.Intn(delta + 1)
	}

	return x
}

func CharacterNumSkills(ch decode.Character) int {
	for i, s := range ch.Skills {
		if s.Level == 0 {
			return i
		}
	}

	return len(ch.Skills)
}
