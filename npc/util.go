package npc

import "math/rand"

// dice rolls xDy and returns the sum.
func dice(count int, sides int) int {
	total := 0

	for i := 0; i < count; i++ {
		total += rand.Intn(sides) + 1
	}

	return total
}

// randRange calculates a random number within the specified bounds
// (inclusive).
func randRange(min int, max int) int {
	x := min

	delta := max - min
	if delta > 0 {
		x += rand.Intn(delta + 1)
	}

	return x
}

func filterIDs(numIDs int, shouldKeep func(id int) bool) []int {
	ids := make([]int, 0, numIDs)

	for i := 0; i < numIDs; i++ {
		if shouldKeep(i) {
			ids = append(ids, i)
		}
	}

	return ids
}
